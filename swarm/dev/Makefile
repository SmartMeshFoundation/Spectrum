.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/SmartMeshFoundation/SMChain/cmd/swarm
	go build -o bin/geth     github.com/SmartMeshFoundation/SMChain/cmd/geth
	go build -o bin/bootnode github.com/SmartMeshFoundation/SMChain/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/SmartMeshFoundation/SMChain/swarm/...
