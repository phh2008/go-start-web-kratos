# Kratos Project Template

## 工具安装
1.安装 protoc   
https://github.com/protocolbuffers/protobuf   
2.执行 make init，如果没有安装 make 就install 以下包：     
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest   
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest   
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest   
go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest   
go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest   
go install github.com/google/wire/cmd/wire@latest    
>安装完以上工具添加 protoc 和 gopath/bin 到环境变量   

## Create a service
```
# Create a template project
kratos new server

cd server
# Add a proto template
kratos proto add api/server/server.proto
# Generate the proto code
kratos proto client api/server/server.proto
# Generate the source code of service by proto file
kratos proto server api/server/server.proto -t internal/service

go generate ./...
go build -o ./bin/ ./...
./bin/server -conf ./configs
```
## Generate other auxiliary files by Makefile
```
# Download and update dependencies
make init
# Generate API files (include: pb.go, http, grpc, validate, swagger) by proto file
make api
# Generate all files
make all
```
## Automated Initialization (wire)
```
# install wire
go get github.com/google/wire/cmd/wire

# generate wire
cd cmd/server
wire
```

## Docker
```bash
# build
docker build -t <your-docker-image-name> .

# run
docker run --rm -p 8000:8000 -p 9000:9000 -v </path/to/your/configs>:/data/conf <your-docker-image-name>
```

