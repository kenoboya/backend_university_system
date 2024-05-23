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
	SubjectID   uint64 `db:"subject_id" json:"subject_id"`
	Name        string `db:"name" json:"name"`
	Semester    []int8 `db:"semester" json:"semester"`
	SubjectType string `db:"subject_type" json:"subject_type"`
}

type Lesson struct {
	LessonID    uint64    `db:"lesson_id" json:"lesson_id"`
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
	Name        string   `db:"name" json:"name"`
	Semester    []uint16 `db:"semester" json:"semester"`
	SubjectType string   `db:"subject_type" json:"subject_type"`
}

type CreateLessonInput struct {
	TeacherID   uint64    `db:"teacher_id" json:"teacher_id"`
	SubjectID   uint64    `db:"subject_id" json:"subject_id"`
	LectureRoom string    `db:"lecture_room" json:"lecture_room"`
	Date        time.Time `db:"date" json:"date"`
	LessonType  string    `db:"lesson_type" json:"lesson_type"`
}

type AttendanceRecord struct {
	StudentBriefInfo
	LessonID uint64 `db:"lesson_id" json:"lesson_id"`
	Status   string `db:"status" json:"status"`
}

type Grade struct {
	StudentBriefInfo
	LessonID uint64 `db:"lesson_id" json:"lesson_id"`
	Grade    uint16 `db:"status" json:"status"`
}
