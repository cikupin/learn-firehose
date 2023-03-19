.PHONY: build
build:
	go build -ldflags="-w -s" -tags=timetzdata -o bin/learn-firehose

.PHONY: build-for-scratch
build-for-scratch:
	CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -tags=timetzdata -o bin/learn-firehose

.PHONY: image-scratch
image-scratch: build-for-scratch
	docker build -t learn-firehose:latest .