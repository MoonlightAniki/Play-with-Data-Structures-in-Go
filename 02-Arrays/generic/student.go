package generic

import "fmt"

type Student struct {
	name  string
	score int
}

func GetStudent(name string, score int) (student *Student) {
	student = &Student{
		name:  name,
		score: score,
	}
	return
}

func (student *Student) String() string {
	return fmt.Sprintf("Student(name:%s, score:%d)", student.name, student.score)
}
