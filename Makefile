#NAME		:=	$(shell grep -E "^module " ./go.mod | sed -E -e "s/^module //g" -e "s/.*\///g")
DATE		:=	$(shell date)
GIT_STATUS		:=	$(shell git status --short)

all:	commit build commit
#all:	commit go-clean clean go-update test build commit

build:
#	go build -ldflags="-s -w" -trimpath -o "./bin/${NAME}" ./src/...
#	go build -ldflags="-s -w" -trimpath -o "./bin/" ./...

clean:
	-gh auth logout
	-go clean -i -r -x -cache -testcache -modcache -fuzzcache
	-rm -v go.mod
	-rm -v go.sum
	-find ./ -name ".DS_Store" -delete
	-find ./ -name "._.DS_Store" -delete

commit:
ifneq (${GIT_STATUS},)
	git add . && git commit -m "${DATE}" && git push
endif

init:
	gh auth login --with-token < ~/.git_token
#	go mod init ${NAME}
	go mod init ${PACKAGE}
	go get -u ./...
	go mod tidy

install:
	@echo ${NAME} ${PACKAGE} ${DATE} ${GIT_STATUS}

release:
	git add .
	git commit -m "${DATE}"
	git tag v${VERSION}
	git push origin v${VERSION}
	gh release create v${VERSION} --generate-notes --latest=true

race:
#	go run -race ./...

run:
#	go run -ldflags="-s -w" -trimpath ./...

test:
#	go test ./...

update:
	go get -u ./...
	go mod tidy

include Makefile.local
