#!make

build:
	mdbook build
deploy:
	exit 1
prepare:
	curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
	cargo install mdbook