package main

import "fmt"

type shape interface{
	area() int
	perimeter() int
}
type rectangle struct{
	l,b int
}
func (r rectangle) area()int{
	return r.l*r.b
}
func (r rectangle) perimeter()int{
	return 2*r.l+2*r.b
}
type square struct{
	a int
}
func(s square) area()int{
	return s.a*s.a
}
func (s square) perimeter()int{
	return 4*s.a
}
func printable (s shape) {
     fmt.Println(s)
	 fmt.Println(s.area())
	 fmt.Println(s.perimeter())
}

func main(){
  r:=rectangle{2,3}
  s:=square{5}
  printable(r)
  printable(s)

}