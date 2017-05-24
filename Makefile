OUTFILES := $(patsubst %.go,%,$(wildcard *.go))

%: %.go
	go build -o $@ $<

all: $(OUTFILES)

clean:
	rm -rf $(OUTFILES)
