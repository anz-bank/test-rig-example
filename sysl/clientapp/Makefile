all: sysl

input = sysl/clientapp/model.sysl
app = clientapp
down = dep1 dep2
out = gen
# Current go import path
basepath = github.com/anz-bank/test-rig-example

TRANSLOCATION = github.com/anz-bank/sysl-go/codegen/transforms
TRANSFORM = svc_error_types.sysl svc_handler.sysl svc_interface.sysl svc_router.sysl svc_types.sysl
DOWNSTREAMTRANSFORMS = svc_client.sysl svc_error_types.sysl svc_types.sysl
GRAMMAR = github.com/anz-bank/sysl-go/codegen/grammars/go.gen.g
START = goFile

.PHONY: clean gen
sysl: clean gen downstream format

# Generate files with internal git service
gen:
	$(foreach file, $(TRANSFORMS), $(shell sysl codegen --basepath=$(basepath)/${out}/ --transform $(TRANSLOCATION)/$(file) --grammar ${GRAMMAR} --start ${START} --outdir=${out}/${app} --app-name ${app} $(input)))

downstream:
	$(foreach file, $(DOWNSTREAMTRANSFORMS), $(foreach downstream, $(down), $(shell sysl codegen --basepath=$(basepath)/${out}/ --transform $(TRANSLOCATION)/$(file) --grammar ${GRAMMAR} --start ${START} --outdir=${out}/${downstream} --app-name ${downstream} $(input))))

format:
	gofmt -s -w ${out}/${app}/*
	goimports -w ${out}/${app}/*
	$(foreach path, $(down), $(shell gofmt -s -w ${out}/${path}/*))
	$(foreach path, $(down), $(shell goimports -w ${out}/${path}/*))

# Remove the generated files
clean:
	rm -rf $(out)
