.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/SmartMeshFoundation/Spectrum/cmd/swarm
	go build -o bin/geth     github.com/SmartMeshFoundation/Spectrum/cmd/geth
	go build -o bin/bootnode github.com/SmartMeshFoundation/Spectrum/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/SmartMeshFoundation/Spectrum/swarm/...
