package structs

import "fmt"

/*
Go's structs are typed collections of fields.
They're useful for groupoing data together to form records
*/

/* This person struct type has name and age fields */
type person struct {
	name string
	age  int
}

/*
newPerson constructs a new person struct with the given name.
you can safely return a pointer to local variable as a local local variable wtill survive the scope of the function.
*/
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 27
	return &p
}

func PrintPeople() {
	defer fmt.Println()

	bob := person{"Bob", 20}
	fmt.Println(bob)

	/* you can name the fields when initializing a struct. */
	alice := person{name: "Alice", age: 30}
	fmt.Println(alice)

	/* Omitted fields will be zero-valued. */
	fred := person{name: "Fred"}
	fmt.Println(fred)

	/* An & prefix yields a pointer to the struct. */
	ann := &person{name: "Ann", age: 40}
	fmt.Println(ann)

	/* It's idiomatic to encapsulate new struct creation in constructor functions. */
	jon := newPerson("Jon")
	fmt.Println(jon)

	/* Access struct fields with a dot. */
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	/* You can also use dots with struct pointers - the pointers are automatically dereferenced. */
	sp := &s
	fmt.Println(sp.age)

	/* Structs are mutable. */
	sp.age = 51
	fmt.Println(s.age)
}
