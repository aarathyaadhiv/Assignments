package main

import (
	"fmt"
	"time"
)




  func main(){
	fmt.Println(time.Now().Format("01-02-2006 15.04.05 friday"))
	fmt.Println(time.Date(2019,time.December,29,11,25,25,25,time.UTC).Format("02-01-2006 sunday 15.04.05"))

  }