# Клонирование googleapis
```
yay -S protoc-gen-go protoc-gen-go-grpc
go get google.golang.org/protobuf/cmd/protoc-gen-go
git clone https://github.com/googleapis/googleapis.git
```
# Генерация proto-файлов
```
 protoc -I . -I googleapis --go_out=. --go-grpc_out=. *.proto 
```