package model

import time "time"

type Schedule struct {
	ID               uint
	ClassroomID      int
	ProductID        int
	StudentID        int
	TeacherID        int
	OpenclassID      int
	GradeID          int
	No               int
	Status           int
	Type             int
	StartTime        int
	StartAt          time.Time
	EndAt            time.Time
	LeaveFlag        int
	SmsFlag          int
	SendsmsAt        time.Time
	CommentID        int
	CommentTeacherID int
	MonitorFlag      int
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
