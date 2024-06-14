package grpcVacancies

import (
	"context"
	"errors"
	"fmt"
	pb "moscowhack/gen/go/vacancies"
)

type Vacancies interface {
	GetVacanciesService() (map[string]*pb.GetVacanciesItem, error)
	GetVacanciesByIdService(int32) (*pb.GetVacanciesByIdResponse, error)
	GetVacanciesByFilterService(map[string]string) (map[string]*pb.GetVacanciesItem, error)
	AddVacanciesService(map[string]string) (int32, error)
	DelVacanciesService(int32) error
}

type Endpoint struct {
	Vacancies Vacancies
	pb.UnimplementedVacanciesServiceServer
}

func (e *Endpoint) GetVacancies(_ context.Context, _ *pb.GetVacanciesRequest) (*pb.GetVacanciesResponse, error) {
	vacanciesData, err := e.Vacancies.GetVacanciesService()
	if err != nil {
		return &pb.GetVacanciesResponse{}, err
	}

	return &pb.GetVacanciesResponse{Vacancies: vacanciesData}, nil
}

func (e *Endpoint) GetVacanciesById(_ context.Context, req *pb.GetVacanciesByIdRequest) (*pb.GetVacanciesByIdResponse, error) {
	if req.Id == 0 {
		return nil, errors.New("id вакансии не указан")
	}

	newsData, err := e.Vacancies.GetVacanciesByIdService(req.Id)
	if err != nil {
		return &pb.GetVacanciesByIdResponse{}, err
	}

	return newsData, nil
}

func (e *Endpoint) GetVacanciesByFilter(_ context.Context, req *pb.GetFilterVacanciesRequest) (*pb.GetVacanciesResponse, error) {
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

	newsData, err := e.Vacancies.GetVacanciesByFilterService(vacanciesMap)
	if err != nil {
		return &pb.GetVacanciesResponse{}, err
	}

	return &pb.GetVacanciesResponse{Vacancies: newsData}, nil
}

func (e *Endpoint) AddVacancies(_ context.Context, req *pb.AddVacanciesRequest) (*pb.AddVacanciesResponse, error) {
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

	id, err := e.Vacancies.AddVacanciesService(vacanciesMap)
	if err != nil {
		return &pb.AddVacanciesResponse{Err: error.Error(err)}, err
	}

	return &pb.AddVacanciesResponse{Id: id}, nil
}

func (e *Endpoint) DelVacancies(_ context.Context, req *pb.DelVacanciesRequest) (*pb.DelVacanciesResponse, error) {
	if req.Id == 0 {
		return nil, errors.New("id новости не указан")
	}

	err := e.Vacancies.DelVacanciesService(req.Id)
	if err != nil {
		return &pb.DelVacanciesResponse{Err: error.Error(err)}, err
	}

	return &pb.DelVacanciesResponse{Err: ""}, nil
}
