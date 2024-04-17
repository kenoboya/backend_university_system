package psql

type Repositories struct {
	Students Students
}

func NewRepositories(student Students) *Repositories {
	return &Repositories{student}
}
