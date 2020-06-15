# Phrase Service
Phrase service provides REST api for converting phrase to numeric hash.
## Run service
### With docker-compose
```
docker-compose up
```

### With docker
```
docker build --tag phrase-service .
docker run -p 8080:8080 phrase-service
```

### Standalone
```
go build -o service phraseservice/cmd
./service
```

## Run linter
Install [golangci-lint](https://github.com/golangci/golangci-lint).

Run from root directory:
```
golangci-lint run
```

## Run tests
### Unit
```
go test ./...
```
### Functional
To run functional tests use [http client](https://www.jetbrains.com/help/go/http-client-in-product-code-editor.html) and run all requests in phraseservice.http file.