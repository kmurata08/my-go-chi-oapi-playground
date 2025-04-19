oapi:
	oapi-codegen -config oapi-config/user.yaml api-docs/modules/user/openapi.yaml > internal/user/api.gen.go
