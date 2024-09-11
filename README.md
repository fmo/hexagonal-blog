# ORM

Gorm.io orm is being used we have 2 dependencies

```
gorm.io/driver/mysql
gorm.io/gorm
```

## Generate proto files

Go to proto/post folder and run following

```
protoc \
--go_out=../../golang/post/ \
--go_opt=paths=source_relative \
--go-grpc_out=../../golang/post/ \
--go-grpc_opt=paths=source_relative \
./post.proto
```
