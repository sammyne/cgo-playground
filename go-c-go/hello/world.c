#include <stdio.h>

extern void HelloFromGo();

int helloFromC() {
    printf("Hi from C\n");
    //call Go function
    HelloFromGo();
    return 0;
}