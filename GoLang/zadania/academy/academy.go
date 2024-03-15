package academy

import "math"

type Student struct {
	Name       string
	Grades     []int
	Project    int
	Attendance []bool
}

// AverageGrade returns an average grade given a
// slice containing all grades received during a
// semester, rounded to the nearest integer.
func AverageGrade(grades []int) int {
	n := len(grades)
	if n == 0 {
		return 0
	}

	var sum int
	for i := 0; i < n; i++ {
		sum += grades[i]
	}

	// return int(math.Round(sum/float64(len(grades))));

	average := float64(sum) / float64(n)
	average = math.Round(average)

	return int(average)
}

// AttendancePercentage returns a percentage of class
// attendance, given a slice containing information
// whether a student was present (true) of absent (false).
//
// The percentage of attendance is represented as a
// floating-point number ranging from 0 to 1.
func AttendancePercentage(attendance []bool) float64 {
	n := len(attendance)

	var sum float64
	for i := 0; i < n; i++ {
		if attendance[i] {
			sum++
		}
	}

	return sum / float64(n)
}

// FinalGrade returns a final grade achieved by a student,
// ranging from 1 to 5.
//
// The final grade is calculated as the average of a project grade
// and an average grade from the semester, with adjustments based
// on the student's attendance. The final grade is rounded
// to the nearest integer.

// If the student's attendance is below 80%, the final grade is
// decreased by 1. If the student's attendance is below 60%, average
// grade is 1 or project grade is 1, the final grade is 1.
func FinalGrade(s Student) int {
	if AttendancePercentage(s.Attendance) < 0.6 || s.Project == 1 || AverageGrade(s.Grades) == 1 {
		return 1
	}

	fg := float64(s.Project+AverageGrade(s.Grades)) / 2
	fg = math.Round(fg)

	if AttendancePercentage(s.Attendance) < 0.8 && fg > 1 {
		return int(fg - 1)
	}
	return int(fg)

}

// GradeStudents returns a map of final grades for a given slice of
// Student structs. The key is a student's name and the value is a
// final grade.
func GradeStudents(students []Student) map[string]uint8 {
	gradedStudents := map[string]uint8{}
	for i := 0; i < len(students); i++ {
		gradedStudents[students[i].Name] = uint8(FinalGrade(students[i]))
	}
	return gradedStudents
}
