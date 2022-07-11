gen:
	@echo "Remove Current Proto Generation"
	@rm -rf "proto_model/*.go";
	@echo "Running Progo Generator ..."
	@protoc --proto_path=proto  --go_out=proto_model   --go-grpc_out=proto_model  --go-grpc_opt=paths=source_relative  --go_opt=paths=source_relative proto/*.proto
	@echo "Success"
	@make run-server

run-server:
	@echo "Running Server"
	@go run "server/main.go"

r:
	@echo "Running Server"
	@go run "server/main.go"
	
run-client:
	@echo "Running Client"
	@go run "client/main.go"
	
c:
	@echo "Running Client"
	@go run "client/main.go"
	
