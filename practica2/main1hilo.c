#include "threads.h"
#include <unistd.h>
#include <stdlib.h>
#include <stdio.h>

int print = 0;
int count = 0;
void hola(void){
	for(int i = 0; i < 10000000; i++){
		if(print != 1){
			if(count!=0){
				fprintf(stderr, "%d\r\n", count);
			}
			fprintf(stderr, "hola ");
			print = 1;
			count = 1;
		}
		count++;
		yieldthread();
	}
	fprintf(stderr, "%d\n\rSale hilo 1\n\r", count);
	count=0;
	exitsthread();
}

void adios(void){
	for(int i = 0; i < 10000000; i++){
		if(print != 2){
			if(count!=0){
				fprintf(stderr, "%d\r\n", count);
			}
			fprintf(stderr, "adios ");
			print = 2;
		}
		count++;
		yieldthread();
	}
	fprintf(stderr, "Sale hilo 2\n\r");
	exitsthread();
}

int main(){
	initthreads();
	sleep(1);
	yieldthread();
	fprintf(stderr, "funciona");
	exitsthread();
	exit(1);
}
