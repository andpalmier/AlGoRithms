/*
Given the array of strings that are all animal names, ort the list in alphabetical order with a bubble sort.
*/

package main
import "fmt"



func main() {
    var animals[]string= []string{"dog", "cat", "alligator", "cheetah", "rat", "moose", "cow", "mink", "porcupine", "dung beetle", "camel", "steer", "bat", "hamster", "horse", "colt", "bald eagle", "frog", "rooster"}
    fmt.Println("Not sorted: ", animals)
    alphabeticalbubblesort(animals)
    fmt.Println("Sorted: ", animals)
}

func alphabeticalbubblesort(animals []string){
    for i:=0;i<len(animals);i++{
	if !sweep(animals,i){
	    return
	}
    }
}

func sweep(animals []string, step int) bool {

    n := len(animals)
    swap := false

    i := 0
    j := 1

    for j < n-step {

	first := animals[i]
	second := animals[j]

	// go returns true to "a"<"b"
	if first[0] > second[0]{
	    swap = true
	    fmt.Println("switching: ",first,second)
	    animals[i] = second
	    animals[j] = first
	}

	j++
	i++
    }
    return swap
}
