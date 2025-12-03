#include <stdio.h>
#include <stdlib.h>

int main(void) {
    FILE *fp = fopen("input_day1.txt", "r");
    // FILE *fp = fopen("test_day1.txt", "r");
    if (!fp) {
        perror("cannot read input file");
        return 1;
    }

    int number = 50;
    int zeroes = 0;
    int passed_zero = 0;

    char *line = NULL;
    size_t len = 0;
    while (getline(&line, &len, fp) != -1) {
        char dir;
        int val;
        int old_number = number;

        int scan_res = sscanf(line, "%c%d", &dir, &val);
        if (scan_res != 2) {
            printf("scanning failed (scanned %d): %s", scan_res, line);
            continue;
        }

        while (val > 99) {
              val = val - 100;
              passed_zero++;
        }

        if (dir == 'R') {
            number = old_number + val;
        } else if (dir == 'L') {
            number = old_number - val;
        }

        if ((old_number < 100 && number > 100) || (old_number > 0 && number < 0)) {
           passed_zero = passed_zero + 1;
        }
        if (number > 99) {
            number = number - 100;
        } else if (number < 0) {
            number = number + 100;
        }

        if (number == 0) {
            zeroes++;
        }
    }

    printf("Zeroes: %d\nPassed Zero: %d\n", zeroes, zeroes + passed_zero);

    free(line);
    fclose(fp);

    return 0;
}
