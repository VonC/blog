cmd /v /c "set GOOS=linux&& set GOARCH=amd64&& go build -o hello ..\hello.go"
docker build -t hello .