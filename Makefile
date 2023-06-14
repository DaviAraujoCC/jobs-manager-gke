IMG ?= hu/jobs-manager
TARGET ?= development
PWD=$(shell pwd)
PORT=8080
INTERACTIVE:=$(shell [ -t 0 ] && echo i || echo d)
PROJECT_NAME=jobs-manager


run: docker-build
	@echo 'Running on http://localhost:$(PORT)/healthcheck'
	@docker run -t${INTERACTIVE} --rm \
		-v ${PWD}:/usr/app:delegated \
		-w /usr/app \
		--expose 8080 \
		-p $(PORT):8080 \
		--name ${PROJECT_NAME} \
		${IMG}
	

docker-build:
	@echo "Building Docker image..."
	@docker build --target ${TARGET} -t ${IMG} .

docker-push:
	@echo "Pushing Docker image..."
	@docker push ${IMG}