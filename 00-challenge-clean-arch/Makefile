run:
	cd cmd/ordersystem && go run main.go wire_gen.go

generateGraphQL:
	cd internal/infra/graph && go run github.com/99designs/gqlgen generate

generateProto:
	protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto


evans:
	evans internal/infra/grpc/protofiles/order.proto -p 50051

migrateCreate:
	migrate create -ext=sql -dir=internal/infra/database/migrations -seq init

migrateUp:
	migrate -path internal/infra/database/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose up

migrateDown:
	migrate -path internal/infra/database/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose down