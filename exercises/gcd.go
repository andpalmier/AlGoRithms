package main
import "fmt"

// take two integers and return the greatest common divisor

func main(){
    fmt.Println(gcd(3,6))
    fmt.Println(gcd(256,988))
}

func gcd(x int, y int) int{
    min := x
    max := y
    if (x>y){
	min = y
	max = x
    }
    gcd := 1
    for i:=2;i<=max/2;i++{
	if max%i==0 && min%i==0{
	    gcd = i
	}
    }
    return gcd
}
