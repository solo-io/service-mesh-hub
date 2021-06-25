#----------------------------------------------------------------------------------
# Base
#----------------------------------------------------------------------------------
OUTDIR ?= _output
PROJECT ?= gloo-mesh

DOCKER_REPO ?= gcr.io/gloo-mesh
GLOOMESH_IMAGE ?= $(DOCKER_REPO)/gloo-mesh
CA_IMAGE ?= $(DOCKER_REPO)/cert-agent

BACKUP_DOCKER_REPO ?= quay.io/solo-io
BACKUP_GLOOMESH_IMAGE ?= $(BACKUP_DOCKER_REPO)/gloo-mesh
BACKUP_CA_IMAGE ?= $(BACKUP_DOCKER_REPO)/cert-agent

SOURCES := $(shell find . -name "*.go" | grep -v test.go)
export RELEASE ?= "true"
ifeq ($(TAGGED_VERSION),)
	TAGGED_VERSION := $(shell git describe --tags --dirty --always)
	export RELEASE := "false"
endif

VERSION ?= $(shell echo $(TAGGED_VERSION) | cut -c 2-)
.PHONY: print-version
print-version:
ifeq ($(TAGGED_VERSION),)
	exit 1
endif
	echo $(VERSION)

LDFLAGS := "-X github.com/solo-io/$(PROJECT)/pkg/common/version.Version=$(VERSION)"
GCFLAGS := all="-N -l"

print-info:
	@echo RELEASE: $(RELEASE)
	@echo TAGGED_VERSION: $(TAGGED_VERSION)
	@echo VERSION: $(VERSION)

#----------------------------------------------------------------------------------
# Code Generation
#----------------------------------------------------------------------------------

DEPSGOBIN=$(shell pwd)/$(OUTDIR)/.bin
export PATH:=$(DEPSGOBIN):$(PATH)
export GOBIN:=$(DEPSGOBIN)

.PHONY: fmt
fmt:
	go run golang.org/x/tools/cmd/goimports -w $(shell ls -d */ | grep -v vendor)

.PHONY: mod-download
mod-download:
	go mod download

.PHONY: clear-vendor-any
clear-vendor-any:
	rm -rf vendor_any

# Dependencies for code generation
mockgen: $(DEPSGOBIN)/mockgen
$(DEPSGOBIN)/mockgen:
	go build -o $@ github.com/golang/mock/mockgen

protoc-gen-go: $(DEPSGOBIN)/protoc-gen-go
$(DEPSGOBIN)/protoc-gen-go:
	go build -o $@ github.com/golang/protobuf/protoc-gen-go

protoc-gen-jsonshim: $(DEPSGOBIN)/protoc-gen-jsonshim
$(DEPSGOBIN)/protoc-gen-jsonshim:
	go build -o $@ istio.io/tools/cmd/protoc-gen-jsonshim

protoc-gen-ext: $(DEPSGOBIN)/protoc-gen-ext
$(DEPSGOBIN)/protoc-gen-ext:
	go build -o $@ github.com/solo-io/protoc-gen-ext

protoc-plugins: protoc-gen-go protoc-gen-ext protoc-gen-jsonshim


# Call all generated code targets
.PHONY: generated-code
generated-code: operator-gen \
				manifest-gen \
				go-generate \
				generated-reference-docs \
				fmt
	go mod tidy

#----------------------------------------------------------------------------------
# Go generate
#----------------------------------------------------------------------------------

# Run go-generate on all sub-packages
go-generate: mockgen
	go generate -v ./...

#----------------------------------------------------------------------------------
# Operator Code Generation
#----------------------------------------------------------------------------------

# Generate Operator Code
.PHONY: operator-gen
operator-gen: clear-vendor-any protoc-plugins
	go run -ldflags=$(LDFLAGS) -gcflags=$(GCFLAGS) codegen/generate.go

#----------------------------------------------------------------------------------
# Docs Code Generation
#----------------------------------------------------------------------------------

# Generate Reference documentation
.PHONY: generated-reference-docs
generated-reference-docs: clear-vendor-any update-licenses
	SKIP_CHANGELOG_GENERATION=true go run docs/cmd/docsgen.go

#----------------------------------------------------------------------------------
# Build
#----------------------------------------------------------------------------------

.PHONY: build-all-images
build-all-images: gloo-mesh-image cert-agent-image

#----------------------------------------------------------------------------------
# Build gloo-mesh controller + image
#----------------------------------------------------------------------------------

# for local development only; to build docker image, use gloo-mesh-linux-amd-64
.PHONY: gloo-mesh
gloo-mesh: $(OUTDIR)/gloo-mesh
$(OUTDIR)/gloo-mesh: $(SOURCES)
	go build -ldflags=$(LDFLAGS) -gcflags=$(GCFLAGS) -o $@ cmd/gloo-mesh/main.go

.PHONY: gloo-mesh-linux-amd64
gloo-mesh-linux-amd64: $(OUTDIR)/gloo-mesh-linux-amd64
$(OUTDIR)/gloo-mesh-linux-amd64: $(SOURCES)
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags=$(LDFLAGS) -gcflags=$(GCFLAGS) -o $@ cmd/gloo-mesh/main.go

# build image with gloo-mesh binary
# this is an alternative to using operator-gen to build the image
.PHONY: gloo-mesh-image
gloo-mesh-image: gloo-mesh-linux-amd64
	cp $(OUTDIR)/gloo-mesh-linux-amd64 build/gloo-mesh/ && \
	docker build -t $(GLOOMESH_IMAGE):$(VERSION) build/gloo-mesh/
	docker tag $(GLOOMESH_IMAGE):$(VERSION) $(BACKUP_GLOOMESH_IMAGE):$(VERSION) || true
	rm build/gloo-mesh/gloo-mesh-linux-amd64

.PHONY: gloo-mesh-image-push
gloo-mesh-image-push: gloo-mesh-image
ifeq ($(RELEASE),"true")
	docker push $(GLOOMESH_IMAGE):$(VERSION)
	docker push $(BACKUP_GLOOMESH_IMAGE):$(VERSION) || true
endif

.PHONY: gloo-mesh-image-load
gloo-mesh-image-load: gloo-mesh-image
	kind load docker-image --name mgmt-cluster $(GLOOMESH_IMAGE):$(VERSION)

#----------------------------------------------------------------------------------
# Build cert-agent + image
#----------------------------------------------------------------------------------

# for local development only; to build docker image, use cert-agent-linux-amd-64
.PHONY: cert-agent
cert-agent: $(OUTDIR)/cert-agent
$(OUTDIR)/cert-agent: $(SOURCES)
	go build -ldflags=$(LDFLAGS) -gcflags=$(GCFLAGS) -o $@ cmd/cert-agent/main.go

.PHONY: cert-agent-linux-amd64
cert-agent-linux-amd64: $(OUTDIR)/cert-agent-linux-amd64
$(OUTDIR)/cert-agent-linux-amd64: $(SOURCES)
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags=$(LDFLAGS) -gcflags=$(GCFLAGS) -o $@ cmd/cert-agent/main.go

# build image with cert-agent binary
# this is an alternative to using operator-gen to build the image
.PHONY: cert-agent-image
cert-agent-image: cert-agent-linux-amd64
	cp $(OUTDIR)/cert-agent-linux-amd64 build/cert-agent/ && \
	docker build -t $(CA_IMAGE):$(VERSION) build/cert-agent/
	docker tag $(CA_IMAGE):$(VERSION) $(BACKUP_CA_IMAGE):$(VERSION) || true
	rm build/cert-agent/cert-agent-linux-amd64

.PHONY: cert-agent-image-push
cert-agent-image-push: cert-agent-image
ifeq ($(RELEASE),"true")
	docker push $(CA_IMAGE):$(VERSION)
	docker push $(BACKUP_CA_IMAGE):$(VERSION) || true
endif

.PHONY: cert-agent-image-load
cert-agent-image-load: cert-agent-image
	kind load docker-image --name mgmt-cluster $(CA_IMAGE):$(VERSION)


#----------------------------------------------------------------------------------
# Build gloo-mesh cli (meshctl)
#----------------------------------------------------------------------------------

PACKR := go run github.com/gobuffalo/packr/packr

.PHONY: meshctl-linux-amd64
meshctl-linux-amd64: $(OUTDIR)/meshctl-linux-amd64
$(OUTDIR)/meshctl-linux-amd64: $(SOURCES)
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux $(PACKR) build -ldflags=$(LDFLAGS) -gcflags=$(GCFLAGS) -o $@ cmd/meshctl/main.go

.PHONY: meshctl-darwin-amd64
meshctl-darwin-amd64: $(OUTDIR)/meshctl-darwin-amd64
$(OUTDIR)/meshctl-darwin-amd64: $(SOURCES)
	CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin $(PACKR) build -ldflags=$(LDFLAGS) -gcflags=$(GCFLAGS) -o $@ cmd/meshctl/main.go

.PHONY: meshctl-windows-amd64
meshctl-windows-amd64: $(OUTDIR)/meshctl-windows-amd64.exe
$(OUTDIR)/meshctl-windows-amd64.exe: $(SOURCES)
	CGO_ENABLED=0 GOARCH=amd64 GOOS=windows $(PACKR) build -ldflags=$(LDFLAGS) -gcflags=$(GCFLAGS) -o $@ cmd/meshctl/main.go

.PHONY: build-cli
build-cli: meshctl-linux-amd64 meshctl-darwin-amd64 meshctl-windows-amd64

.PHONY: install-cli
install-cli:
	$(PACKR) build -ldflags=$(LDFLAGS) -gcflags=$(GCFLAGS) -o ${GOPATH}/bin/meshctl cmd/meshctl/main.go

#----------------------------------------------------------------------------------
# Push images
#----------------------------------------------------------------------------------

.PHONY: push-all-images
push-all-images: gloo-mesh-image-push cert-agent-image-push

#----------------------------------------------------------------------------------
# Helm chart
#----------------------------------------------------------------------------------
HELM_ROOTDIR := install/helm
# Include helm makefile so its targets can be ran from the root of this repo
include $(HELM_ROOTDIR)/helm.mk

# Generate Manifests from Helm Chart
.PHONY: chart-gen
chart-gen: clear-vendor-any
	go run -ldflags=$(LDFLAGS) -gcflags=$(GCFLAGS) codegen/generate.go -chart

.PHONY: manifest-gen
manifest-gen: install/gloo-mesh-default.yaml
	go run golang.org/x/tools/cmd/goimports -w $(shell ls -d */ | grep -v vendor)
	go mod tidy
install/gloo-mesh-default.yaml: chart-gen
	helm dependency update $(HELM_ROOTDIR)/gloo-mesh
	helm template --include-crds --namespace gloo-mesh $(HELM_ROOTDIR)/gloo-mesh > $@

#----------------------------------------------------------------------------------
# Test
#----------------------------------------------------------------------------------

# run all tests
# set TEST_PKG to run a specific test package
.PHONY: run-tests
run-tests:
	go run github.com/onsi/ginkgo/ginkgo -r -failFast -trace $(GINKGOFLAGS) \
		-ldflags=$(LDFLAGS) \
		-gcflags=$(GCFLAGS) \
		-progress \
		-race \
		-compilers=4 \
		-skipPackage=$(SKIP_PACKAGES) $(TEST_PKG)

# regen code+manifests, image build+push, and run all tests
# convenience for local testing
.PHONY: test-everything
test-everything: clean-generated-code generated-code manifest-gen run-tests

##----------------------------------------------------------------------------------
## Release
##----------------------------------------------------------------------------------

.PHONY: upload-github-release-assets
upload-github-release-assets: build-cli
ifeq ($(RELEASE),"true")
	go run ci/upload_github_release_assets.go
endif

#----------------------------------------------------------------------------------
# Third Party License Management
#----------------------------------------------------------------------------------
.PHONY: update-licenses
update-licenses:
	# check for GPL licenses, if there are any, this will fail
	GO111MODULE=on go run hack/oss_compliance/main.go osagen -c "GNU General Public License v2.0,GNU General Public License v3.0,GNU Affero General Public License v3.0"

	GO111MODULE=on go run hack/oss_compliance/main.go osagen -s "Mozilla Public License 2.0,GNU General Public License v2.0,GNU General Public License v3.0,GNU Lesser General Public License v2.1,GNU Lesser General Public License v3.0,GNU Affero General Public License v3.0"> docs/content/static/content/osa_provided.txt
	GO111MODULE=on go run hack/oss_compliance/main.go osagen -i "Mozilla Public License 2.0,GNU General Public License v2.0,GNU General Public License v3.0,GNU Lesser General Public License v2.1,GNU Lesser General Public License v3.0,GNU Affero General Public License v3.0"> docs/content/static/content/osa_included.txt


#----------------------------------------------------------------------------------
# Clean
#----------------------------------------------------------------------------------

.PHONY: clean
clean: clean-helm
	rm -f install/gloo-mesh-default.yaml
	rm -rf  _output/ vendor_any/

.PHONY: clean-generated-code
clean-generated-code:
	find pkg -name "*.pb.go" -type f -delete
	find pkg -name "*.hash.go" -type f -delete
	find pkg -name "*.gen.go" -type f -delete
	find pkg -name "*deepcopy.go" -type f -delete
