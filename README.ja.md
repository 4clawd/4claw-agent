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

`agent/` ディレクトリは、`cli/` デスクトップアプリの背後で動作する**コア実行プロセス**です。

要点:

- `agent/` = コアエンジン（Go バイナリ、設定解析、ツール、gateway runtime）
- `cli/` = 操作 UI（Electron、エージェント管理、トレイ/ウィンドウ UX）

## `cli/` との関係

`cli/` は runtime を置き換えるものではなく、runtime を管理します:

1. `config.json` を書き込む
2. `4claw gateway --config ...` を開始/停止する
3. ログを取得する
4. バックアップの作成/インポート/エクスポートを行う

つまり、`agent/` が実行の中核で、`cli/` はコントロールプレーンです。

## `agent/` の主な機能

- Go 製 single-binary runtime
- 各種チャネル向け gateway モード
- 設定ベースの model/provider ルーティング
- ツールシステム（filesystem、shell、web、schedule、skills）
- エージェント実行ループとタスクオーケストレーション
- 設定駆動のポータブル運用

## `cli/` スクリーンショット（コントロール層）

このリポジトリは runtime core ですが、操作 UI は以下です:

### メイン画面

![CLI Main](./docs/images/main.png)

### 設定画面

![CLI Settings](./docs/images/setting.png)

## リポジトリ構成

- `cmd/`: エントリーポイントとコマンド
- `config/`: 設定サンプル
- `internal/`: コアロジック
- `docs/`: ドキュメントと roadmap
- `docs/images/`: README 用画像

## クイックスタート

```bash
make deps
make build
```

実行:

```bash
./4claw gateway
```

カスタム設定で実行:

```bash
./4claw gateway -c /path/to/config.json
```

## デスクトップと併用する流れ

1. `agent/` でバイナリをビルド/取得
2. `cli/resources/bin/` に配置
3. `cli/` アプリを起動
4. UI で複数 runtime を管理

## 分離設計の利点

- runtime を軽量・移植可能に維持
- UI と core を分離し開発速度を向上
- サーバーで headless 実行可能
- デスクトップで安全にマルチエージェント運用

## License

MIT
