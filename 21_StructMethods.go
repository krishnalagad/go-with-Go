package main

import "fmt"

type Student struct {
	name   string
	grades []int
	age    int
}

// setters and getters for struct: Student
func (s Student) setName(name string) {
	s.name = name
}
func (s Student) getName() string {
	return s.name
}
func (s Student) setGrades(grades []int) {
	s.grades = grades
}
func (s Student) getGrades() []int {
	return s.grades
}
func (s Student) setAge(age int) {
	s.age = age
}
func (s Student) getAge() int {
	return s.age
}

func main() {
	// setting values to object in constructor
	s1 := Student{"Krishna", []int{60, 56, 54}, 22}
	fmt.Println(s1)

	// using getters and setters
	s2 := Student{}
	s2.setName("Krishna")
	s2.setGrades([]int{55, 67, 45})
	s2.setAge(21)
	fmt.Println(s2)
}