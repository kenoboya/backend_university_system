package service

import "test-crud/internal/repository/psql"

type SubjectsService struct {
	repo psql.Subjects
}

func NewSubjectsService(repo psql.Subjects) *SubjectsService {
	return &SubjectsService{repo}
}

type LessonsService struct {
	repo psql.Lessons
}

func NewLessonsService(repo psql.Lessons) *LessonsService {
	return &LessonsService{repo}
}
