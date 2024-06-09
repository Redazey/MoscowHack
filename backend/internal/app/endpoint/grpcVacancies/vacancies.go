package grpcVacancies

import (
	"context"
	pb "moscowhack/gen/go/vacancies"
)

type Vacancies interface {
	GetVacanciesService(ctx context.Context) (map[string]*pb.VacanciesItem, error)
	//GetVacanciesByIdService(ctx context.Context, id int32) (map[string]*pb.VacanciesItem, error)
	//GetVacanciesByCategoryService(ctx context.Context, categoryId string) (map[string]*pb.VacanciesItem, error)
	//AddVacanciesService(ctx context.Context, title string, text string, datetime string, categories string) (int32, error)
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

/*func (e *Endpoint) GetVacanciesById(ctx context.Context, req *pb.VacanciesRequest) (*pb.VacanciesResponse, error) {
	if req.Id == 0 {
		return nil, errors.New("id новости не указан")
	}

	newsData, err := e.Vacancies.GetVacanciesByIdService(ctx, req.Id)
	if err != nil {
		return &pb.VacanciesResponse{}, err
	}

	return &pb.VacanciesResponse{Vacancies: newsData}, nil
}

func (e *Endpoint) GetVacanciesByCategory(ctx context.Context, req *pb.VacanciesRequest) (*pb.VacanciesResponse, error) {
	if req.Name == "" {
		return nil, errors.New("id категории не указан")
	}

	newsData, err := e.Vacancies.GetVacanciesByCategoryService(ctx, req.CategoryVacancies)
	if err != nil {
		return &pb.VacanciesResponse{}, err
	}

	return &pb.VacanciesResponse{Vacancies: newsData}, nil
}

func (e *Endpoint) AddVacancies(ctx context.Context, req *pb.VacanciesRequest) (*pb.ChangeVacanciesResponse, error) {
	if req.Name == "" {
		return nil, errors.New("заголовок новости не указан")
	}

	id, err := e.Vacancies.AddVacanciesService(ctx, req.Name, req.DepartmentCompany, req.Description, req.CategoryVacancies, req.Requirements, req.WorkingConditions, req.GeolocationCompany)
	if err != nil {
		return &pb.ChangeVacanciesResponse{Err: error.Error(err)}, err
	}

	return &pb.ChangeVacanciesResponse{Id: id}, nil
}

func (e *Endpoint) DelVacancies(ctx context.Context, req *pb.VacanciesRequest) (*pb.ChangeVacanciesResponse, error) {
	if req.Id == 0 {
		return nil, errors.New("id новости не указан")
	}

	err := e.Vacancies.DelVacanciesService(ctx, req.Id)
	if err != nil {
		return &pb.ChangeVacanciesResponse{Err: error.Error(err)}, err
	}

	return &pb.ChangeVacanciesResponse{Err: ""}, nil
}*/
