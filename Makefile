TEST?=./...
NAME = $(shell awk -F\" '/^const Name/ { print $$2 }' main.go)
VERSION = $(shell awk -F\" '/^const Version/ { print $$2 }' main.go)
DOCKERREPO=ryane

all: push

build:
	docker build -t $(DOCKERREPO)/$(NAME) .
	docker tag $(DOCKERREPO)/$(NAME) $(DOCKERREPO)/$(NAME):$(VERSION)

push: build
	docker push $(DOCKERREPO)/$(NAME):$(VERSION)
	docker push $(DOCKERREPO)/$(NAME):latest
