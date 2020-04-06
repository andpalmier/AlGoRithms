// queue implementation

package main

import "fmt"
import "errors"

type Queue struct{
    // no need to store length with sliced
    slice []int
}

// push i into the queue
func (q* Queue) push(i int){
    q.slice = append(q.slice,i)
}

// pop element from the queue
// error, if queue is empty
func (q* Queue) pop() (int,error){
    if len(q.slice)>0{
	ret := q.slice[0]
	q.slice = q.slice[1:len(q.slice)]
	return ret,nil
    } else {
	return -1,errors.New("queue is empty!")
    }
}

func (q* Queue) String() string {
    return fmt.Sprint(q.slice)
}

func main(){
    var q* Queue = new(Queue)
    q.push(123)
    fmt.Println(q.slice)
    fmt.Println(q.pop())
    fmt.Println(q.pop())
}
