GOFILES = $(wildcard *.go)
INTFILES := $(patsubst %.go,%.6,$(wildcard *.go))
OUTFILES := $(patsubst %.go,%.out,$(wildcard *.go))

%.out: %.go
	go build -o $@ $<

all: $(OUTFILES)
