package main

import (
	"fmt"
	"time"
	"week6/slideWindow"
)

func main() {
	w := slideWindow.GetSlideWindow(2)
	fmt.Println(w.Visit())
	fmt.Println(w.Visit())
	fmt.Println(w.Visit())
	fmt.Println(w.Visit())
	time.Sleep(time.Second)
	fmt.Println(w.Visit())
	fmt.Println(w.Visit())
	fmt.Println(w.Visit())
	fmt.Println(w.Visit())
	time.Sleep(time.Second)
	fmt.Println(w.Visit())
	fmt.Println(w.Visit())
	fmt.Println(w.Visit())
	fmt.Println(w.Visit())
	fmt.Println("sleep next minute")
	time.Sleep(time.Second * 57)
	fmt.Println(w.Visit())
	fmt.Println(w.Visit())
	fmt.Println(w.Visit())
	fmt.Println(w.Visit())
	time.Sleep(time.Second)
	fmt.Println(w.Visit())
	fmt.Println(w.Visit())
	fmt.Println(w.Visit())
	fmt.Println(w.Visit())
	time.Sleep(time.Second)
	fmt.Println(w.Visit())
	fmt.Println(w.Visit())
	fmt.Println(w.Visit())
	fmt.Println(w.Visit())

}
