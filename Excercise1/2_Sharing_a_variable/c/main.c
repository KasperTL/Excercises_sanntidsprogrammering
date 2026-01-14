#include <pthread.h>
#include <stdio.h>



int i = 0;

void* incrementingThreadFunction(){
    // TODO: increment i 1_000_000 times
    for (int j = 0; j < 1000000; j++) {
        i++;
    };
    return NULL;
}

void* decrementingThreadFunction(){
    // TODO: decrement i 1_000_000 times
    for (int j = 0; j < 1000000; j++) {
        i--;
    };
    return NULL;
};



int main() {
    pthread_t thread_1, thread_2;

    pthread_create(&thread_1, NULL, incrementingThreadFunction, NULL);
    pthread_create(&thread_2, NULL, decrementingThreadFunction, NULL);

    pthread_join(thread_1, NULL);
    pthread_join(thread_2, NULL);

    printf("The magic number is: %d\n", i);

};