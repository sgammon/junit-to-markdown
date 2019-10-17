IMAGE?=junit-to-markdown
TAG?=latest
GO111MODULE := on
export GO111MODULE

.PHONY: dockerbuild
dockerbuild:
	docker build -t $(IMAGE):$(TAG) .

.PHONY: dockerpush
dockerpush: dockerbuild
	docker push $(IMAGE):$(TAG)

.PHONY: deps
deps:
	go get

.PHONY: build
build:
	go build

.PHONY: clean
clean:
	go clean
	rm -f junit-to-markdown