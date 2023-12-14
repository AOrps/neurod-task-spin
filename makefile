SHELL = /bin/sh
EXE = "spins"

all: compile

compile:
	go build -o $(EXE)

run: compile
	./$(EXE)

clean:
	rm -f $(EXE)

prune: clean
	rm -f neurod-task-spin *~
