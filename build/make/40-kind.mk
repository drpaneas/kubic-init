# running `kind`
# see the build/kind/README.md doc for more instructions

# a temporary directory for building the base image
KIND_TMP_DIR  = /tmp/kind-kubic-build

# name:tag for the images
BASE_IMG_NAME = kindest/base:kubic
NODE_IMG_NAME = kindest/node:kubic

# any extra args for the build
BASE_BUILD_ARGS =
NODE_BUILD_ARGS =

get-kind:
	@echo ">>> Getting kind..."
	@command kind >/dev/null 2>&1 || GO111MODULE=off go get github.com/kubernetes-sigs/kind

kind-base-image: get-kind
	# kind build base-image doesn't work if a go.mod is in any top directory
	# so we must do some dirty hack and copy everything to a temp dir
	@echo ">>> Copying to a temporary directory..."
	rm -rf $(KIND_TMP_DIR)
	mkdir -p $(KIND_TMP_DIR)
	cp -r build/kind/base/* $(KIND_TMP_DIR)/
	@echo ">>> Building base image..."
	cd $(KIND_TMP_DIR) && \
		kind build base-image \
		--image "$(BASE_IMG_NAME)" \
		--loglevel debug \
		--source $(KIND_TMP_DIR) $(BASE_BUILD_ARGS)
	rm -rf $(KIND_TMP_DIR)

kind-node-image: get-kind
	@echo ">>> Building node image..."
	kind build node-image --image "$(NODE_IMG_NAME)" \
		--loglevel debug \
		--base-image "$(BASE_IMG_NAME)" $(NODE_BUILD_ARGS)

kind-images: kind-base-image kind-node-image
