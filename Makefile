DOCKER ?= docker
BIN_DIR := $(CURDIR)/bin

.PHONY: all build clean distclean deployment_image push

all: build

build: distclean
	@mkdir -p $(BIN_DIR)
	@$(DOCKER) build -t kubernetes-nvidia-drivers:build -f Dockerfile.build --build-arg USER_ID="$(shell id -u)" $(CURDIR)
	@$(DOCKER) run --rm --net=host -v $(BIN_DIR):/go/bin:Z kubernetes-nvidia-drivers:build

clean:
	-@$(DOCKER) images -q kubernetes-nvidia-drivers | xargs $(DOCKER) rmi 2>/dev/null

distclean:
	@rm -rf $(BIN_DIR)
