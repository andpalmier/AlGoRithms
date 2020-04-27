// tree

package main
import "fmt"
import "errors"
import "strings"
import "strconv"

func main(){
    var t *Node = new (Node)
    t.value = 3
    r := &Node{value:5,}
    t.right = r
    var s *Stack = new(Stack)
    s.push(t)
    fmt.Println(s)
    s.pop()
    fmt.Println(s)
}

// Node struct
type Node struct{
    value int
    left *Node
    right *Node
}

func printNode (n *Node){
    fmt.Println(n.value)
    if(n.left!=nil){
	fmt.Println(n.left.value)
    } else {
	fmt.Println("Left node is empty")
    }
    if(n.right!=nil){
	fmt.Println(n.right.value)
    } else {
	fmt.Println("Right node is empty")
    }
}

// stack struct
type Stack struct{
    slice []*Node
}

func (s* Stack) String() []string{
    if len(s.slice)==0{
	return "empty stack!"
    } else {
	var ret []string
	for i,_ := range s.slice{
	    ret = append(ret,strconv.Itoa(s.slice[i].value))
	}
	return ret
    }
}

func (s* Stack) push(n *Node){
    s.slice = append(s.slice,n)
}

func (s* Stack) pop() (*Node, error){
    if (len(s.slice)>0){
	ret := s.slice[len(s.slice)-1]
	s.slice = s.slice[0:len(s.slice)-1]
	return ret,nil
    } else {
	t := new (Node)
	t.value = -1
	return t,errors.New("empty stack!")
    }
}

