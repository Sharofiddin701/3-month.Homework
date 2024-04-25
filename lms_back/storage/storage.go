package storage

import (
	"context"
	"lms_back/api/models"
)

type IStorage interface {
	CloseDB()
	Admin() IAdminStorage
	Branches() IBranchStorage
	Group() IGroupStorage
	Payment() IPaymentStorage
	Student() IStudentStorage
	Teacher() ITeacherStorage
	Schedule() IScheduleStorage
	Task() ITaskStorage
	Lesson() ILessonStorage
	Report() IReportAdminStorage
}

type IAdminStorage interface {
	Create(context.Context, models.Admin) (models.Admin, error)
	GetAll(context.Context, models.GetAllAdminsRequest) (models.GetAllAdminsResponse, error)
	GetByID(context.Context, string) (models.Admin, error)
	Update(context.Context, models.Admin) (models.Admin, error)
	Delete(context.Context, string) error
}

type IBranchStorage interface {
	Create(context.Context, models.Branch) (models.Branch, error)
	GetAll(context.Context, models.GetAllBranchesRequest) (models.GetAllBranchesResponse, error)
	GetByID(context.Context, string) (models.Branch, error)
	Update(context.Context, models.Branch) (models.Branch, error)
	Delete(context.Context, string) error
}

type IGroupStorage interface {
	Create(context.Context, models.Group) (models.Group, error)
	GetAll(context.Context, models.GetAllGroupsRequest) (models.GetAllGroupsResponse, error)
	GetByID(context.Context, string) (models.Group, error)
	Update(context.Context, models.Group) (models.Group, error)
	Delete(context.Context, string) error
}

type IPaymentStorage interface {
	Create(context.Context, models.Payment) (models.Payment, error)
	GetAll(context.Context, models.GetAllPaymentsRequest) (models.GetAllPaymentsResponse, error)
	GetByID(context.Context, string) (models.Payment, error)
	Update(context.Context, models.Payment) (models.Payment, error)
	Delete(context.Context, string) error
}

type IStudentStorage interface {
	Create(context.Context, models.Student) (models.Student, error)
	GetAll(context.Context, models.GetAllStudentsRequest) (models.GetAllStudentsResponse, error)
	GetByID(context.Context, string) (models.Student, error)
	Update(context.Context, models.Student) (models.Student, error)
	Delete(context.Context, string) error
}

type ITeacherStorage interface {
	Create(context.Context, models.Teacher) (models.Teacher, error)
	GetAll(context.Context, models.GetAllTeachersRequest) (models.GetAllTeachersResponse, error)
	GetByID(context.Context, string) (models.Teacher, error)
	Update(context.Context, models.Teacher) (models.Teacher, error)
	Delete(context.Context, string) error
}

type IScheduleStorage interface {
	Create(context.Context, models.Schedule) (models.Schedule, error)
	GetAll(context.Context, models.GetAllSchedulesRequest) (models.GetAllSchedulesResponse, error)
	GetByID(context.Context, string) (models.Schedule, error)
	Update(context.Context, models.Schedule) (models.Schedule, error)
	Delete(context.Context, string) error
}

type ITaskStorage interface {
	Create(context.Context, models.Task) (models.Task, error)
	GetAll(context.Context, models.GetAllTasksRequest) (models.GetAllTasksResponse, error)
	GetByID(context.Context, string) (models.Task, error)
	Update(context.Context, models.Task) (models.Task, error)
	Delete(context.Context, string) error
}

type ILessonStorage interface {
	Create(context.Context, models.Lesson) (models.Lesson, error)
	GetAll(context.Context, models.GetAllLessonsRequest) (models.GetAllLessonsResponse, error)
	GetByID(context.Context, string) (models.Lesson, error)
	Update(context.Context, models.Lesson) (models.Lesson, error)
	Delete(context.Context, string) error
}
type IReportAdminStorage interface {
	Get(context.Context, models.AdminPrimaryKey) (models.ReportAdmin, error)
}
