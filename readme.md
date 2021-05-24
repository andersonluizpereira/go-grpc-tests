#Como gerar o arquivo proto, preparando para o server
protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb