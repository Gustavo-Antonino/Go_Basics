package main

import (
	"time"
	"fmt"
)

// Channel (canal) é a forma de comunucação entre as goroutines
// O canal é um tipo de dado que pode ser usado para passar dados entre as goroutines
//

func doisTresQuatroVezes(base int, c chan int){
	time.Sleep(time.Second)
	c  <- 2 * base // enviandos dados para o canal 

	time.Sleep((time.Second))
	c <- 3 * base 

	time.Sleep(3 * time.Second)
	c <- 4 * base 
}

func main(){
	c := make(chan int)
	go doisTresQuatroVezes(2, c)
	fmt.Println("A")

	a, b := <-c, <-c // recebendo os dados do canal 
	fmt.Println("B")
	fmt.Println(a, b)
	

	fmt.Println(<-c)
	
}