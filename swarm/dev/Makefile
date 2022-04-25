.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/MeshBoxTech/mesh-chain/cmd/swarm
	go build -o bin/geth     github.com/MeshBoxTech/mesh-chain/cmd/geth
	go build -o bin/bootnode github.com/MeshBoxTech/mesh-chain/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/MeshBoxTech/mesh-chain/swarm/...
