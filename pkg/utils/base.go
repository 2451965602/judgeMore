package utils

const (
	Grade20 = "2020"
	Grade21 = "2021"
	Grade22 = "2022"
	Grade23 = "2023"
	Grade24 = "2024"
	Grade25 = "2025"
)

func IsGradeValid(grade string) bool {
	if grade != Grade20 &&
		grade != Grade21 &&
		grade != Grade22 &&
		grade != Grade23 &&
		grade != Grade24 &&
		grade != Grade25 {
		return false
	}
	return true
}
