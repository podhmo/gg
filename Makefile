readme:
	echo '# gg' > README.md
	echo '```console' >> README.md
	gg -h 2>&1 >> README.md || echo ok
	echo '```' >> README.md

build:
	mkdir -p bin
	vgo build -o bin/gg
