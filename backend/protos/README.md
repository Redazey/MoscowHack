# Клонирование googleapis
```
yay -S protoc-gen-go protoc-gen-go-grpc
go get google.golang.org/protobuf/cmd/protoc-gen-go
git clone https://github.com/googleapis/googleapis.git
```
# Генерация proto-файлов
- auth.proto
```
protoc -I . -I googleapis --go_out=. auth.proto 
```
- news.proto
```
protoc -I . -I googleapis --go_out=. news.proto 
```