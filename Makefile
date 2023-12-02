#!make

# include .env
# If .env exists
ifneq ("$(wildcard .env)","")
	include .env
	export $(shell sed 's/=.*//' .env)
endif
.ONESHELL:
.PHONY: build clean deploy install rebuild serve
.SHELLFLAGS += -e

book: book/uk interactive-examples/docs

book/uk: book.toml content content/files/uk/index.md content/files/uk/SUMMARY.md original-content revamp/exe/move-media revamp/exe/format-links revamp/exe/inject-authors revamp/exe/rewire-paths revamp/exe/run-macros revamp/exe/strip-frontmatter  revamp/exe/writer src
	mdbook build
	./revamp/exe/move-media

book/interactive-examples: interactive-examples/docs revamp/exe/fix-interactive-examples
	cp -r interactive-examples/docs ./book/interactive-examples
	./revamp/exe/fix-interactive-examples

book/live-samples: book/uk
	cp -r live-samples ./book/live-samples

book/robots.txt: src/robots.txt
	cp ./src/robots.txt ./book/robots.txt

build: book book/interactive-examples book/live-samples book/robots.txt

content:
	git submodule add git@github.com:webdoky/content.git ./content

content/files/uk/index.md: src/index-template.md
	cp ./src/index-template.md ./content/files/uk/index.md

content/files/uk/SUMMARY.md: revamp/exe/summary
	./revamp/exe/summary

clean:
	rm -rf book

deploy: build revamp/exe/populate-algolia
	./revamp/exe/populate-algolia
	exit 1

install:
	curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
	cargo install mdbook
	go get ...

interactive-examples:
	git submodule add git@github.com:webdoky/interactive-examples.git

interactive-examples/docs: interactive-examples/node_modules
	cd interactive-examples && yarn build

interactive-examples/node_modules:
	cd interactive-examples && . ${HOME}/.nvm/nvm.sh && nvm use 18 && yarn install

node_modules:
	. ${HOME}/.nvm/nvm.sh && nvm use 18 && yarn install

original-content:
	git submodule add git@github.com:mdn/content.git ./original-content

revamp/exe/summary: revamp/generators/src/summary
	cd revamp/generators/src/summary && go build -o ../../../exe/

revamp/exe/fix-interactive-examples: revamp/postprocessors/src/fix-interactive-examples
	cd revamp/postprocessors/src/fix-interactive-examples && go build -o ../../../exe/

revamp/exe/move-media: revamp/postprocessors/src/move-media
	cd revamp/postprocessors/src/move-media && go build -o ../../../exe/

revamp/exe/populate-algolia: revamp/postprocessors/src/populate-algolia
	cd revamp/postprocessors/src/populate-algolia && go build -o ../../../exe/

revamp/exe/format-links: revamp/preprocessors/src/format-links revamp/preprocessors/src/helpers revamp/preprocessors/src/preprocessor
	cd revamp/preprocessors/src/format-links && go build -o ../../../exe/

revamp/exe/inject-authors: revamp/preprocessors/src/inject-authors revamp/preprocessors/src/preprocessor
	cd revamp/preprocessors/src/inject-authors && go build -o ../../../exe/

revamp/exe/rewire-paths: revamp/preprocessors/src/rewire-paths revamp/preprocessors/src/helpers revamp/preprocessors/src/preprocessor
	cd revamp/preprocessors/src/rewire-paths && go build -o ../../../exe/

revamp/exe/run-macros: revamp/preprocessors/src/run-macros revamp/preprocessors/src/helpers revamp/preprocessors/src/preprocessor
	cd revamp/preprocessors/src/run-macros && go build -o ../../../exe/

revamp/exe/strip-frontmatter: revamp/preprocessors/src/strip-frontmatter revamp/preprocessors/src/helpers revamp/preprocessors/src/preprocessor
	cd revamp/preprocessors/src/strip-frontmatter && go build -o ../../../exe/

revamp/exe/writer: revamp/preprocessors/src/writer revamp/preprocessors/src/helpers revamp/preprocessors/src/preprocessor
	cd revamp/preprocessors/src/writer && go build -o ../../../exe/

rebuild: clean build
