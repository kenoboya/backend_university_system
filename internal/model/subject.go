package model

import "time"

const (
	LessonLecture    = "lecture"
	LessonLaboratory = "laboratory"
	LessonPractical  = "practical"
	LessonSession    = "session"

	SubjectSelected    = "selected"
	SubjectGeneral     = "general"
	SubjectSpecialized = "specialized"

	StatusPresent = "present"
	StatusAbsent  = "absent"
)

type Subject struct {
	SubjectID   int64  `db:"subject_id" json:"subject_id"`
	Name        string `db:"name" json:"name"`
	Semester    []int8 `db:"semester" json:"semester"`
	SubjectType string `db:"subject_type" json:"subject_type"`
}

type Lesson struct {
	LessonID    int64     `db:"lesson_id" json:"lesson_id"`
	LectureRoom string    `db:"lecture_room" json:"lecture_room"`
	Date        time.Time `db:"date" json:"date"`
	LessonType  string    `db:"lesson_type" json:"lesson_type"`
	Subject
	Teacher
}

type CreateSubjectInput struct {
	Name        string `db:"name" json:"name"`
	Semester    []int8 `db:"semester" json:"semester"`
	SubjectType string `db:"subject_type" json:"subject_type"`
}
type UpdateSubjectInput struct {
	Name        string `db:"name" json:"name"`
	Semester    []int8 `db:"semester" json:"semester"`
	SubjectType string `db:"subject_type" json:"subject_type"`
}

type CreateLessonInput struct {
	TeacherID   int64     `db:"teacher_id" json:"teacher_id"`
	SubjectID   int64     `db:"subject_id" json:"subject_id"`
	LectureRoom string    `db:"lecture_room" json:"lecture_room"`
	Date        time.Time `db:"date" json:"date"`
	LessonType  string    `db:"lesson_type" json:"lesson_type"`
}
