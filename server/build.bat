set GOOS=linux
go build server.go
docker build -t usdzserver . 