PROTOC_CMD = protoc
WEB_PROTOC_ARGS = --proto_path=./valuedig/web/v1/ --go_opt=paths=source_relative --go_out=./go/web/v1/ --go-grpc_out=./go/web/v1/ ./valuedig/web/v1/web.proto
TSD_PROTOC_ARGS = --proto_path=./valuedig/tsd/v1/ --go_opt=paths=source_relative --go_out=./go/tsd/v1/ --go-grpc_out=./go/tsd/v1/ ./valuedig/tsd/v1/tsd.proto

HTOML_TAG_FIX_CMD = htoml-tag-fix
WEB_HTOML_TAG_FIX_ARGS = go/web/v1/
TSD_HTOML_TAG_FIX_ARGS = go/tsd/v1/

BUILDCOLOR="\033[34;1m"
BINCOLOR="\033[37;1m"
ENDCOLOR="\033[0m"

ifndef V
	QUIET_BUILD = @printf '%b %b\n' $(BUILDCOLOR)BUILD$(ENDCOLOR) $(BINCOLOR)$@$(ENDCOLOR) 1>&2;
	QUIET_INSTALL = @printf '%b %b\n' $(BUILDCOLOR)INSTALL$(ENDCOLOR) $(BINCOLOR)$@$(ENDCOLOR) 1>&2;
endif


all: go-v1
	@echo ""
	@echo "build complete"
	@echo ""

go-v1:
	$(QUIET_BUILD)$(PROTOC_CMD) $(WEB_PROTOC_ARGS) $(CCLINK)
	$(QUIET_BUILD)$(HTOML_TAG_FIX_CMD) $(WEB_HTOML_TAG_FIX_ARGS) $(CCLINK)
	$(QUIET_BUILD)$(PROTOC_CMD) $(TSD_PROTOC_ARGS) $(CCLINK)
	$(QUIET_BUILD)$(HTOML_TAG_FIX_CMD) $(TSD_HTOML_TAG_FIX_ARGS) $(CCLINK)

