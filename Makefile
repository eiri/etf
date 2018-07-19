MAKEFLAGS += --silent
.DEFAULT_GOAL := all

TESTDIR = testdata
TERMFILES := $(shell find $(TESTDIR) -name '*.term')
BINFILES := $(TERMFILES:%.term=%.bin)

define generate_testdata
	erl -noshell -eval 'file:write_file("$(1)", erlang:term_to_binary($(2)))' -s init stop
endef

define show_testdata
	erl -noshell -eval 'F = "$(1)", {ok, Bin} = file:read_file(F), io:fwrite("| ~-20s | ~-40w | ~-40w |~n", [F, Bin, erlang:binary_to_term(Bin)])' -s init stop
endef

define show_f
	echo $(1)
endef


.PHONY: all
all: deps test build

.PHONY: deps
deps:
	go get -t ./...

.PHONY: test
test:
	go test -v ./...

.PHONY: build
build:
	go build -v

.PHONY: clean
clean:
	go clean ./...

.PHONY: distclean
distclean: clean
	rm -f $(TESTDIR)/*.bin

.PHONY: generate-testdata
generate-testdata: $(BINFILES)

$(TESTDIR)/%.bin: $(TESTDIR)/%.term
	$(eval TERM := $(shell cat $^))
	$(call generate_testdata,$@,$(TERM))

.PHONY: show-testdata
show-testdata: $(BINFILES)
	printf '%.0s=' {1..110}
	echo
	printf '| %-20s | %-40s | %-40s |\n' 'FILE' 'BINARY' 'TERM'
	printf '%.0s-' {1..110}
	echo
	$(foreach bf,$^,$(call show_testdata,$(bf));)
	printf '%.0s=' {1..110}
	echo "\n"
