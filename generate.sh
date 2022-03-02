# in .proto file set below
#option go_package="./tic_tac_toe_proto/;package_name"
protoc --go_out=. --go-grpc_out=. ./tic_tac_toe_proto/tic_tac_toe.proto

go run server/main.go server/main_grpc.go

docker run --name tic_tac_toe_postgres -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5433:5432 -d postgres:10.10

docker exec -it tic_tac_toe_postgres bash

docker run -it -p 3000:3000 -p 50051:50051 go-dock-6 /bin/bash
docker build . -t  go-dock-6