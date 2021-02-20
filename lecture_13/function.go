package main
import "fmt"
//example-1
/*
func add(x int, y int) int{
	//body
r := x+y
return r	
}
*/
/*
//example-2
func add(x,y int) int{
	//body
	r := x+y
	return r	
}
*/
/*
//example-3 Name Return Values
func add(x,y int) (r int){
	//body
	r = x+y
	return r	
}
*/
/*
//example-4
func add(x,y int) (r int){
	//body
	r = x+y
	return 	
}
*/
/*
//example-5 Returning Multiple Values
func rectangle(l int, b int) (area int, parameter int) {
	parameter = 2 * (l + b)
	area = l * b
	return // Return statement without specify variable name

}
*/
/*
//example-6 Passing Address to a Function
func update(a *int, t *string) {
	*a = *a + 5      // defrencing pointer address
	*t = *t + " Doe" // defrencing pointer address
	return
}
*/



func main(){
	//x:=add(10, 30)
	//a, p := rectangle(10, 10)
	//fmt.Println(a, p)
	/*
	number := 10
	name:= "Nizam"
	update(&number, &name)
	fmt.Println(number, name)
	*/
//example-7  Anonymous Functions
/*a := func(x, y int) (r int){
 r=x*y
 return
}
fmt.Println(a(10,10))
*/

a := func(x, y int) (r int){
 r=x*y
 return
}(10,10)
fmt.Println(a)
}