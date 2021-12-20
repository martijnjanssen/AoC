DAY := $(shell date +%Y/%-d)

all: $(DAY)

total:
	@cd 2021; \
	go run main.go all

bench:
	@cd 2021; \
	go run main.go bench

2021/%:
	@cd 2021; \
	go run main.go $(subst 2021/,,$@)
