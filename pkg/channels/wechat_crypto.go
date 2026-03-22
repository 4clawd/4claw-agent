package channels

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"strings"
)

func encryptAesECB(plaintext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	plaintext = pkcs7Pad(plaintext, block.BlockSize())
	out := make([]byte, len(plaintext))
	for bs, be := 0, block.BlockSize(); bs < len(plaintext); bs, be = bs+block.BlockSize(), be+block.BlockSize() {
		block.Encrypt(out[bs:be], plaintext[bs:be])
	}
	return out, nil
}

func decryptAesECB(ciphertext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(ciphertext)%block.BlockSize() != 0 {
		return nil, fmt.Errorf("ciphertext length %d is not a multiple of block size", len(ciphertext))
	}
	out := make([]byte, len(ciphertext))
	for bs, be := 0, block.BlockSize(); bs < len(ciphertext); bs, be = bs+block.BlockSize(), be+block.BlockSize() {
		block.Decrypt(out[bs:be], ciphertext[bs:be])
	}
	return pkcs7Unpad(out, block.BlockSize())
}

func pkcs7Pad(src []byte, blockSize int) []byte {
	padding := blockSize - (len(src) % blockSize)
	if padding == 0 {
		padding = blockSize
	}
	return append(src, bytes.Repeat([]byte{byte(padding)}, padding)...)
}

func pkcs7Unpad(src []byte, blockSize int) ([]byte, error) {
	if len(src) == 0 || len(src)%blockSize != 0 {
		return nil, fmt.Errorf("invalid padded plaintext length")
	}
	padding := int(src[len(src)-1])
	if padding <= 0 || padding > blockSize || padding > len(src) {
		return nil, fmt.Errorf("invalid padding size")
	}
	for _, b := range src[len(src)-padding:] {
		if int(b) != padding {
			return nil, fmt.Errorf("invalid padding bytes")
		}
	}
	return src[:len(src)-padding], nil
}

func parseWeChatAESKey(encoded string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}
	if len(decoded) == 16 {
		return decoded, nil
	}
	if len(decoded) == 32 && isHexString(string(decoded)) {
		return decodeHexString(string(decoded))
	}
	return nil, fmt.Errorf("unexpected aes key length %d", len(decoded))
}

func isHexString(s string) bool {
	for _, r := range s {
		switch {
		case r >= '0' && r <= '9':
		case r >= 'a' && r <= 'f':
		case r >= 'A' && r <= 'F':
		default:
			return false
		}
	}
	return true
}

func decodeHexString(s string) ([]byte, error) {
	s = strings.TrimSpace(s)
	if len(s)%2 != 0 {
		return nil, fmt.Errorf("invalid hex length")
	}
	out := make([]byte, len(s)/2)
	for i := 0; i < len(s); i += 2 {
		var v byte
		for j := 0; j < 2; j++ {
			v <<= 4
			switch c := s[i+j]; {
			case c >= '0' && c <= '9':
				v |= c - '0'
			case c >= 'a' && c <= 'f':
				v |= c - 'a' + 10
			case c >= 'A' && c <= 'F':
				v |= c - 'A' + 10
			default:
				return nil, fmt.Errorf("invalid hex digit")
			}
		}
		out[i/2] = v
	}
	return out, nil
}
