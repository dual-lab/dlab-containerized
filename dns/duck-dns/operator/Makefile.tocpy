#!/usr/bin/make -f
# ============================================================================ #
#
#  Author       : dmike
#  Date         : 17,November 2022
#  Description  : Principal Makefile for duckdns operator
#
# ============================================================================ #

# ============================================================================ #
#
# Make flag.
# Scope:
#	- do not use built-in rules
#	- do not use built-in variables
#	- do not print working directories
#
# ============================================================================ #
MAKEFLAGS += -rR --no-print-directory -s
# ############################################################################ #
PHONY := _all
_all:

$(CURDIR)/Makefile Makefile: ;

PHONY += all
_all: all

# ############################################################################ #
#
# Common variables
#
# ############################################################################ #
GEN_CODE_IMAGE := slok/kube-code-generator:v1.25.0
PROJECT_GROUP := duckdns:v1alpha1
PROJECT_PACKAGE := github.com/dual-lab/dlab-containerized/dns/duck-dns/operator
PKG_PACKAGE := $(PROJECT_PACKAGE)/pkg
DOCKER := podman
GO := go
# ############################################################################ #
#
# Code generation recipe. Use a docker image that contains the gen-code
# binary.
#
# ############################################################################ #
PHONY += generate
generate:
		@echo "code generation"
	 	@$(DOCKER) run --rm \
	 	-v $(CURDIR):/go/src/$(PROJECT_PACKAGE) \
	 	-e PROJECT_PACKAGE=$(PROJECT_PACKAGE) \
	 	-e CLIENT_GENERATOR_OUT=$(PKG_PACKAGE)/generated \
	 	-e APIS_ROOT=$(PKG_PACKAGE)/apis \
	 	-e GROUPS_VERSION=$(PROJECT_GROUP) \
	 	-e GENERATION_TARGETS="deepcopy,client,informer,lister" \
	 	$(GEN_CODE_IMAGE)

# ############################################################################## #
#
# Clean recipe + clean variable
#
# ############################################################################## #
CLEAN_DEEPCOPY := $(PKG_PACKAGE)/apis/duckdns/**/zz_generated.deepcopy.go
CLEAN_GENERATED := $(PKG_PACKAGE)/generated
CLEAN_FILES = $(CLEAN_DEEPCOPY) $(CLEAN_GENERATED)

PHONY += clean
clean: $(CLEAN_DEEPCOPY) $(CLEAN_GENERATED)

$(CLEAN_FILES):
		@echo "removing $@"
		@rm -rf $@

.PHONY : $(PHONY)