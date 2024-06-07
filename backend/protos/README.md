# Клонирование googleapis
# Linux
```
yay -S protoc-gen-go protoc-gen-go-grpc
go get google.golang.org/protobuf/cmd/protoc-gen-go@latest
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
git clone https://github.com/googleapis/googleapis.git
```
# Windows
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
git clone https://github.com/googleapis/googleapis.git
```
# Генерация proto-файлов (генерацию проводим в backend)
```
protoc -I googleapis -I protos protos/*.proto --go_out=gen/go/. --go-grpc_out=gen/go/.    
```