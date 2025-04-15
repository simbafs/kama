# kama

> /kam-á/  
> 橘子的臺語    
> Tangerine in Taiwanese Hokkien

[English version](./README.md)

`kama` 是一個用來建構 **前後端分離（frontend-backend separated）** 網頁應用的 Golang 工具。

在開發階段，它會將所有未處理的 HTTP 請求轉發給前端開發伺服器（例如 Vite、Next.js）。  
在正式部署時，它會從內嵌的檔案系統中提供靜態資源，並允許以本地目錄覆寫內嵌檔案，無需重新編譯即可調整靜態檔案內容。

## 特點

- 支援 Go 標準 `net/http` 與 `gin` 框架
- 開發模式自動轉發請求至前端 dev server
- 部署模式內嵌並提供靜態檔案
- 可選擇使用本機目錄覆寫內嵌檔案

## 安裝

```bash
go get github.com/simbafs/kama
```

# 快速總覽

```go
//go:embed all:static
var embededFS embed.FS

kama.New(embeddedFS,
  kama.WithDevServer("http://localhost:3001"),
  kama.WithPath("static"),
)
```

# 使用方式

請參考 [\_example/](./_example/) 資料夾中的範例。完整範例請參考 [counter](https://github.com/simbafs/counter)

# License

[MIT](./LICENSE)
