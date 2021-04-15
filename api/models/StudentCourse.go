package models

type StudentCourse struct {
	ID       uint32 `json:"id"`
	StudentID   uint32 `json:"student_id"`
	CourseID uint32 `json:"course_id"`
}
