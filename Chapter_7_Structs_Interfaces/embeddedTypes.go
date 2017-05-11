package main

import "fmt"

type Person struct {
	Name string
}

// Method for a person to say their name
func (p *Person) Talk(){
	fmt.Println("Hi my name is", p.Name)
}

type Android struct {
	// Android is a person
	Person
	Model string
}

func main(){
	a := new(Android)
	// Android can set the name of the person
	a.Person.Name = "Robo"
	// Android can use Person methods because it has type Person 
	a.Talk()
}
