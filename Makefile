GOPATH:=$(shell go env GOPATH)

.PHONY: config
config:
	cp ./config-example.json ./config.json
	cp ./config-example.json ./test-config.json
