package main

import "fmt"




func main(){
	var word string
	fmt.Println("enter the string")
	fmt.Scanf("%s",&word)
	flag:=0
	for i:=0;i<((len(word))/2)+1;i++{
		if word[i]!=word[len(word)-1-i]{
			flag=1
			break
		}
	}
	if flag==0{
		fmt.Println("word is a palindrome")
	}else{
		fmt.Println("word is not a palindrome")
	}

}