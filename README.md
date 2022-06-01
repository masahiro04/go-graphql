## 概要
goでGraphQLを扱うためのサンプルリポ


## 稼働方法
```
# graphql関連ファイルの生成
docker-compose run app go run github.com/99designs/gqlgen

# 実行
docker-compose up -d
```

## endpoints
```
http://localhost:8080/
http://localhost:8080/query
```
