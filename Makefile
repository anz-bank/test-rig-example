all: codegen

codegen: clientapp dbfront testrig

clientapp:
	make -f ./sysl/clientapp/Makefile all

dbfront:
	make -f ./sysl/dbfront/Makefile all

VARS=template.json
OUT=cmd
INPUT=sysl/dbfront/model.sysl

testrig:
	sysl test-rig --template="$(VARS)" --output-dir="$(OUT)" $(INPUT)

clean:
	rm -rf cmd/*

.PHONY: testrig
