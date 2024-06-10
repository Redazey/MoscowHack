# Клонирование googleapis
# Linux
```
yay -S protoc-gen-go protoc-gen-go-grpc
go get google.golang.org/protobuf/cmd/protoc-gen-go@latest
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
cd protos && git clone https://github.com/googleapis/googleapis.git && cd ..
```
# Windows
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
cd protos && git clone https://github.com/googleapis/googleapis.git && cd ..
```
# Генерация proto-файлов (генерацию проводим в backend)
```
protoc -I protos/googleapis -I protos protos/*.proto --go_out=gen/go/. --go-grpc_out=gen/go/.    
```