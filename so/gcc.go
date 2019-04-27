package main

//go:generate gcc -c -fPIC -o hi.o hi.c
//go:generate gcc -shared -o libhi.so hi.o
//go:generate rm hi.o
