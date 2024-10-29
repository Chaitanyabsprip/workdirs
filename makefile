.DEFAULT_GOAL:=install
SOURCES := $(shell find . -type f -name '*.go')

clean:
	@rm -rd ./bin

install: ${SOURCES}
	@go install ./cmd/workdirs
	@go install ./cmd/worktrees

uninstall: clean
	@rm "$(which workdirs)"
	@rm "$(which worktrees)"
