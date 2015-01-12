BUILD_DIR=bin

all: build

build:
	go install -v ./cmd/...

clean:
	rm -rf $(BUILD_DIR)
