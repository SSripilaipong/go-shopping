.PHONY: build
build:
	GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o build/main main.go

deploy:
	cd app/terraform; ../../script/deploy-terraform.sh

terraform-init:
	cd app/terraform; ../../script/init-terraform.sh

test:
	go test ./tests/...