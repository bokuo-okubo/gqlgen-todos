TOOLS=\
	github.com/99designs/gqlgen@v0.17.4

init:
	@for tool in $(TOOLS) ; do \
		go install $$tool; \
	done

gqlgen:
	go run github.com/99designs/gqlgen@v0.17.4 generate