#include <stdio.h>

void hello(char *name, unsigned int name_len) {
	*name = 'D';
	name[name_len-1] = 'd';
}

#define _GNU_SOURCE
#include <stdio.h>
#include <stdlib.h>

void alphabet() {
	// Input file.
	FILE *rp;
	rp = fopen("~/input", "r+");
	if (rp == NULL)
		exit(EXIT_FAILURE);

	// Output file.
	FILE *op;
	op = fopen("~/output", "w");
	if (op == NULL)
		exit(EXIT_FAILURE);

	char   *line = NULL; // The buffer we will be reading each line into.
	size_t  len  = 0;		// The length of each line.
	size_t  read;		// The number of characters read from the file.

	while ( (read = getline(&line, &len, rp)) != -1 ) {
		fputs(line, op);
	}
	
	// Close the file pointers.
	fclose(rp);
	fclose(op);

	// Free buffer.
	if (line)
		free(line);
}