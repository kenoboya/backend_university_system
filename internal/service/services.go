package service

type Services struct {
	Students Students
}

func NewServices(student Students) *Services {
	return &Services{student}
}
