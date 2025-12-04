#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <inttypes.h>
#include <string.h>
#include <math.h>

int main(void) {
    FILE *fp = fopen("input_day3.txt", "r");
    // FILE *fp = fopen("test_day3.txt", "r");
    if (!fp) {
        perror("cannot read input file");
        return 1;
    }

    char *line = NULL;
    size_t len = 0;

    int total = 0;

    while (getline(&line, &len, fp) != -1) {
        int a = -1;
        int a_pnt = -1;
        int b = -1;

        for (int i = 0; i < strlen(line)-2; i++) {
            int curr;
            curr = line[i] - '0';
            if (curr > a) {
               a = curr;
               a_pnt = i;
            }
        }

        for (int i = a_pnt + 1; i < strlen(line)-1; i++) {
            int curr;
            curr = line[i] - '0';
            if (curr > b) {
               b = curr;
            }
        }

        total += a*10+b;
    }

    printf("a: %d\n", total);
}
