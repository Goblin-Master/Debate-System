.PHONY: user-http-generate user-run
user-http-generate:
# --home 是指定模板目录
	goctl api go --api api/http/user.api --dir ./user --home ./template
user-run:
	go mod tidy && cd user && go run user.go