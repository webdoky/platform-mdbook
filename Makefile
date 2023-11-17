#!make

include .env
export $(shell sed 's/=.*//' .env)

.ONESHELL:
.PHONY: build clean deploy install rebuild serve
.SHELLFLAGS += -e

book: book/uk interactive-examples/docs

book/uk: book.toml content content/files/uk/index.md content/files/uk/SUMMARY.md node_modules original-content postprocessors/move-media preprocessors/format-links preprocessors/inject-authors preprocessors/rewire-paths preprocessors/run-macros preprocessors/strip-frontmatter preprocessors/summary preprocessors/writer src
	. ${HOME}/.nvm/nvm.sh && nvm use 18 && mdbook build
	./postprocessors/move-media

book/interactive-examples: interactive-examples/docs postprocessors/fix-interactive-examples
	cp -r interactive-examples/docs ./book/interactive-examples
	./postprocessors/fix-interactive-examples

book/live-samples: book/uk
	cp -r live-samples ./book/live-samples

book/robots.txt: src/robots.txt
	cp ./src/robots.txt ./book/robots.txt

build: book book/interactive-examples book/live-samples book/robots.txt

content:
	git submodule add git@github.com:webdoky/content.git ./content

content/files/uk/index.md: src/index-template.md
	cp ./index-template.md ./content/files/uk/index.md

content/files/uk/SUMMARY.md:
	echo "# Зміст" > ./content/files/uk/SUMMARY.md

clean:
	rm -rf book

deploy: build postprocessors/populate-algolia
	./postprocessors/populate-algolia
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

postprocessors/fix-interactive-examples: postprocessors/src/fix-interactive-examples
	cd postprocessors/src/fix-interactive-examples && go build -o ../../fix-interactive-examples

postprocessors/move-media: postprocessors/src/move-media
	cd postprocessors/src/move-media && go build -o ../../move-media

postprocessors/populate-algolia: postprocessors/src/populate-algolia
	cd postprocessors/src/populate-algolia && go build -o ../../populate-algolia

preprocessors/format-links: preprocessors/src/format-links preprocessors/src/helpers preprocessors/src/preprocessor
	cd preprocessors/src/format-links && go build -o ../../format-links

preprocessors/inject-authors: preprocessors/src/inject-authors preprocessors/src/preprocessor
	cd preprocessors/src/inject-authors && go build -o ../../inject-authors

preprocessors/rewire-paths: preprocessors/src/rewire-paths preprocessors/src/helpers preprocessors/src/preprocessor
	cd preprocessors/src/rewire-paths && go build -o ../../rewire-paths

preprocessors/run-macros: preprocessors/src/run-macros preprocessors/src/helpers preprocessors/src/preprocessor
	cd preprocessors/src/run-macros && go build -o ../../run-macros

preprocessors/strip-frontmatter: preprocessors/src/strip-frontmatter preprocessors/src/helpers preprocessors/src/preprocessor
	cd preprocessors/src/strip-frontmatter && go build -o ../../strip-frontmatter

preprocessors/summary: preprocessors/src/summary preprocessors/src/helpers preprocessors/src/preprocessor
	cd preprocessors/src/summary && go build -o ../../summary

preprocessors/writer: preprocessors/src/writer preprocessors/src/helpers preprocessors/src/preprocessor
	cd preprocessors/src/writer && go build -o ../../writer

rebuild: clean build

serve: interactive-examples/docs content/files/uk/SUMMARY.md
	mdbook serve --port 3002
