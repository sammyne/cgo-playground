
.PHONY: all
all: foo bar

bar: foo
	@echo "building bar"
	@gcc -o bar bar.c -I. -L. -lfoo -Wl,-rpath=.

foo:
	@echo "building libfoo.so and libfoo.h"
	@go build -o libfoo.so -buildmode=c-shared foo.go

run: bar
	@./bar

.PHONY: clean
clean:
	rm -f bar *.so *.h
