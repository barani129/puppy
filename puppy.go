package puppy


import "fmt"

func Bark() string {
	return "woof! woof!"
}

func Barks() string {
	return "woofs! woofs!"
}

func main(){
	s := Bark()
	v := Barks()
	fmt.Prinln(s)
	fmt.Println(v)
}	
