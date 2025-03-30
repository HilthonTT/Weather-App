include .env

.PHONY: gen-docs
gen-docs:
	@swag init --parseDependency -g ./api/main.go -d cmd,internal,services && swag fmt 
