package json_demo

type Dob struct {
	Day   int
	Month int
	Year  int
}

type Person struct {
	Name  string
	Email string
	Dob   Dob
}