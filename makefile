SHELL = /bin/sh
EXE = "spins"
VERSION = $(shell cat Dockerfile | awk '/^LABEL version/{gsub("\"",""); print }' | cut -d'=' -f2)

all: compile

compile:
	go build -o $(EXE)

run: compile
	./$(EXE)

clean:
	rm -f $(EXE)

prune: clean
	rm -f neurod-task-spin *~
	find . -type f -name '*~' -exec rm -f {} \;

# builds a container image
container:
	docker build -t neurod-task-spin:$(VERSION) -f Dockerfile .

# prunes container image generated from
container-prune:
	docker rmi neurod-task-spin
