.DEFAULT_GOAL := swagger

install_swagger:
	go install github.com/go-swagger/go-swagger/cmd/swagger@latest

swagger:
	@echo Ensure you have the swagger CLI or this command will fail.
	@echo You can install the swagger CLI with: go install github.com/go-swagger/go-swagger/cmd/swagger@latest
	@echo ....

	swagger generate spec -o ./swagger.yaml --scan-models