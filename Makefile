OAPI = $(shell which oapi-codegen)
OPENAPI = $(shell which openapi-generator)

oapi-generate-backend:
	$(OAPI) -package generated server/greeting.yaml > server/generated/greeting.gen.go

openapi-generate-frontend:
	${OPENAPI} generate -i server/greeting.yaml -g typescript-axios -o client/src/api/greeting-api-client