APP_NAME="poc-go-validate-arch"
DOCKER_IMAGE_INTERNAL_TOOLS="internal-tools-${APP_NAME}:latest"

internal-tools-install:
	@docker build -t "${DOCKER_IMAGE_INTERNAL_TOOLS}" -f Dockerfile.internal-tools .

internal-tools-exec:
	@docker run --rm -w=/go/app -v=${PWD}:/go/app "${DOCKER_IMAGE_INTERNAL_TOOLS}" $(command)

lint:
	make lint-arch
	make lint-directories

lint-arch:
	@echo "validates: architecture and dependencies"
	@make internal-tools-exec command="arch-go"

lint-directories:
	@echo "validates: directory structure"
	@make internal-tools-exec command="directory-validator ."
