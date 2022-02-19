all:
	docker build -t gcr.io/my-gcp-project/my-kube-scheduler:1.0 -f scheduler/Dockerfile .

test:
	go test -v ./...

build:
	go build -o bin/scheduler ./cmd/scheduler