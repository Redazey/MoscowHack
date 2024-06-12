package grpcVacancies

import (
	"context"
	"errors"
	"fmt"
	pb "moscowhack/gen/go/vacancies"
)

type Vacancies interface {
	GetVacanciesService(ctx context.Context) (map[string]*pb.GetVacanciesItem, error)
	GetVacanciesByIdService(ctx context.Context, id int32) (*pb.GetVacanciesByIdResponse, error)
	GetVacanciesByFilterService(ctx context.Context, vacanciesMap map[string]string) (map[string]*pb.GetVacanciesItem, error)
	AddVacanciesService(ctx context.Context, vacanciesMap map[string]string) (int32, error)
	//DelVacanciesService(ctx context.Context, vacanciesID int32) error
}

type Endpoint struct {
	Vacancies Vacancies
	pb.UnimplementedVacanciesServiceServer
}

func New(vacancies Vacancies) *Endpoint {
	return &Endpoint{
		Vacancies: vacancies,
	}
}

func (e *Endpoint) GetVacancies(ctx context.Context, req *pb.GetVacanciesRequest) (*pb.GetVacanciesResponse, error) {
	vacanciesData, err := e.Vacancies.GetVacanciesService(ctx)
	if err != nil {
		return &pb.GetVacanciesResponse{}, err
	}

	return &pb.GetVacanciesResponse{Vacancies: vacanciesData}, nil
}

func (e *Endpoint) GetVacanciesById(ctx context.Context, req *pb.GetVacanciesByIdRequest) (*pb.GetVacanciesByIdResponse, error) {
	if req.Id == 0 {
		return nil, errors.New("id вакансии не указан")
	}

	newsData, err := e.Vacancies.GetVacanciesByIdService(ctx, req.Id)
	if err != nil {
		return &pb.GetVacanciesByIdResponse{}, err
	}

	return newsData, nil
}

func (e *Endpoint) GetVacanciesByFilter(ctx context.Context, req *pb.GetFilterVacanciesRequest) (*pb.GetVacanciesResponse, error) {
	vacanciesMap := map[string]string{
		"departmentCompany":   req.DepartmentCompany,
		"categoryVacancies":   req.CategoryVacancies,
		"experienceStartYear": fmt.Sprint(req.ExperienceStartYear),
		"experienceEndYear":   fmt.Sprint(req.ExperienceEndYear),
		"educationId":         fmt.Sprint(req.EducationId),
		"salary":              fmt.Sprint(req.Salary),
		"workHoursPerDay":     fmt.Sprint(req.WorkHoursPerDay),
		"workSchedule":        req.WorkSchedule,
		"salaryTaxIncluded":   fmt.Sprint(req.SalaryTaxIncluded),
		"geolocationCompany":  req.GeolocationCompany,
	}

	newsData, err := e.Vacancies.GetVacanciesByFilterService(ctx, vacanciesMap)
	if err != nil {
		return &pb.GetVacanciesResponse{}, err
	}

	return &pb.GetVacanciesResponse{Vacancies: newsData}, nil
}

func (e *Endpoint) AddVacancies(ctx context.Context, req *pb.AddVacanciesRequest) (*pb.AddVacanciesResponse, error) {
	vacanciesMap := map[string]string{
		"name":               req.Name,
		"departmentCompany":  req.DepartmentCompany,
		"description":        req.Description,
		"categoryVacancies":  req.CategoryVacancies,
		"experienceYears":    fmt.Sprint(req.ExperienceYears),
		"educationId":        fmt.Sprint(req.EducationId),
		"workMode":           fmt.Sprint(req.WorkMode),
		"salary":             fmt.Sprint(req.Salary),
		"workHoursPerDay":    fmt.Sprint(req.WorkHoursPerDay),
		"workSchedule":       req.WorkSchedule,
		"salaryTaxIncluded":  fmt.Sprint(req.SalaryTaxIncluded),
		"geolocationCompany": req.GeolocationCompany,
		"skills":             req.Skills,
		"backendStack":       req.BackendStack,
		"frontendStack":      req.FrontendStack,
		"databaseStack":      req.DatabaseStack,
	}

	id, err := e.Vacancies.AddVacanciesService(ctx, vacanciesMap)
	if err != nil {
		return &pb.AddVacanciesResponse{Err: error.Error(err)}, err
	}

	return &pb.AddVacanciesResponse{Id: id}, nil
}

/*func (e *Endpoint) DelVacancies(ctx context.Context, req *pb.DelVacanciesRequest) (*pb.DelVacanciesResponse, error) {
	if req.Id == 0 {
		return nil, errors.New("id новости не указан")
	}

	err := e.Vacancies.DelVacanciesService(ctx, req.Id)
	if err != nil {
		return &pb.DelVacanciesResponse{Err: error.Error(err)}, err
	}

	return &pb.DelVacanciesResponse{Err: ""}, nil
}*/
