package main

import "fmt"

var marks = map[string]int{
	"maths":     10,
	"physics":   20,
	"english":   30,
	"chemistry": 40,
	"biology":   50,
}

func main() {
	select_subject := userInput()
	subject_marks := getsubjectMarks(marks, select_subject)
	fmt.Println(select_subject, "marks is:", subject_marks)
	printMap(marks)
	println("======================================")
	iterateOverMaps(marks)
	updateMarks(marks, "maths", 100)
	println("======================================")
	deleteKeyFromMap(marks, select_subject)
	iterateOverMaps(marks)
}

//* print the map
func printMap(marks map[string]int) {
	fmt.Println(marks)
}

//* get marks of a subject
func getsubjectMarks(marks map[string]int, subject string) int {
	return marks[subject]
}

//* get user input
func userInput() string {
	var subject string
	fmt.Println("Enter subject name: ")
	fmt.Scanln(&subject)
	is_exist := checkSubjectExists(marks, subject)
	if is_exist == false {
		fmt.Println("Enter valid subject name")
		userInput()
	}
	println("Provided subject is:" + " " + subject)
	return subject
}

//* iteracte over maps
func iterateOverMaps(marks map[string]int) {
	for key, value := range marks {
		fmt.Println("marks of", key, "is", value)
	}
}

//* update marks of a subject
func updateMarks(marks map[string]int, subject string, new_marks int) {
	marks[subject] = new_marks
}

//* check if subject is present in marks map
func checkSubjectExists(marks map[string]int, subject string) bool {
	if _, ok := marks[subject]; ok {
		fmt.Println("subject exists")
		return true
	} else {
		fmt.Println("subject does not exists")
		return false
	}
}

//* delete a subject from marks
func deleteKeyFromMap(marks map[string]int, subject string) {
	delete(marks, subject)
}
