WORKDIR=/mnt/qwkstow
CE=podman run -it --rm -v ./:${WORKDIR}:Z -w ${WORKDIR}
GO=${CE} docker.io/library/golang:1.27-alpine go
GDB=${CE} localhost/go-gdb:1.27 gdb -tui


.PHONY: test

build-container:
	podman build -t go-gdb:1.27 -f container/Debug.containerfile .

shell:
	${CE} /bin/sh

test:
	${GO} test -v ./lib/util

debug:
	rm -rf build/*
	${GO} build -ldflags=-compressdwarf=false -o ${WORKDIR}/build/util.test ${WORKDIR}/lib/util
	${CE}	docker.io/library/golang:1.27-alpine chmod 111 build/util.test
	${GDB} build/util.test
