package psql

import (
	"context"
	"test-crud/internal/model"

	"github.com/jmoiron/sqlx"
)

type NewsRepository struct {
	db *sqlx.DB
}

func NewNewsRepository(db *sqlx.DB) *NewsRepository {
	return &NewsRepository{db}
}

func (r NewsRepository) Create(ctx context.Context, news model.CreateNewsInput) error {
	_, err := r.db.NamedExec("INSERT INTO news(title, author, text, title_image, images, links) VALUES(:title, :author, :text, :title_image, :images, :links)", news)
	if err != nil {
		return err
	}
	return nil
}
func (r NewsRepository) GetAll(ctx context.Context) ([]model.News, error) {
	news := []model.News{}
	err := r.db.Select(&news, "SELECT news.*, people.name, people.surname, people.birth_date, people.photo FROM news JOIN people ON news.author_id = people.person_id")
	if err != nil {
		return news, err
	}
	return news, nil
}
func (r NewsRepository) GetById(ctx context.Context, id int64) (model.News, error) {
	var news model.News
	err := r.db.Get(&news, "SELECT news.*, people.name, people.surname, people.birth_date, people.photo FROM news JOIN people ON news.author_id = people.person_id WHERE news_id = $1", id)
	if err != nil {
		return news, err
	}
	return news, nil
}
func (r NewsRepository) Update(ctx context.Context, id int64, news model.UpdateNewsInput) error {
	_, err := r.db.Exec("UPDATE news SET title = $1, text = $2, title_image = $3, images = $4, links = $5 WHERE news_id = $6",
		news.Title, news.Text, news.TitleImage, news.Images, news.Links, id)
	if err != nil {
		return err
	}
	return nil
}
func (r NewsRepository) Delete(ctx context.Context, id int64) error {
	_, err := r.db.Exec("DELETE FROM news WHERE news_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
