CMD := ../../bin/gg resolve

default:
	${CMD} . > 00.output
	${CMD} ../../ >> 00.output
	${CMD} ${GOPATH}/src/golang.org/x/tools/go/loader/ >> 00.output
	${CMD} $(shell go env | grep ROOT | cut -d = -f 2)/src/go/printer >> 00.output
