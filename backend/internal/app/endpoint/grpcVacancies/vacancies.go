package grpcVacancies

import (
	"context"
	"errors"
	"fmt"
	pb "moscowhack/gen/go/vacancies"
)

type Vacancies interface {
	GetVacanciesService(ctx context.Context) (map[string]*pb.VacanciesItem, error)
	GetVacanciesByIdService(ctx context.Context, id int32) (map[string]*pb.VacanciesIdItem, error)
	GetVacanciesByFilterService(ctx context.Context, vacanciesMap map[string]string) (map[string]*pb.VacanciesItem, error)
	AddVacanciesService(ctx context.Context, vacanciesMap map[string]string) (int32, error)
	//DelVacanciesService(ctx context.Context, newsID int32) error
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

func (e *Endpoint) GetVacancies(ctx context.Context, req *pb.VacanciesRequest) (*pb.VacanciesResponse, error) {
	vacanciesData, err := e.Vacancies.GetVacanciesService(ctx)
	if err != nil {
		return &pb.VacanciesResponse{}, err
	}

	return &pb.VacanciesResponse{Vacancies: vacanciesData}, nil
}

func (e *Endpoint) GetVacanciesById(ctx context.Context, req *pb.VacanciesRequest) (*pb.VacanciesResponse, error) {
	if req.Id == 0 {
		return nil, errors.New("id вакансии не указан")
	}

	newsData, err := e.Vacancies.GetVacanciesByIdService(ctx, req.Id)
	if err != nil {
		return &pb.VacanciesResponse{}, err
	}

	return &pb.VacanciesResponse{Vacancies: newsData}, nil
}

func (e *Endpoint) GetVacanciesByFilter(ctx context.Context, req *pb.VacanciesFilterRequest) (*pb.VacanciesResponse, error) {
	vacanciesMap := map[string]string{
		"departmentCompany":   req.DepartmentCompany,
		"categoryVacancies":   req.CategoryVacancies,
		"experienceStartYear": string(req.ExperienceStartYear),
		"experienceEndYear":   string(req.ExperienceEndYear),
		"educationId":         string(req.EducationId),
		"salary":              string(req.Salary),
		"workHoursPerDay":     string(req.WorkHoursPerDay),
		"workSchedule":        req.WorkSchedule,
		"salaryTaxIncluded":   fmt.Sprint(req.SalaryTaxIncluded),
		"geolocationCompany":  req.GeolocationCompany,
	}

	newsData, err := e.Vacancies.GetVacanciesByFilterService(ctx, vacanciesMap)
	if err != nil {
		return &pb.VacanciesResponse{}, err
	}

	return &pb.VacanciesResponse{Vacancies: newsData}, nil
}

func (e *Endpoint) AddVacancies(ctx context.Context, req *pb.VacanciesIdItem) (*pb.ChangeVacanciesResponse, error) {
	vacanciesMap := map[string]string{
		"id":                 string(req.Id),
		"name":               req.Name,
		"departmentCompany":  req.DepartmentCompany,
		"description":        req.Description,
		"categoryVacancies":  req.CategoryVacancies,
		"experienceYears":    string(req.ExperienceYears),
		"educationId":        string(req.EducationId),
		"workMode":           fmt.Sprint(req.WorkMode),
		"salary":             string(req.Salary),
		"workHoursPerDay":    string(req.WorkHoursPerDay),
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
		return &pb.ChangeVacanciesResponse{Err: error.Error(err)}, err
	}

	return &pb.ChangeVacanciesResponse{Id: id}, nil
}

/*func (e *Endpoint) DelVacancies(ctx context.Context, req *pb.VacanciesRequest) (*pb.ChangeVacanciesResponse, error) {
	if req.Id == 0 {
		return nil, errors.New("id новости не указан")
	}

	err := e.Vacancies.DelVacanciesService(ctx, req.Id)
	if err != nil {
		return &pb.ChangeVacanciesResponse{Err: error.Error(err)}, err
	}

	return &pb.ChangeVacanciesResponse{Err: ""}, nil
}*/
