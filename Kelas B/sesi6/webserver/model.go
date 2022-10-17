package webserver

type Employee struct {
	ID       int
	Name     string `json:"name"`
	Age      int
	Division []string
}

var emps = []Employee{
	{
		ID:       1,
		Name:     "MNC",
		Age:      10,
		Division: []string{"Media", "Edukasi", "Entertainment"},
	},
}

func GetEmployees() []Employee {
	return emps
}

func CreateEmployee(emp *Employee) []Employee {
	emps = append(emps, *emp)
	return emps
}
