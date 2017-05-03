# todo: versioning with ldflags
# todo: test

TAG:=$(shell git describe --tags --abbrev=0)

all:
	env GOOS=linux GOARCH=386 gb build
	env GOOS=darwin GOARCH=386 gb build
