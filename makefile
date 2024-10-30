.DEFAULT_GOAL:=install
SOURCES := $(shell find . -type f -name '*.go')

clean:
	@rm -rd ./bin

install: ${SOURCES}
	@go install ./cmd/work
	@go install ./cmd/workdirs
	@go install ./cmd/worktrees

work: ${SOURCES}
	@go run ./cmd/work

uninstall: clean
	@rm "$(which workdirs)"
	@rm "$(which worktrees)"
