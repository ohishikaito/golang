# 前提
## containerの立ち上げ
docker-compose build

docker-compose up -d

## ホットリロード
docker-compose exec app realize start
↑なくても行けるかも？要検証

## 単体で実行
docker-compose exec app go run main.go

or

docker run golang_app go run app/main.go