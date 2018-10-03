package main

import (
	"strconv"
	"syscall/js"
	"time"
)

func add(i []js.Value) {
	int1, int2 := takeValue(i)
	js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", int1+int2)
}

func subtract(i []js.Value) {
	int1, int2 := takeValue(i)
	js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", int1-int2)
}

func trigger() {
	for {
		js.Global().Get("document").Call("getElementById", "timer").Set("textContent", time.Now().Format("2006/01/02 15:04:05"))
		time.Sleep(1 * time.Second)
	}
}

func takeValue(i []js.Value) (int, int) {
	value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	value2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()
	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)
	return int1, int2
}

func registerCallbacks() {
	js.Global().Set("add", js.NewCallback(add))
	js.Global().Set("subtract", js.NewCallback(subtract))
}

func main() {
	c := make(chan struct{}, 0)

	println("WASM Go Initialized")
	// register functions
	registerCallbacks()
	go trigger()
	<-c
}
