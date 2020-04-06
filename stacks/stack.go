// stack implementation

package main

import "fmt"
import "errors"

type Stack struct{
    slice []int
}

func main(){
    var s* Stack = new(Stack)
    s.push(1)
    s.push(1)
    s.push(9)
    fmt.Println(s)
    fmt.Println(s.pop())
    fmt.Println(s.peek())
    fmt.Println(s.pop())
    fmt.Println(s.pop())
    //error!
    s.push(9)
    fmt.Println(s.pop())
}

func (s* Stack) String() string{
    return fmt.Sprint(s.slice)
}

// returns last in the stack and remove it
func (s* Stack) pop()(int,error){
    if (len(s.slice)>0){
	ret := s.slice[len(s.slice)-1]
	s.slice = s.slice[0:len(s.slice)-1]
	return ret,nil
    } else {
	return -1,errors.New("Empty stack!")
    }
}

// returns last in the stack BUT NOT remove it
func (s* Stack) peek()(int,error){
    if (len(s.slice)>0){
	ret := s.slice[len(s.slice)-1]
	return ret,nil
    } else {
	return -1,errors.New("Empty stack!")
    }
}

// add i into the stack
func (s* Stack) push(i int){
    s.slice = append(s.slice,i)
}

