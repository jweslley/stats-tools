BUILD_DIR=bin
VERSION=0.0.1

all: build

build:
	go install -v ./cmd/...

deps:
	go get github.com/gobuild/gobuild3/packer

dist:
	packer --os linux  --arch amd64 --output stats-tools-linux-amd64-$(VERSION).zip
	packer --os linux  --arch 386   --output stats-tools-linux-386-$(VERSION).zip
	packer --os darwin --arch amd64 --output stats-tools-mac-amd64-$(VERSION).zip
	packer --os darwin --arch 386   --output stats-tools-mac-386-$(VERSION).zip

clean:
	rm -rf $(BUILD_DIR) *.zip
