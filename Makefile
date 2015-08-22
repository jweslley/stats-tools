PROGRAM=stats-tools
VERSION=0.0.1
LDFLAGS="-X stats.programVersion=$(VERSION)"

all: test

deps:
	go get ./...

install: deps
	go install -a -v -ldflags $(LDFLAGS) ./cmd/...

test: deps
	go test -v ./...

qa:
	go vet
	golint
	go test -coverprofile=.cover~
	go tool cover -html=.cover~

dist:
	@for os in linux darwin; do \
		for arch in 386 amd64; do \
			target=$(PROGRAM)-$$os-$$arch-$(VERSION); \
			echo Building $$target; \
			mkdir $$target;         \
			cp ./README.md ./LICENSE $$target; \
			for tool in $$(ls ./cmd); do \
				GOOS=$$os GOARCH=$$arch go build -ldflags $(LDFLAGS) -o $$target/$$tool ./cmd/$$tool ; \
			done; \
			tar -zcf $$target.tar.gz $$target; \
			rm -rf $$target;                   \
		done                                 \
	done

clean:
	rm -rf *.tar.gz
