.PHONY: build
build:
	go build -ldflags="-w -s" -tags=timetzdata -o bin/learn-firehose

.PHONY: build-for-scratch
build-for-scratch:
	CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -tags=timetzdata -o bin/learn-firehose

.PHONY: image-scratch
image-scratch: build-for-scratch
	docker build -t learn-firehose:latest .

.PHONY: generate-proto
generate-proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative payload/schema.proto
	protoc --descriptor_set_out=./payload/schema.desc --include_imports ./payload/schema.proto