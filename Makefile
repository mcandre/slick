VERSION=0.0.1

all:

install:
	sh -c 'cd cmd/slick && go install'

uninstall:
	-rm "$$GOPATH/bin/slick"

test:
	! slick -n examples

govet:
	find . -path "*/vendor*" -prune -o -name "*.go" -type f -exec go tool vet -shadow {} \;

golint:
	find . -path '*/vendor/*' -prune -o -name '*.go' -type f -exec golint {} \;

gofmt:
	find . -path '*/vendor/*' -prune -o -name '*.go' -type f -exec gofmt -s -w {} \;

goimports:
	find . -path '*/vendor/*' -prune -o -name '*.go' -type f -exec goimports -w {} \;

errcheck:
	errcheck -blank

opennota-check:
	aligncheck
	structcheck
	varcheck

megacheck:
	megacheck

lint: govet golint gofmt goimports errcheck opennota-check megacheck

port: archive-ports

archive-ports: bin
	zipc -chdir bin "slick-$(VERSION).zip" "slick-$(VERSION)"

bin:
	gox -output="bin/slick-$(VERSION)/{{.OS}}/{{.Arch}}/{{.Dir}}" ./cmd...

clean-ports:
	-rm -rf bin

clean: clean-ports
