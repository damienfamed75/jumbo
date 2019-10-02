#include <stdio.h>

void hello(char *name, unsigned int name_len) {
	*name = 'D';
	name[name_len-1] = 'd';
}