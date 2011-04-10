GOFILES = $(wildcard *.go)
INTFILES := $(patsubst %.go,%.6,$(wildcard *.go))
OUTFILES := $(patsubst %.go,%.out,$(wildcard *.go))

%.6: %.go
	6g $<

%.out: %.6
	6l $<
	if [ -f $< ] rm $<

all: $(OUTFILES)
