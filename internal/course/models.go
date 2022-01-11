package course

import "signmeup/internal/auth"

type Course struct {
	ID                string                           `json:"id" mapstructure:"id"`
	Title             string                           `json:"title" mapstructure:"title"`
	Code              string                           `json:"code" mapstructure:"code"`
	Term              string                           `json:"term" mapstructure:"term"`
	IsArchived        bool                             `json:"isArchived" mapstructure:"isArchived"`
	CoursePermissions map[string]auth.CoursePermission `json:"coursePermissions" mapstructure:"coursePermissions"`
}

type GetCourseRequest struct {
	CourseID     string     `json:"courseID"`
}

type CreateCourseRequest struct {
	Title     string     `json:"title"`
	Code      string     `json:"code"`
	Term      string     `json:"term"`
	CreatedBy *auth.User `json:"omitempty"`
}

type DeleteCourseRequest struct {
	CourseID     string     `json:"courseID"`
}
