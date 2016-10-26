#include "threads.h"
#include <unistd.h>
#include <stdlib.h>
#include <stdio.h>


void
f1(void){
	int i;
	fprintf(stderr, "SOY f1, id %d\n", curidthread());
	for(i = 0; i < 1024*1024*1024; i++){
		if(i == 1024*1024*1024/2){
			fprintf(stderr, "F1 cya later\n");
			yieldthread();
			fprintf(stderr, "F1 hello again\n");
		}
	}
	exitsthread();
}

void
f2(void  *val){
	fprintf(stderr, "f2 recibe %d\n", *(int *)val);
	sleep(1);
	yieldthread();
	fprintf(stderr, "Hello again f2!\n");
	exitsthread();
}

void
f3(void * ptr){
	fprintf(stderr, "f3 recibe %s\n", (char *) ptr);
	yieldthread();
	fprintf(stderr, "f3 hello again! %d\n", curidthread());
	exitsthread();
}

int main(){
	fprintf(stderr, "main: empezamos");
	int var;
	char str[] = "EH";
	fprintf(stderr, "main: inicializame los threads");
	initthreads();
	fprintf(stderr, "main: threads incializados");
	for(var = 1; var; var++)
		if(createthread((void *)(void *)f2, &var, 16*1024) < 0)
			break;
	sleep(1);
	yieldthread();
	sleep(1);
	yieldthread();
	createthread((void *)(void *)f1, NULL, 15*1024);
	createthread((void *)(void *)f2, &var, 4*1024);
	createthread(f3, (void *) str, 4*1024);
	fprintf(stderr, "MAIN SAYS BYE\n");
	exitsthread();
	exit(1);
}
