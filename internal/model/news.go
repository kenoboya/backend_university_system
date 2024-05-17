package model

import "time"

type News struct {
	NewsID        int64     `db:"news_id" json:"news_id"`
	Title         string    `db:"title" json:"title"`
	Author        int64     `db:"author_id" json:"author_id"`
	PublishedTime time.Time `db:"published_time" json:"published_time"`
	Text          string    `db:"text" json:"text"`
	TitleImage    string    `db:"title_image" json:"title_image"`
	Images        *[]string `db:"images" json:"images"`
	Links         *[]string `db:"links" json:"links"`
	// categoryID string `db:"" json:"published_time"`
	PersonBriefInfo
}

type CreateNewsInput struct {
	Title      string    `db:"title" json:"title"`
	Author     int64     `db:"person_id" json:"person_id"`
	Text       string    `db:"text" json:"text"`
	TitleImage string    `db:"title_image" json:"title_image"`
	Images     *[]string `db:"images" json:"images"`
	Links      *[]string `db:"links" json:"links"`
}

type UpdateNewsInput struct {
	Title      string    `db:"title" json:"title"`
	Text       string    `db:"text" json:"text"`
	TitleImage string    `db:"title_image" json:"title_image"`
	Images     *[]string `db:"images" json:"images"`
	Links      *[]string `db:"links" json:"links"`
}
