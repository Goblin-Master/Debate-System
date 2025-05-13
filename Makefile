.PHONY: generate run
generate:
# --home 是指定模板目录
	goctl api go --api app.api --dir . --home ./template
run:
	go mod tidy && go run user.go