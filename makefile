SHELL = /bin/sh
EXE = "spins"
VERSION = $(shell cat Dockerfile | awk '/^LABEL version/{gsub("\"",""); print }' | cut -d'=' -f2)
DOCKERNAME = neurod-task-spin

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
	docker build -t $(DOCKERNAME):$(VERSION) -t $(DOCKERNAME):latest -f Dockerfile .

# prunes container image generated from
container-prune:
	docker rmi $(DOCKERNAME)

container-run: container
	docker run --network=host --rm -it $(DOCKERNAME)
