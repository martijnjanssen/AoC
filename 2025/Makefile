CC = gcc
CFLAGS = -Wall -Wextra -O2

SRC := $(wildcard day_*.c)
FILES := $(SRC:.c=)

all: $(FILES)

%: %.c
	$(CC) $(CFLAGS) $< -o $@

clean:
	rm -f $(EXE)
