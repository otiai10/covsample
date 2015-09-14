package sample

import (
	"log"
	"runtime"
)

func Caller() {
	a, f, b, c := runtime.Caller(0)
	log.Println(a, f, b, c)
}
