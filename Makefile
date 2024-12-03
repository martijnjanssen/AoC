DAY := $(shell date +%Y/%-d)

all: $(DAY)

total:
	@cd 2024; \
	go run main.go all

pprof:
	@cd 2024; \
	PPROF=yes go run main.go loop $(subst 2024/,,${DAY}); \
	go tool pprof --text test.prof

bench:
	@cd 2024; \
	go run main.go bench

2021/%:
	@cd 2021; \
	go run main.go $(subst 2021/,,$@)

2024/%:
	@cd 2024; \
	go run main.go $(subst 2024/,,$@)
