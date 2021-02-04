gen-gqlgen:
	go run cmd/gqlgen/main.go
GCP_PROJECT := $(shell gcloud config get-value project)
VERSION := $$(make -s show-version)
CURRENT_REVISION := $(shell git rev-parse --short HEAD)
GOBIN ?= $(shell go env GOPATH)/bin
export GO111MODULE ?= on

.PHONY: test
test:
	go test -v ./...

#.PHONY: deploy
#deploy:
#	gcloud app deploy -q

.PHONY: show-version
show-version: $(GOBIN)/gobump
	@gobump show -r .

$(GOBIN)/gobump:
	@cd && go get github.com/x-motemen/gobump/cmd/gobump

.PHONY: tag
tag:
	git tag -a "v$(VERSION)" -m "Release $(VERSION)"
	git push --tags

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	golint -set_exit_status ./...

.PHONY: vet
vet:
	go vet ./...

build:
	gcloud builds submit --tag gcr.io/$(GCP_PROJECT)/$(IMAGE)

deploy: build
	gcloud run deploy $(IMAGE) --image gcr.io/$(GCP_PROJECT)/$(IMAGE) --platform managed --region us-central1
