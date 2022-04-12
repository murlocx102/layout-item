# This how we want to name the binary output
BINARY=layout-item

VERSION=`date +%Y%m%d%H%M%S`

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-w -s -X main.Version=${VERSION}"

# Remove private gopath information from the compiling machine. e: /home/xxx/work/go/... =>  go/...
RPATHFLAGS=-gcflags "-trimpath=$(GOPATH)" -asmflags "-trimpath=$(GOPATH)"  -trimpath

# Builds the project
app:
	go build -o ../../bin/${BINARY}  ${RPATHFLAGS}  ${LDFLAGS}  ../../cmd/admin/main.go

# Build linux
linux: binary-linux

# Cleans our project: deletes binaries
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi


.PHONY: clean

# 编译前置
binary-linux:
	GOOS=linux GOARCH=amd64  go build -o ../../bin/${BINARY}-linux  ${RPATHFLAGS}  ${LDFLAGS}  ../../cmd/admin/main.go && \
	mkdir -p "../../bin/configs" && \
	cp ../../configs/custom.yaml ../../bin/configs/

production-tar: binary-linux
	tar -cvzf ../../bin/${BINARY}.tar.gz -C ../../ ./configs/custom.yaml -C ./bin/ ${BINARY}-linux	

