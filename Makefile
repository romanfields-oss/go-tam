swagger?=docker run --rm -it --env GOPATH=$(HOME)/go:/go --user $$(id -u):$$(id -g) --volume $(HOME):$(HOME) --workdir $$(pwd) quay.io/goswagger/swagger:v0.36.5


clean:
	rm -rf tam/

integration:
	go test ./... -tags=integration

.PHONY: clean integration