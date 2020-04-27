package main

import "fmt"
// get the index and return the fibonacci number at that index
func main(){
    fmt.Println(IterativeFibonacci(3))
    fmt.Println(IterativeFibonacci(7))
}

func RecursiveFibonacci(n int)int{
    if (n==0){
	return 0
    } else if (n==1) {
	return 1
    } else {
	return RecursiveFibonacci(n-1)+RecursiveFibonacci(n-2)
    }
}

func IterativeFibonacci(n int)int{
    if n <= 1 {
	return n
    }
    x:=0
    y:=1
    i:=2
    var temp int
    for (i<=n){
	temp = x+y
	y = x
	x = temp
	i++
    }
    return x+y
}
