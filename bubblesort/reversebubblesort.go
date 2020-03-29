/*
Execise from https://www.calhoun.io/lets-learn-algorithms-implementing-bubble-sort/

Given the array of numbers, sort the list in decreasing order with a bubble sort. That is, sort is in reverse of how you normally would.
*/

package main
import "fmt"

func main(){
    var numbers []int=[]int{21, 4, 2, 13, 10, 0, 19, 11, 7, 5, 23, 18, 9, 14, 6, 8, 1, 20, 17, 3, 16, 22, 24, 15, 12}
    fmt.Println("Not reversed: ",numbers)
    reversebubblesort(numbers)
    fmt.Println("Reversed: ",numbers)

}

func reversebubblesort(numbers []int){
    for i:=0;i<len(numbers);i++{

	if !sweep(numbers,i){
	    return
	}
    }
}

func sweep(numbers []int,step int) bool{
    n:= len(numbers)
    swap := false

    i := n-1
    j := n-2

    for j>=0 {
	last := numbers[i]
	secondlast := numbers[j]

	if last>secondlast{
	    fmt.Println("Switching: ",last,secondlast)
	    numbers[j] = last
	    numbers[i] = secondlast
	    swap = true
	}
	j--
	i--
    }
    return swap
}
