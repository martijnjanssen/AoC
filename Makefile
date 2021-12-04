DAY := $(shell date +%Y/%-d)

all: $(DAY)

2021/%:
	@cd $(subst /,/day_,$@); \
	go run main.go

