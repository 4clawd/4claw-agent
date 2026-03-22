package channels

import (
	"bytes"
	"context"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

type weChatUploadedFile struct {
	FileKey                     string
	DownloadEncryptedQueryParam string
	AESKeyHex                   string
	FileSize                    int
	FileSizeCiphertext          int
}

func downloadAndDecryptWeChatMedia(ctx context.Context, cdnBaseURL, encryptedQueryParam, aesKeyBase64 string) ([]byte, error) {
	key, err := parseWeChatAESKey(aesKeyBase64)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, buildWeChatCDNDownloadURL(cdnBaseURL, encryptedQueryParam), nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("cdn download %d: %s", resp.StatusCode, string(body))
	}
	ciphertext, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return decryptAesECB(ciphertext, key)
}

func saveInboundWeChatMedia(data []byte, suffix string) (string, error) {
	dir := filepath.Join(os.TempDir(), "4claw_wechat_media")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", err
	}
	f, err := os.CreateTemp(dir, "wechat-*"+suffix)
	if err != nil {
		return "", err
	}
	defer f.Close()
	if _, err := f.Write(data); err != nil {
		return "", err
	}
	return f.Name(), nil
}

func maybeDownloadRemoteMedia(ctx context.Context, mediaRef string) (string, error) {
	if mediaRef == "" {
		return "", nil
	}
	if strings.HasPrefix(mediaRef, "http://") || strings.HasPrefix(mediaRef, "https://") {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, mediaRef, nil)
		if err != nil {
			return "", err
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			body, _ := io.ReadAll(resp.Body)
			return "", fmt.Errorf("remote media download %d: %s", resp.StatusCode, string(body))
		}
		ext := extensionFromContentType(resp.Header.Get("Content-Type"))
		if ext == "" {
			ext = filepath.Ext(mediaRef)
		}
		if ext == "" {
			ext = ".bin"
		}
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		return saveInboundWeChatMedia(data, ext)
	}
	if strings.HasPrefix(mediaRef, "file://") {
		u, err := url.Parse(mediaRef)
		if err != nil {
			return "", err
		}
		return u.Path, nil
	}
	if filepath.IsAbs(mediaRef) {
		return mediaRef, nil
	}
	return filepath.Abs(mediaRef)
}

func uploadWeChatMediaFile(ctx context.Context, api *weChatAPIClient, cdnBaseURL, filePath, toUserID string, mediaType int) (*weChatUploadedFile, error) {
	plaintext, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	fileKeyBytes := make([]byte, 16)
	if _, err := rand.Read(fileKeyBytes); err != nil {
		return nil, err
	}
	aesKeyBytes := make([]byte, 16)
	if _, err := rand.Read(aesKeyBytes); err != nil {
		return nil, err
	}
	ciphertext, err := encryptAesECB(plaintext, aesKeyBytes)
	if err != nil {
		return nil, err
	}
	fileKey := fmt.Sprintf("%x", fileKeyBytes)
	rawMD5 := fmt.Sprintf("%x", md5.Sum(plaintext))
	uploadResp, err := api.getUploadURL(ctx, weChatGetUploadURLRequest{
		FileKey:    fileKey,
		MediaType:  mediaType,
		ToUserID:   toUserID,
		RawSize:    len(plaintext),
		RawFileMD5: rawMD5,
		FileSize:   len(ciphertext),
		NoNeedThumb: true,
		AESKey:     fmt.Sprintf("%x", aesKeyBytes),
	})
	if err != nil {
		return nil, err
	}
	if uploadResp.UploadParam == "" {
		return nil, fmt.Errorf("wechat upload url response missing upload_param")
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, buildWeChatCDNUploadURL(cdnBaseURL, uploadResp.UploadParam, fileKey), bytes.NewReader(ciphertext))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/octet-stream")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("wechat cdn upload %d: %s", resp.StatusCode, string(body))
	}
	downloadParam := resp.Header.Get("x-encrypted-param")
	if downloadParam == "" {
		return nil, fmt.Errorf("wechat cdn upload response missing x-encrypted-param")
	}
	return &weChatUploadedFile{
		FileKey:                     fileKey,
		DownloadEncryptedQueryParam: downloadParam,
		AESKeyHex:                   fmt.Sprintf("%x", aesKeyBytes),
		FileSize:                    len(plaintext),
		FileSizeCiphertext:          len(ciphertext),
	}, nil
}

func buildWeChatCDNDownloadURL(cdnBaseURL, encryptedQueryParam string) string {
	return strings.TrimRight(cdnBaseURL, "/") + "/download?encrypted_query_param=" + url.QueryEscape(encryptedQueryParam)
}

func buildWeChatCDNUploadURL(cdnBaseURL, uploadParam, fileKey string) string {
	return strings.TrimRight(cdnBaseURL, "/") + "/upload?encrypted_query_param=" + url.QueryEscape(uploadParam) + "&filekey=" + url.QueryEscape(fileKey)
}

func extensionFromContentType(contentType string) string {
	if contentType == "" {
		return ""
	}
	if exts, _ := mime.ExtensionsByType(strings.TrimSpace(strings.Split(contentType, ";")[0])); len(exts) > 0 {
		return exts[0]
	}
	return ""
}

func mimeFromFilename(name string) string {
	if mimeType := mime.TypeByExtension(strings.ToLower(filepath.Ext(name))); mimeType != "" {
		return strings.Split(mimeType, ";")[0]
	}
	return "application/octet-stream"
}

func mediaItemFromUpload(filePath string, uploaded *weChatUploadedFile) weChatMessageItem {
	ext := strings.ToLower(filepath.Ext(filePath))
	baseMedia := &weChatCDNMedia{
		EncryptQueryParam: uploaded.DownloadEncryptedQueryParam,
		AESKey:            base64.StdEncoding.EncodeToString([]byte(uploaded.AESKeyHex)),
		EncryptType:       1,
	}
	switch {
	case isVideoExt(ext):
		return weChatMessageItem{
			Type: weChatMessageTypeVideo,
			VideoItem: &weChatVideoItem{
				Media:     baseMedia,
				VideoSize: uploaded.FileSizeCiphertext,
			},
		}
	case isImageExt(ext):
		return weChatMessageItem{
			Type: weChatMessageTypeImage,
			ImageItem: &weChatImageItem{
				Media:   baseMedia,
				MidSize: uploaded.FileSizeCiphertext,
			},
		}
	default:
		return weChatMessageItem{
			Type: weChatMessageTypeFile,
			FileItem: &weChatFileItem{
				Media:    baseMedia,
				FileName: filepath.Base(filePath),
				Len:      fmt.Sprintf("%d", uploaded.FileSize),
			},
		}
	}
}

func isImageExt(ext string) bool {
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif", ".webp", ".bmp":
		return true
	default:
		return false
	}
}

func isVideoExt(ext string) bool {
	switch ext {
	case ".mp4", ".mov", ".mkv", ".avi", ".webm":
		return true
	default:
		return false
	}
}
