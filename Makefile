# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: smc android ios smc-cross swarm evm all test clean
.PHONY: smc-linux smc-linux-386 smc-linux-amd64 smc-linux-mips64 smc-linux-mips64le
.PHONY: smc-linux-arm smc-linux-arm-5 smc-linux-arm-6 smc-linux-arm-7 smc-linux-arm64
.PHONY: smc-darwin smc-darwin-386 smc-darwin-amd64
.PHONY: smc-windows smc-windows-386 smc-windows-amd64

GOBIN = $(shell pwd)/build/bin
GO ?= latest

smc:
	build/env.sh go run build/ci.go install ./cmd/smc
	@echo "Done building."
	@echo "Run \"$(GOBIN)/smc\" to launch smc."

swarm:
	build/env.sh go run build/ci.go install ./cmd/swarm
	@echo "Done building."
	@echo "Run \"$(GOBIN)/swarm\" to launch swarm."

all:
	build/env.sh go run build/ci.go install

android:
	build/env.sh go run build/ci.go aar --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/smc.aar\" to use the library."

ios:
	build/env.sh go run build/ci.go xcode --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/Smc.framework\" to use the library."

test: all
	build/env.sh go run build/ci.go test

clean:
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# The devtools target installs tools required for 'go generate'.
# You need to put $GOBIN (or $GOPATH/bin) in your PATH to use 'go generate'.

devtools:
	env GOBIN= go get -u golang.org/x/tools/cmd/stringer
	env GOBIN= go get -u github.com/jteeuwen/go-bindata/go-bindata
	env GOBIN= go get -u github.com/fjl/gencodec
	env GOBIN= go install ./cmd/abigen

# Cross Compilation Targets (xgo)

smc-cross: smc-linux smc-darwin smc-windows smc-android smc-ios
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/smc-*

smc-linux: smc-linux-386 smc-linux-amd64 smc-linux-arm smc-linux-mips64 smc-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/smc-linux-*

smc-linux-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/smc
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/smc-linux-* | grep 386

smc-linux-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/smc
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/smc-linux-* | grep amd64

smc-linux-arm: smc-linux-arm-5 smc-linux-arm-6 smc-linux-arm-7 smc-linux-arm64
	@echo "Linux ARM cross compilation done:"
	@ls -ld $(GOBIN)/smc-linux-* | grep arm

smc-linux-arm-5:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-5 -v ./cmd/smc
	@echo "Linux ARMv5 cross compilation done:"
	@ls -ld $(GOBIN)/smc-linux-* | grep arm-5

smc-linux-arm-6:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-6 -v ./cmd/smc
	@echo "Linux ARMv6 cross compilation done:"
	@ls -ld $(GOBIN)/smc-linux-* | grep arm-6

smc-linux-arm-7:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-7 -v ./cmd/smc
	@echo "Linux ARMv7 cross compilation done:"
	@ls -ld $(GOBIN)/smc-linux-* | grep arm-7

smc-linux-arm64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm64 -v ./cmd/smc
	@echo "Linux ARM64 cross compilation done:"
	@ls -ld $(GOBIN)/smc-linux-* | grep arm64

smc-linux-mips:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/smc
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/smc-linux-* | grep mips

smc-linux-mipsle:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/smc
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/smc-linux-* | grep mipsle

smc-linux-mips64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/smc
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/smc-linux-* | grep mips64

smc-linux-mips64le:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/smc
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/smc-linux-* | grep mips64le

smc-darwin: smc-darwin-386 smc-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/smc-darwin-*

smc-darwin-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/smc
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/smc-darwin-* | grep 386

smc-darwin-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/smc
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/smc-darwin-* | grep amd64

smc-windows: smc-windows-386 smc-windows-amd64
	@echo "Windows cross compilation done:"
	@ls -ld $(GOBIN)/smc-windows-*

smc-windows-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/386 -v ./cmd/smc
	@echo "Windows 386 cross compilation done:"
	@ls -ld $(GOBIN)/smc-windows-* | grep 386

smc-windows-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/amd64 -v ./cmd/smc
	@echo "Windows amd64 cross compilation done:"
	@ls -ld $(GOBIN)/smc-windows-* | grep amd64
