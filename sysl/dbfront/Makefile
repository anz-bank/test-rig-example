all: sysl

input = sysl/dbfront/model.sysl
app = dbfront
down = 
out = gen
# Current go import path
basepath = github.com/anz-bank/test-rig-example

SYSL = ./syslw
TRANSLOCATION = github.com/anz-bank/sysl-go/codegen/transforms
TRANSFORMS= svc_error_types.sysl svc_handler.sysl svc_interface.sysl svc_router.sysl svc_types.sysl
DOWNSTREAMTRANSFORMS = svc_client.sysl svc_error_types.sysl svc_types.sysl
GRAMMAR=github.com/anz-bank/sysl-go/codegen/grammars/go.gen.g
START=goFile

.PHONY: clean gen
sysl: gen downstream format


# Generate files with internal git service
gen:
	$(shell mkdir -p ${out}/${app})
	$(foreach file, $(TRANSFORMS), $(shell $(SYSL) codegen --basepath=$(basepath)/${out}/ --transform $(TRANSLOCATION)/$(file) --grammar ${GRAMMAR} --start ${START} --outdir=${out}/${app} --app-name ${app} $(input)))
	$(shell $(SYSL) generate-db-scripts --verbose --root="." --output-dir="$(out)/$(app)" --db-type="postgres" --app-names="$(app)" $(input))

downstream:
	$(foreach name, $(down), $(shell mkdir -p ${out}/${name}))
	$(foreach file, $(DOWNSTREAMTRANSFORMS), $(foreach downstream, $(down), $(shell $(SYSL) codegen --basepath=$(basepath)/${out}/ --transform $(TRANSLOCATION)/$(file) --grammar ${GRAMMAR} --start ${START} --outdir=${out}/${downstream} --app-name ${downstream} $(input))))

format:
	gofmt -s -w ${out}/${app}/*.go
	goimports -w ${out}/${app}/*.go
	$(foreach path, $(down), $(shell gofmt -s -w ${out}/${path}/*.go))
	$(foreach path, $(down), $(shell goimports -w ${out}/${path}/*.go))

# Remove the generated files
clean:
	rm -rf $(out)
