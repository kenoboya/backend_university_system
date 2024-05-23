package service

import (
	"context"
	"test-crud/internal/model"
	"test-crud/internal/repository/psql"
)

type NewsService struct {
	repo psql.News
}

func NewNewsService(repo psql.News) *NewsService {
	return &NewsService{repo}
}

func (s *NewsService) Create(ctx context.Context, news model.CreateNewsInput) error {
	return s.repo.Create(ctx, news)
}
func (s *NewsService) GetList(ctx context.Context) ([]model.News, error) {
	return s.repo.GetAll(ctx)
}
func (s *NewsService) GetNews(ctx context.Context, newsID uint64) (model.News, error) {
	return s.repo.GetById(ctx, newsID)
}
func (s *NewsService) Update(ctx context.Context, newsID uint64, news model.UpdateNewsInput) error {
	return s.repo.Update(ctx, newsID, news)
}
func (s *NewsService) Delete(ctx context.Context, newsID uint64) error {
	return s.repo.Delete(ctx, newsID)
}
