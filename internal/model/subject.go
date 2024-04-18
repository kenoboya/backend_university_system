package model

import "time"

type Subject struct {
	ID        int64  `db:"subject_id" json:"subject_id"`
	Name      string `db:"name" json:"name"`
	Specialty Specialty
}

type Lesson struct {
	ID          int64     `db:"lesson_id" json:"lesson_id"`
	LectureRoom string    `db:"lecture_room" json:"lecture_room"`
	TimeStart   time.Time `db:"time_start" json:"time_start"`
	TimeEnd     time.Time `db:"time_end" json:"time_end"`
	LessonType  string    `db:"lesson_type" json:"lesson_type"`
	Subject     Subject
	Group       Group
	Teacher     Teacher
}

type CreateSubjectInput struct {
	SpecialtyID *int64 `db:"specialty_id" json:"specialty_id"`
	Name        string `db:"name" json:"name"`
}
type UpdateStudentInput struct {
	Name string `db:"name" json:"name"`
}

type CreateOrUpdateLessonInput struct {
	TeacherID   int64     `db:"teacher_id" json:"teacher_id"`
	SubjectID   int64     `db:"subject_id" json:"subject_id"`
	GroupID     string    `db:"group_id" json:"group_id"`
	LectureRoom string    `db:"lecture_room" json:"lecture_room"`
	TimeStart   time.Time `db:"time_start" json:"time_start"`
	TimeEnd     time.Time `db:"time_end" json:"time_end"`
	LessonType  string    `db:"lesson_type" json:"lesson_type"`
}
