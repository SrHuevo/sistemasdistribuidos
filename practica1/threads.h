#ifndef _THREADS_H
#define _THREADS_H

void initthreads(void);
int createthread(void (*mainf)(void*), void *arg, int stacksize);
void exitsthread(void);
void yieldthread(void);
int curidthread(void);

#endif // _THREADS_H
