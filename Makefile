run:
	go run main.go

install-oapi-codegen:
	@which oapi-codegen > /dev/null || go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.3.0

generate-oapi: install-oapi-codegen
	oapi-codegen -generate types -package api ./api/openapi.yaml > ./api/types.go
	oapi-codegen -generate server -package api ./api/openapi.yaml > ./api/server.go
