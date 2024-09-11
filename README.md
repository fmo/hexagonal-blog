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

## Tagging the version

After creating golang/post module with 

```
go mod init github.com/fmo/hexagonal-blog/golang/post
```

Push the code first then run following to tag the version

```
git tag golang/post/v0.0.2 
git push origin golang/post/v0.0.2
```

And pull it from posts folder

```
go get github.com/fmo/hexagonal-blog/golang/post@v0.0.2 
```

## Make the request

```
grpcurl -d '{"title": "Hola", "body": "Volla"}' -plaintext localhost:3003 Post/Create
```
