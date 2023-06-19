package main

import "fmt"



func main(){
  var mymap map[string] int = map[string]int {"aarathy":1,"akhil":2}
  fmt.Println(mymap)
  supermap :=map[string]int{"en":1,"fr":2}
  fmt.Println(supermap["en"])
  supermap["fm"]=3
  supermap["fr"]=5
  fmt.Println(supermap)
  for key,value := range supermap{
	fmt.Println(key,"=>",value)
  }
  newmap :=make(map[string]int)
  newmap["ss"]=0
  newmap["pp"]=1
  delete(newmap,"ss")
  fmt.Println(newmap)
  arr:= [10] int {1,2,3,4,5,6,7,8,9,10}
  
  slice1 := append(arr[:4], arr[5:]...)
  fmt.Println(slice1)

}