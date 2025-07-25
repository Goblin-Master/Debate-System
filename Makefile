.PHONY: user-generate user-run user-build
user-generate:
# --home 是指定模板目录
	goctl api go --api api/http/user.api --dir ./user --home ./template
user-run:
	go mod tidy && cd user && go run user.go
user-build:
	cd user/deploy && goctl docker --go ../user.go --exe user

.PHONY: chat-generate chat-run
chat-generate:
# --home 是指定模板目录
	goctl api go --api api/http/chat.api --dir ./chat --home ./template
chat-run:
	go mod tidy && cd chat && go run chat.go

.PHONY: grpc
grpc:
	@buf generate api/proto
