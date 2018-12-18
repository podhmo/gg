readme:
	echo '# gg' > README.md
	echo '```console' >> README.md
	gg -h 2>&1 >> README.md || echo ok
	echo '```' >> README.md

test: bin/gg
	for i in $(shell find . -mindepth 2 -name Makefile); do make -C `dirname $$i`; done
	git diff

bin/gg: main.go
	mkdir -p bin
	vgo build -o bin/gg
