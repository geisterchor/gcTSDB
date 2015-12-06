default: all

clean:
	rm -rf target/

install:
	go install

all:
	mkdir -p target
	GO15VENDOREXPERIMENT=1 go build -o target/gcTSDB

docker:
	docker run -i --rm -v "$(shell pwd):/go/src/geisterchor.com/gcTSDB" -w /go/src/geisterchor.com/gcTSDB golang:1.5.2 bash -c "make test && make"
	docker build -f Dockerfile -t geisterchor/gctsdb:$(DOCKER_LABEL) .

run: all
	bash -c "./target/gcTSDB"

test:
	GO15VENDOREXPERIMENT=1 go test ${TESTFLAGS} -v ./...
