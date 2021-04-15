# 前提
## containerの立ち上げ
docker-compose build

docker-compose up -d

## 単体で実行
docker-compose exec app go run main.go

or

docker run golang_app go run main.go

# デバッカー
`.vscode/launch.json`で
`${workspaceFolder}/app`って書いて、
appディレクトリ配下のファイルを読むようにしてる