export GOPATH:=$(CURDIR)/Godeps/_workspace:$(GOPATH)

default: all

clean:
	rm -rf target/

install:
	GOPATH=$(shell godep path):$(GOPATH) go install

all:
	mkdir -p target
	go build -o target/gcTSDB

docker:
	docker run -i --rm -v "$(shell pwd):/go/src/geisterchor.com/gctsdb" -w /go/src/geisterchor.com/gctsdb golang bash -c "make test && make"
	docker build -f Dockerfile -t geisterchor/gctsdb:$(DOCKER_LABEL) .

run: all
	bash -c "./target/gcTSDB"

test:
	go test ${TESTFLAGS} -v ./...
