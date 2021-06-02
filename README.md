# go-service-template

## Set Environment

```
export $(grep -v '^#' .env | xargs)
```

## Testing

### Install tool

```
go get github.com/fullstorydev/grpcui/...
go install github.com/fullstorydev/grpcui/cmd/grpcui
```

### Start gRPC Web UI

```
grpcui -plaintext localhost:5000
```
