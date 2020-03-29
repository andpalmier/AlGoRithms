/*
***DISCLAIMER***
This code is ugly!

Given a list of users with a first and last name, sort them in alphabetical order by their last name, and if two last names are the same, sort them by their first name.
*/

package main
import "fmt"
import "strings"

func main() {
    var users[]string= []string{"Shane Calhoun","John Smith","Bob Smith","Larry Green","Joseph Page","George Costanza","Jerry Seinfeld","John Calhoun"}
    fmt.Println("Not sorted: ", users)
    alphabeticalbubblesort(users)
    fmt.Println("Sorted: ", users)
}

func alphabeticalbubblesort(users []string){
    for i:=0;i<len(users);i++{
	if !sweep(users,i){
	    return
	}
    }
}

// 0 if are the same
// 1 if ordered
// 2 if to be swapped
func order(first string,second string) int{

    min := len(first)
    if len(first)>len(second){
	min = len(second)
    }
    for i:=0;i<min;i++{
	if first[i]>second[i]{
	    return 2
	} else if first[i]<second[i] {
	    return 1
	}
    }

    // if same surname
    return 0
}

func sweep(users []string, step int) bool {

    n := len(users)
    swap := false

    i := 0
    j := 1

    for j < n-step {

	// get user
	userfirst := users[i]
	usersecond := users[j]

	first := strings.SplitN(users[i]," ",2)
	second := strings.SplitN(users[j]," ",2)

	res := order(first[1],second[1])
	// swap surnames if in wrong position
	if res==2 {
	    swap = true
	    fmt.Println("switching surnames: ",userfirst,usersecond)
	    users[i] = usersecond
	    users[j] = userfirst

	// if same surname, check names
	} else if res==0{
	    // swap names if in wrong position
	    res := order(first[0],second[0])
	    if res==2{
		swap = true
		fmt.Println("switching names: ",userfirst,usersecond)
		users[i] = usersecond
		users[j] = userfirst
	    }
	}

	j++
	i++
    }
    return swap
}
