package puppy


import "fmt"

func Bark() {
	return "woof! woof!"
}

func Barks() {
	return "woofs! woofs!"
}

func main(){
	s := Bark()
	v := Barks()
	fmt.Prinln(s)
	fmt.Println(v)
}	
