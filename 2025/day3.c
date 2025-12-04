#include <stdio.h>
#include <stdlib.h>
#include <inttypes.h>
#include <string.h>
#include <math.h>

uint64_t shift_left(char *line, int llen, int picks) {
    int *a = malloc(picks * sizeof *a);
    // Position picks at the end of the line
    for (int i = 0; i < picks; i++) {
        a[i] = llen - (picks - i);
    }

    // For each number pick, as long as it's within bounds and
    // not on the position of the pick preceding it, shift forward.
    // When it's larger or equal the the previous max pick,
    // shift forward, we want it as far to the left as possible.
    for (int i = 0; i < picks; i++) {
        for (int j = a[i]; j >= 0; j--) {
            if (i > 0 && j <= a[i-1])  {
               break;
            }
            int curr = line[a[i]]-'0';
            if (curr <= line[j]-'0') {
                a[i] = j;
            }
        }
    }

    // Calculate the score for the picks.
    uint64_t num = 0;
    for (int i = 0; i < picks; i++) {
        num += pow(10, picks - 1 - i)*(line[a[i]]-'0');
    }

    free(a);
    return num;
}

int main(void) {
    FILE *fp = fopen("input_day3.txt", "r");
    // FILE *fp = fopen("test_day3.txt", "r");
    if (!fp) {
        perror("cannot read input file");
        return 1;
    }

    char *line = NULL;
    size_t len = 0;

    int64_t total = 0;
    int64_t total_b = 0;

    while (getline(&line, &len, fp) != -1) {
        line[strcspn(line, "\n")] = '\0';
        int llen = strlen(line);
        
        uint64_t num = shift_left(line, llen, 2);
        total += num;

        num = shift_left(line, llen, 12);
        total_b += num;
    }

    printf("a: %" PRIu64 ", b: %" PRIu64 "\n", total, total_b);
}

