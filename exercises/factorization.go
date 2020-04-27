package main

import "fmt"

func main(){
    fmt.Println(Factor([]int{2,3},6))
    fmt.Println(Factor([]int{2,3},15))
}

// take a sorted list of primes and a number and return the factorization of the given number
func Factor(primes []int, number int) []int{
    var result []int
    for _, prime := range primes {
	for (number % prime == 0){
	    result = append(result,prime)
	    number = number/prime
	}
    }
    if number > 1 {
	result = append(result,number)

    }

    return result
}
