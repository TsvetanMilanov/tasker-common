SHELL               := /bin/bash
BUILD_DIR           := .build
GO_NAMESPACE        := github.com/TsvetanMilanov/tasker-common

define find_recursive
$(shell find $1 -iname "$2")
endef

.PHONY: \
	build

$(BUILD_DIR)/vendor: $(CURDIR)/Gopkg.toml $(CURDIR)/Gopkg.lock
	dep ensure

$(BUILD_DIR)/build: $(call find_recursive,$(CURDIR),"*.go") $(BUILD_DIR)/vendor
	go build $(GO_NAMESPACE)/common

build: $(BUILD_DIR)/build
