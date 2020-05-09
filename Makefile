testapi:
	@cd ./api/; \
	godog

testcore:
	@cd ./core/; \
	godog

testall: testcore testapi

all: testall
