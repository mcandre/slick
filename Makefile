all:

install:
	sh -c 'cd cmd/slick && go install'

uninstall:
	-rm "$$GOPATH/bin/slick"

test:
	! slick -n examples
