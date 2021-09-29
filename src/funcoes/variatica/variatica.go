package main

import "fmt"

var (
	students = map[string][]float64{
		"Rafael": {7.1, 7.2, 8.5, 6.0},
		"Bruno":  {3.1, 7.2, 4.5, 6.0},
		"José":   {7.1, 7.2, 8.5, 6.0},
	}
)

func main() {
	fmt.Printf("Média: %.1f\n", media(9.0, 8.4, 6.4, 3.2))
	verifyApproved(students)
	fmt.Println(students)
}

func verifyApproved(alunos map[string][]float64) {
	for aluno, medias := range alunos {
		media := media(medias...)
		fmt.Printf("Aluno: %s - Média: %.2f\n", aluno, media)
	}
}

func media(notes ...float64) float64 {
	total := 0.0
	for _, num := range notes {
		total += num
	}
	return total / float64(len(notes))
}
