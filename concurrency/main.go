package main

func main() {
	ch := make(chan int, 2)
	ch <- 2
}
