<div align="center">
  <a href="./README.md"><kbd>English (Default)</kbd></a>
  <a href="./README.zh.md"><kbd>简体中文</kbd></a>
  <a href="./README.ja.md"><kbd>日本語</kbd></a>
  <a href="./README.fr.md"><kbd>Français</kbd></a>
  <a href="./README.pt-br.md"><kbd>Português (Brasil)</kbd></a>
  <a href="./README.vi.md"><kbd>Tiếng Việt</kbd></a>
</div>

<br />

<div align="center">
  <img src="./docs/images/logo.png" alt="4claw logo" width="128" />
</div>

<div align="center">
  <img src="./docs/images/banner.png" alt="4claw banner" width="100%" />
</div>

# 4claw Agent Core

Thư mục `agent/` là **tiến trình lõi** phía sau ứng dụng desktop trong `cli/`.

Tóm tắt:

- `agent/` = engine lõi (binary Go, xử lý config, hệ thống tools, runtime gateway)
- `cli/` = lớp giao diện điều khiển (Electron UI, quản lý agent, tray/window UX)

## Quan hệ với `cli/`

Ứng dụng desktop không thay thế runtime lõi, mà điều khiển runtime:

1. ghi `config.json`
2. start/stop `4claw gateway --config ...`
3. đọc log runtime
4. quản lý backup/import/export

Vì vậy `agent/` là tiến trình thực thi chính, `cli/` là control plane.

## Năng lực chính trong `agent/`

- Runtime Go dạng single-binary
- Gateway mode cho tích hợp kênh
- Routing model/provider theo cấu hình
- Tool system (filesystem, shell, web, schedule, skills)
- Vòng lặp thực thi agent và điều phối tác vụ
- Triển khai linh hoạt theo config

## Ảnh chụp giao diện `cli/` (lớp điều khiển)

Mặc dù repo này là runtime core, dưới đây là UI desktop dùng để vận hành runtime:

### Màn hình chính

![CLI Main](./docs/images/main.png)

### Màn hình cài đặt

![CLI Settings](./docs/images/setting.png)

## Cấu trúc repo

- `cmd/`: entrypoint và command
- `config/`: mẫu cấu hình
- `internal/`: logic lõi
- `docs/`: tài liệu và roadmap
- `docs/images/`: logo/banner/screenshot dùng cho README

## Quick Start

```bash
make deps
make build
```

Chạy gateway:

```bash
./4claw gateway
```

Chạy với config tùy chỉnh:

```bash
./4claw gateway -c /path/to/config.json
```

## Quy trình dùng cùng desktop

1. Build/tải binary từ `agent/`.
2. Đặt binary vào `cli/resources/bin/`.
3. Chạy ứng dụng Electron trong `cli/`.
4. Quản lý nhiều runtime instance bằng UI.

## Lợi ích của mô hình tách lớp

- Runtime nhẹ, dễ đóng gói
- UI tách rời core, dễ phát triển
- Hỗ trợ chạy headless trên server
- Hỗ trợ vận hành đa agent an toàn trên desktop

## License

MIT
