#include <stdlib.h>
#include <unistd.h>
#include <stdio.h>
#include "threads.h"
#include <ucontext.h>
#include <sys/time.h>
#include <err.h>

enum{
	AVILIABLE = 0,
	READY = 1,
	STOP = 2,
	RUNNING = 3,
	MAX_THREAD = 32
};

typedef struct thread_t thread_t;

typedef struct timeval timeval;

struct thread_t{
	void * stack_p;
	int stacksize;
	ucontext_t u_c;
	int id;
	timeval s_time;
	int state;
};

thread_t ttable[MAX_THREAD];
unsigned int current_thread = 0;
unsigned int id = -1;

int getId(){
	return id++;
}

void initthreads(void){
	fprintf(stderr, "inicializando los hilos");
	if(getcontext(&ttable[0].u_c))
		err(1, "getcontext initthread");
	if(gettimeofday(&ttable[0].s_time, NULL))
		err(1, "gettimeofday init");
	ttable[0].id = getId();
	ttable[0].state = RUNNING;
	fprintf(stderr, "inicializando los hilos");
}

int createthread(void (*mainf)(void*), void *arg, int stacksize) {
	fprintf(stderr, "creando hilo...");
	int tindex_void = -1;
	for(int i = 0; i < MAX_THREAD; i++){
		if(&ttable[i].state == AVILIABLE){
			tindex_void = i;
			break; 
		}
	}
	if(tindex_void == -1 ||
			getcontext(&ttable[tindex_void].u_c) || 		
			gettimeofday(&ttable[0].s_time, NULL))
		return -1;
	
	ttable[tindex_void].id = getId();
	int mallocea = 0;
	if(&ttable[tindex_void].stacksize == NULL) {
		mallocea = 1;
	} else if(ttable[tindex_void].stacksize != stacksize){
		free(ttable[tindex_void].stack_p);
		mallocea = 1;
	}
	if(mallocea){
		ttable[tindex_void].stack_p = malloc(stacksize);
	}
	ttable[tindex_void].u_c.uc_stack.ss_sp = ttable[tindex_void].stack_p;
	ttable[tindex_void].u_c.uc_stack.ss_size = ttable[tindex_void].stacksize;
	ttable[tindex_void].u_c.uc_link = &ttable[current_thread].u_c;
	ttable[tindex_void].state = READY;
	makecontext(&ttable[tindex_void].u_c, (void(*))mainf, 1, arg);
	fprintf(stderr, "hilo creado");
	return ttable[tindex_void].id;
} 

void exitsthread(void) {
	fprintf(stderr, "intentando finalizar el hilo");
	ttable[current_thread].state = AVILIABLE;
	yieldthread();
	fprintf(stderr, "hilo finalizado");
}

int nextCT(){
	fprintf(stderr, "buscando siguiente ct");
	do {
		if(current_thread + 1 == MAX_THREAD){
			current_thread = 1;
		} else { 
			current_thread ++;
		}
	} while(ttable[current_thread].state != READY);
	fprintf(stderr, "encontrado el siguiente ct: %d", current_thread);
	return current_thread;
}

int quantumSpent(){
	fprintf(stderr, "quantum gastado.....");
	if(ttable[current_thread].state == AVILIABLE){
		return 1;
	}
	timeval t;
	if(gettimeofday(&t, NULL)){
		err(1, "gettimeofday");
	}
	long now = t.tv_usec / 1000 + t.tv_sec * 1000;
	timeval t2 = ttable[current_thread].s_time;
	long before = t2.tv_usec / 1000 + t2.tv_sec * 1000;
	fprintf(stderr, "tiempo del cuantum: %ld", now-before);
	return now - before > 200;
}

void yieldthread(void) {	
	fprintf(stderr, "cambiando hilos...");
	if(!quantumSpent()){
		fprintf(stderr, "hilo no cambiado");
		return;
	}
	int ct = current_thread;
	int next_ct = nextCT();
	if(gettimeofday(&ttable[ct].s_time, NULL)){
		err(1, "gettimeofday");
	}
	ttable[ct].state = READY;
	ttable[next_ct].state = RUNNING;
	swapcontext(&ttable[ct].u_c, &ttable[next_ct].u_c);
	fprintf(stderr, "hilo cambiado");
}

int curidthread(void) {
	fprintf(stderr, "obteniendo id...");
	return ttable[current_thread].id;
}
