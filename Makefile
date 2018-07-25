MAKEFLAGS += --silent
.DEFAULT_GOAL := all

TESTDIR = testdata
TERMFILES := $(shell find $(TESTDIR) -name '*.term')
GOLDFILES := $(TERMFILES:%.term=%.golden)

define generate_testdata
	erl -noshell -eval 'file:write_file("$(1)", erlang:term_to_binary($(2)))' -s init stop
endef

define show_testdata
	erl -noshell -eval 'F = "$(1)", {ok, Bin} = file:read_file(F), io:fwrite("| ~-35s | ~-55w | ~-55w |~n", [F, Bin, erlang:binary_to_term(Bin)])' -s init stop
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
	rm -f $(TESTDIR)/*.golden

.PHONY: generate-testdata
generate-testdata: $(GOLDFILES)

$(TESTDIR)/%.golden: $(TESTDIR)/%.term
	$(eval TERM := $(shell cat $^))
	$(call generate_testdata,$@,$(TERM))

.PHONY: show-testdata
show-testdata: $(GOLDFILES)
	printf '%.0s=' {1..155}
	echo
	printf '| %-35s | %-55s | %-55s |\n' 'FILE' 'BINARY' 'TERM'
	printf '%.0s-' {1..155}
	echo
	$(foreach bf,$^,$(call show_testdata,$(bf));)
	printf '%.0s=' {1..155}
	echo "\n"
