# Name of the binary
BINARY=main
DOCKER_IMAGE=<company>/<image>

# Values we want to pass for VERSION and BUILD
VERSION=1.0.0
BUILD=`git rev-parse HEAD`

GO_ENV=GOOS=`go env GOOS` GOARCH=`go env GOARCH` CGO_ENABLED=0

# Setup the -ldflags option
LDFLAGS=-ldflags -w
# Setup the -cflags option
CFLAGS=-a -tags netgo

# Default target
.DEFAULT_GOAL: $(BINARY)

# Builds the project
build:
	${GO_ENV} go build ${CFLAGS} ${LDFLAGS} -o ${BINARY} .

# Installs project: copies binaries
install:
	go install ${BINARY}

# Builds to linux target (Docker)
docker: GO_ENV=GOOS=linux GOARCH=amd64 CGO_ENABLED=0
docker: clean build tag clean
tag:
	docker build . -t ${DOCKER_IMAGE}:${BUILD}

# Cleans our project: deletes binaries
clean:
	if [ -f ${BINARY} ]; then rm ${BINARY}; fi

.PHONY: clean install
