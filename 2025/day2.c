#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <inttypes.h>
#include <string.h>
#include <math.h>

int main(void) {
    FILE *fp = fopen("input_day2.txt", "r");
    // FILE *fp = fopen("test_day2.txt", "r");
    if (!fp) {
        perror("cannot read input file");
        return 1;
    }

    char *line = NULL;
    size_t len = 0;
    if (getline(&line, &len, fp) == -1) {
       perror("cannot read line");
       return 1;
    }

    uint64_t invalids = 0;

    char *pair = strtok(line, ",");
    uint64_t a, b;

    while (pair != NULL) {
        int scan_res = sscanf(pair, "%" SCNu64 "-%" SCNu64, &a, &b);
        if (scan_res != 2) {
            printf("scanning failed (scanned %d): %s\n", scan_res, pair);
            continue;
        }

        char str[50]="", x[25]="", y[25]="";
        int str_len;

        for (uint64_t i = a; i <= b; i++) {
            sprintf(str, "%" PRIu64, i);
            str_len = strlen(str) - 1;

            if (str_len % 2 == 0) {
                continue;
            }

            strncpy(x, str, str_len/2+1);
            strncpy(y, str + (str_len/2+1), str_len);

            if (strcmp(x, y) == 0) {
                invalids += i;
            }
        }

        pair = strtok(NULL, ",");
    }

    printf("Invalids: %" SCNu64 "\n", invalids);
}
