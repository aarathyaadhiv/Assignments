package main

import (
	"fmt"
	"sync"
)




func main(){
	mych:=make(chan int )
	wg:= &sync.WaitGroup{}
	wg.Add(2)
	go func(ch chan int,wg *sync.WaitGroup){
		ch<-5
		wg.Done()
	}(mych,wg)
	go func(ch chan int,wg *sync.WaitGroup){
		val:= <-ch
		fmt.Println(val)
		wg.Done()
	}(mych,wg)
	wg.Wait()
	
}