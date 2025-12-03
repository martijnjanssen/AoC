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

    uint64_t invalids_a = 0;
    uint64_t invalids_b = 0;

    char *pair = strtok(line, ",");
    uint64_t a, b;

    while (pair != NULL) {
        int scan_res = sscanf(pair, "%" SCNu64 "-%" SCNu64, &a, &b);
        if (scan_res != 2) {
            printf("scanning failed (scanned %d): %s\n", scan_res, pair);
            continue;
        }

        int str_len;
        // Size of the array for the string to fit based on the highest number in the pair
        int malloc_len = (int)(log10(b)+1);
        char *str = (char*) malloc(malloc_len);
        char *x = (char*) malloc(malloc_len/2);
        char *y = (char*) malloc(malloc_len/2);

        // For each between the lower and upper bound
        for (uint64_t i = a; i <= b; i++) {
            // Create the string from integer
            sprintf(str, "%" PRIu64, i);
            str_len = strlen(str);

            // For each length less than half of the string length
            for (int l = str_len/2; l >= 1; l--) {
                // If remainder is non-zero
                if (str_len % l != 0) {
                    continue;
                }

                // reset values, string buffers
                memset(x,0,strlen(x));
                memset(y,0,strlen(y));
                int equal = 1;

                // Compare all subsequent offsets with the first portion
                strncpy(x, str, l);
                for (int ofst = 0; ofst <= str_len - l; ofst += l) {
                    strncpy(y, str + ofst, l);

                    if (strcmp(x, y) != 0) {
                        equal = 0;
                    }
                }

                // If all offsets are the same
                if (equal > 0) {
                    invalids_b += i;
                    // If there are only two offsets
                    if (l == str_len/2 && str_len % 2 == 0) {
                       invalids_a += i;
                    }
                    break;
                }
            }
        }

        pair = strtok(NULL, ",");
    }

    printf("a: %" PRIu64 ", b: %" PRIu64 "\n", invalids_a, invalids_b);
}
