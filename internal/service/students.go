package service

import (
	"context"
	"test-crud/internal/model"
	"time"
)

type StudentRepository interface {
	Create(ctx context.Context, student model.CreateStudentInput) error
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, id int64, student model.UpdateStudentInput) error
	GetById(ctx context.Context, id int64) (model.Student, error)
	GetAll(ctx context.Context) ([]model.Student, error)
}

type Students struct {
	repo StudentRepository
}

func NewStudents(repo StudentRepository) *Students {
	return &Students{repo}
}

func (st *Students) Create(ctx context.Context, student model.CreateStudentInput) error {
	student.RegisteredAt = time.Now()
	return st.repo.Create(ctx, student)
}
func (st *Students) Delete(ctx context.Context, id int64) error {
	return st.repo.Delete(ctx, id)
}
func (st *Students) Update(ctx context.Context, id int64, student model.UpdateStudentInput) error {
	return st.repo.Update(ctx, id, student)
}
func (st *Students) GetById(ctx context.Context, id int64) (model.Student, error) {
	return st.repo.GetById(ctx, id)
}
func (st *Students) GetAll(ctx context.Context) ([]model.Student, error) {
	return st.repo.GetAll(ctx)
}
