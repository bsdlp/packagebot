go_version = 1.10.3
docker_workdir = /go/src/github.com/bsdlp/packagebot
docker_go := docker run --rm -v $(CURDIR):$(docker_workdir) -w $(docker_workdir) golang:$(go_version)
docker_lambda := docker run --rm -v $(CURDIR):$(docker_workdir) -w $(docker_workdir) bsdlp/lambda-builder:latest

.PHONY: deploy

lint:
	$(docker_go) gofmt -e -d ./
	$(docker_go) go vet ./...

test:
	$(docker_go) go test ./...

build_lambdas:

build: build_lambdas

deploy_lambdas:
deploy:
