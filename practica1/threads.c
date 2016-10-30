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

long gettimeinmillis(timeval t) {
	return t.tv_usec / 1000 + t.tv_sec * 1000;
}

void initthreads(void){
	if(getcontext(&ttable[0].u_c))
		err(1, "getcontext initthread\n\r)");
	if(gettimeofday(&ttable[0].s_time, NULL))
		err(1, "gettimeofday init\n\r)");
	ttable[0].id = getId();
	ttable[0].state = RUNNING;
}

int createthread(void (*mainf)(void*), void *arg, int stacksize) {
	int tindex_void = -1;
	for(int i = 0; i < MAX_THREAD; i++){
		if(ttable[i].state == AVILIABLE){
			tindex_void = i;
			break; 
		}
	}
	if(tindex_void == -1 ||
			getcontext(&ttable[tindex_void].u_c)) 		
		return -1;
	
	ttable[tindex_void].id = getId();
	int mallocea = 0;
	if(&ttable[tindex_void].stacksize == NULL) {
		mallocea = 1;
	} else if(ttable[tindex_void].stacksize != stacksize) {
		free(ttable[tindex_void].stack_p);
		mallocea = 1;
	}
	if(mallocea) {
		ttable[tindex_void].stack_p = malloc(stacksize);
	}
	ttable[tindex_void].u_c.uc_stack.ss_sp = ttable[tindex_void].stack_p;
	ttable[tindex_void].u_c.uc_stack.ss_size = ttable[tindex_void].stacksize;
	ttable[tindex_void].u_c.uc_link = &ttable[current_thread].u_c;
	ttable[tindex_void].state = READY;
	makecontext(&ttable[tindex_void].u_c, (void(*))mainf, 1, arg);
	return ttable[tindex_void].id;
} 

void exitsthread(void) {
	ttable[current_thread].state = AVILIABLE;
	yieldthread();
}

int nextCT() {
	char i = 0;
	do {
		if(current_thread + 1 == MAX_THREAD){
			current_thread = 0;
		} else { 
			current_thread ++;
		}
		i++;
		if(i == MAX_THREAD){
			return -1;
		}
	} while(ttable[current_thread].state != READY);
	return current_thread;
}

int quantumSpent(){
	if(ttable[current_thread].state == AVILIABLE){//Cuando el hilo acaba de morir
		return 1;
	}
	timeval t;
	if(gettimeofday(&t, NULL)){
		err(1, "gettimeofday\n\r)");
	}
	long now = gettimeinmillis(t);
	long before = gettimeinmillis(ttable[current_thread].s_time);
	
	return now - before > 200;
}

void yieldthread(void) {	
	if(!quantumSpent()){
		return;
	}
	int ct = current_thread;
	if(ttable[ct].state != AVILIABLE){ //El hilo acaba de morir
		ttable[ct].state = READY;
	}
	int next_ct = nextCT();
	if(next_ct == -1) {	
		err(1, "No hay hilos para continuar la ejecucion");
	}
	if(gettimeofday(&ttable[ct].s_time, NULL)){
		err(1, "gettimeofday\n\r)");
	}
	if(ct == next_ct) {
		return;
	}
	ttable[next_ct].state = RUNNING;
	swapcontext(&ttable[ct].u_c, &ttable[next_ct].u_c);
}

int curidthread(void) {
	return ttable[current_thread].id;
}
