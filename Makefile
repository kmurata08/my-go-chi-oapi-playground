oapi:
	oapi-codegen -config api-docs/modules/user/cfg.yaml api-docs/modules/user/openapi.yaml

oapi-common:
	oapi-codegen -config api-docs/common/cfg.yaml api-docs/common/openapi.yaml