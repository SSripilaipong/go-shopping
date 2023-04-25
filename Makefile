.PHONY: build
build:
	GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o build/main main.go

deploy:
	cd terraform; ../script/deploy-terraform.sh

terraform-init:
	cd terraform; ../script/init-terraform.sh