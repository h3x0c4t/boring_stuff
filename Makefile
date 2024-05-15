BIN_DIR=$(PWD)/bin

.PHONY: evilmsg
evilmsg:
	@mkdir -p ${BIN_DIR}
	@echo "Building..."
	go build -o ${BIN_DIR}/evilmsg .

.PHONY: clean
clean:
	rm -rf ${BIN_DIR}
