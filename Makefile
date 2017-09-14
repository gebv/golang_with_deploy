VERSION=v0.1
GITHASH=`git log -1 --pretty=format:"%h" || echo "???"`
CURDATE=`date -u +%Y%m%d.%H%M%S`
GITURL=github.com
GITLOGIN=gebv
GITPRNAME=project_name

APPVERSION=${VERSION}-${GITHASH}:${CURDATE}

docker-prebuild:
	docker build -t ${GITPRNAME}-app-builder -f Dockerfile.build .
docker-build: docker-prebuild
	docker run -it --rm --name ${GITPRNAME}-app-make-build \
		-v "${PWD}":/go/src/${GITURL}/${GITLOGIN}/${GITPRNAME} \
		-w /go/src/${GITURL}/${GITLOGIN}/${GITPRNAME} \
		${GITPRNAME}-app-builder make build

build:
	CGO_ENABLED=0 go build \
			-o bin/app \
			-v \
			-ldflags "-X main.VERSION=${APPVERSION}" \
			-a ./main.go
.PHONY: build