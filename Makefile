BUILD_DIR=bin

all: build

build:
	gox -os="linux" -arch="amd64" -output="$(BUILD_DIR)/{{.OS}}_{{.Arch}}/{{.Dir}}" ./cmd/...

clean:
	rm -rf $(BUILD_DIR)
