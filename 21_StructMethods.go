package main

import "fmt"

type Student struct {
	name   string
	grades []int
	age    int
}

// setters and getters for struct: Student
func (s *Student) setName(name string) {
	s.name = name
}
func (s Student) getName() string {
	return s.name
}
func (s *Student) setGrades(grades []int) {
	s.grades = grades
}
func (s Student) getGrades() []int {
	return s.grades
}
func (s *Student) setAge(age int) {
	s.age = age
}
func (s Student) getAge() int {
	return s.age
}

// additional methods
func (s Student) getAverageGrades() float32 {
	sum := 0
	for _, v := range s.grades {
		sum += v
	}
	return float32(sum) / float32(len(s.grades))
}

func (s Student) getMaxGrade() int {
	curGrade := 0
	for _, v := range s.grades {
		if v > curGrade {
			curGrade = v
		}
	}
	return curGrade
}

func main() {
	// setting values to object in constructor
	s1 := Student{"Krishna", []int{60, 56, 54}, 22} // parameterized constructor
	s1.setName("Alex")
	fmt.Println(s1)

	// using getters and setters
	s2 := Student{} // default constructor
	s2.setName("Krishna")
	s2.setGrades([]int{55, 67, 45})
	s2.setAge(21)

	fmt.Println(s2.getName(), s2.getGrades(), s2.getAge())

	fmt.Println(s2.getAverageGrades())
	fmt.Println(s2.getMaxGrade())
}
