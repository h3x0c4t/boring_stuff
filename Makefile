BIN_DIR=$(PWD)/agents

.PHONY: evilmsg
evilmsg:
	@echo "Building..."
	go build -o evilmsg cmd/server/main.go

.PHONY: agent_linux
agent_linux:
	@mkdir -p ${BIN_DIR}
	@echo "Building..."
	rm -f ${BIN_DIR}/agent_linux*
	go build -o ${BIN_DIR}/agent_linux -ldflags "-X main.URL=$(HIT_URL)" cmd/agent_linux/main.go
	chmod +x ${BIN_DIR}/agent_linux
	zip -j ${BIN_DIR}/agent_linux.zip ${BIN_DIR}/agent_linux

.PHONY: clean
clean:
	rm -f evilmsg
	rm -f evilmsg_db.sqlite3
	rm -rf ${BIN_DIR}
