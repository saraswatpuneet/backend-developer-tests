gen:
	cd ./proto && protoc --go_out=.  --go-grpc_out=. --grpc-gateway_out=logtostderr=true:. --go-grpc_opt=paths=source_relative --go_opt=paths=source_relative *.proto  --openapiv2_out=:swagger
