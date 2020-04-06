// Simple binary search implementation

package main

import "fmt"

func main(){
    var lookingFor int = 13
    var sortedList []int = []int{1,3,4,6,7,9,10,11,13}
    fmt.Println("Looking for", lookingFor, "in: ",sortedList)
    index := binarysearch(sortedList,lookingFor)
    if index >= 0 {
	fmt.Println("Number found at index:",index)
    } else {
	fmt.Println(lookingFor, "not found in:",sortedList)
    }
}

func binarysearch(sortedList []int, lookingFor int) int{

    // check beginning and end of list
    end := len(sortedList)-1
    start := 0

    // while there is at least 1 element in the sublist
    for start<=end {

	// update mid value of the sublist
	mid := start+(end-start)/2
	midval := sortedList[mid]
	fmt.Println("midval is", midval)
	fmt.Println("mid is", mid)
	// number found!
	if midval==lookingFor {
	    return mid+1

	// if mid is more than lookingfor -> check left sublist 
	} else if midval >= lookingFor {
	    end = mid

	// else -> check right sublist
	} else {
	    start = mid +1
	    fmt.Println(start,end)
	}

    }

    // return not found otherwise
    return -1
}

