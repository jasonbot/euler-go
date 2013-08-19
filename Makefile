OUTFILES := $(patsubst %.go,%.out,$(wildcard *.go))

%.out: %.go
	go build -o $@ $<

all: $(OUTFILES)
