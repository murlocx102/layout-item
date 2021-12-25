# This how we want to name the binary output
BINARY=layout-item

VERSION=`date +%Y%m%d%H%M%S`

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-w -s -X main.Version=${VERSION}"

# Builds the project
app:
	go build -o ${BINARY}   ${LDFLAGS}  cmd/app/main.go

# Build linux
linux: binary-linux

# Cleans our project: deletes binaries
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi


.PHONY: clean

# 编译前置
binary-linux:
	GOOS=linux GOARCH=amd64  go build -o ${BINARY}-linux   ${LDFLAGS}  cmd/app/main.go

production-tar:
	tar -cvzf ${BINARY}.tar.gz --exclude=configs/custom.yaml configs  ${BINARY}-linux	

