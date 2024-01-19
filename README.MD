# Welcome to Etov(echo to vanilla)
***
## Init
```bash
go mod tidy
```

## Build
```bash
docker buildx build --platform=linux/amd64 -t etov .
```
## Run
```bash
go run main.go
```
or run with docker
```bash
docker run -d -p 8181:8181 -e http_proxy=http://http_proxy/ -e https_proxy=http://https_proxy/ --name etov registry.cn-hangzhou.aliyuncs.com/ceerdecy/etov
```
