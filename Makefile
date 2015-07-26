
deps:
	@go get ./...
	@go get github.com/davecheney/godoc2md
	@go get github.com/golang/lint/golint

doc:
	godoc2md github.com/jjeffery/goda/internal > internal/README.md

