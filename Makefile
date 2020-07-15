SYSL = ./syslw
VARS = template.json
OUT = cmd
INPUT = sysl/dbfront/model.sysl

all: gen

gen: clean clientapp dbfront testrig

clientapp:
	make -f ./sysl/clientapp/Makefile all

dbfront:
	make -f ./sysl/dbfront/Makefile all

testrig:
	${SYSL} test-rig --template="$(VARS)" --output-dir="$(OUT)" $(INPUT)
	$(shell mkdir -p $(OUT)/vendor)

clean:
	make -f ./sysl/clientapp/Makefile clean
	make -f ./sysl/dbfront/Makefile clean
	rm -rf $(OUT)/* docker-compose.yml

.PHONY: testrig
