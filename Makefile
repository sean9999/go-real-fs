REPO=github.com/sean9999/go-real-fs
SEMVER := $$(git tag --sort=-version:refname | head -n 1)
BRANCH := $$(git branch --show-current)
REF := $$(git describe --dirty --tags --always)

.PHONY: test

info:
	@printf "REPO:\t%s\nSEMVER:\t%s\nBRANCH:\t%s\nREF:\t%s\n" $(REPO) $(SEMVER) $(BRANCH) $(REF)

tidy:
	go mod tidy

test:
	go test ./...

clean:
	go clean

docs:
	pkgsite -open .

publish:
	GOPROXY=https://goproxy.io,direct go list -m ${REPO}@${SEMVER}

