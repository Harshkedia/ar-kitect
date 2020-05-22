set GOOS=linux
go build
docker build -t registry.gitlab.com/codeosol/ar-kitect . 