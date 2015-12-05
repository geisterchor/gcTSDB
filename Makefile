export GOPATH:=$(CURDIR)/Godeps/_workspace:$(GOPATH)

default: all

clean:
	rm -rf gcTSDB

install:
	GOPATH=$(shell godep path):$(GOPATH) go install

all:
	mkdir -p target
	go build -o gcTSDB

docker:
	docker run -i --rm -v "$(shell pwd):/go/src/geisterchor.com/gctsdb" -w /go/src/geisterchor.com/gctsdb golang bash -c "make test && make"
	docker build -f Dockerfile -t geisterchor/gctsdb:$(DOCKER_LABEL) .

run: all
	bash -c "./gcTSDB"

test:
	go test ${TESTFLAGS} -v ./...
