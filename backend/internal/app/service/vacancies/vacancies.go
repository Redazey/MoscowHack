package vacancies

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	pb "moscowhack/gen/go/vacancies"
	"moscowhack/pkg/cache"
	"moscowhack/pkg/db"
	"strconv"
	"strings"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) GetVacanciesService() (map[string]*pb.GetVacanciesItem, error) {
	vacanciesMap, err := getVacanciesFromCache("vacancies")
	if err != nil {
		return nil, err
	}
	if vacanciesMap != nil {
		return convertVacanciesMap(vacanciesMap), nil
	}

	// Данных нет
	query := `
		SELECT 
			v.id AS "vacancyId",
			v.name AS "vacancyName",
			v."departmentCompany" as "departmentCompany",
			wc."workMode",
			wc.salary,
			wc."salaryTaxIncluded",
			v."geolocationCompany"
		FROM 
			vacancies v
		JOIN 
			"WorkingConditions" wc ON v."workingConditionsID" = wc.id
		GROUP BY 
			v.id, wc."workMode", wc.salary, wc."workHoursPerDay", wc."workSchedule", wc."salaryTaxIncluded";`
	vacanciesMap, err = fetchVacanciesFromDB(query)
	if err != nil {
		return nil, err
	}

	err = cache.SaveCache("vacancies", vacanciesMap)
	if err != nil {
		return nil, err
	}

	return convertVacanciesMap(vacanciesMap), nil
}

func (s *Service) GetVacanciesByIdService(id int32) (*pb.GetVacanciesByIdResponse, error) {
	vacanciesMap, err := getVacanciesFromCache("vacancies_" + fmt.Sprint(id))
	if err != nil {
		return nil, err
	}
	if vacanciesMap != nil {
		return convertVacanciesByIdMap(vacanciesMap["0"]), nil
	}

	// Данных нет
	query := `
		SELECT 
			v.name AS "vacancyName",
			v."departmentCompany",
			v.description,
			cv.name AS "categoryName",
			r."experienceYears",
			e.name AS "educationName",
			wc."workMode",
			wc.salary,
			wc."workHoursPerDay",
			wc."workSchedule",
			wc."salaryTaxIncluded",
			v."geolocationCompany",
			array_agg(DISTINCT s.name) AS skills,
			array_agg(DISTINCT CASE WHEN st.type = 1 THEN st.name END) AS backend_stack,
			array_agg(DISTINCT CASE WHEN st.type = 2 THEN st.name END) AS frontend_stack,
			array_agg(DISTINCT CASE WHEN st.type = 3 THEN st.name END) AS database_stack
		FROM 
			vacancies v
		JOIN 
			"categoriesVacancies" cv ON v."categoryVacanciesID" = cv.id
		JOIN 
			requirements r ON v."requirementsID" = r.id
		JOIN 
			educations e ON r."educationID" = e.id
		JOIN 
			"WorkingConditions" wc ON v."workingConditionsID" = wc.id
		LEFT JOIN 
			"requirementsSkills" rs ON rs."requirementsID" = r.id
		LEFT JOIN 
			skills s ON rs."skillsID" = s.id
		LEFT JOIN 
			"vacanciesStack" vs ON vs."vacancyId" = v.id
		LEFT JOIN 
			stack st ON vs."stackId" = st.id
		WHERE 
			v.id = $1
		GROUP BY 
			v.id, cv.name, r."experienceYears", e.name, wc."workMode", wc.salary, wc."workHoursPerDay", wc."workSchedule", wc."salaryTaxIncluded";`
	vacanciesMap, err = fetchVacanciesByIdFromDB(query, id)
	if err != nil {
		return nil, err
	}

	err = cache.SaveCache("vacancies_"+fmt.Sprint(id), vacanciesMap)
	if err != nil {
		return nil, err
	}

	return convertVacanciesByIdMap(vacanciesMap["0"]), nil
}

func (s *Service) GetVacanciesByFilterService(vacanciesFilter map[string]string) (map[string]*pb.GetVacanciesItem, error) {
	_, hashMap := cache.ConvertMap(vacanciesFilter)

	vacanciesMap, err := getVacanciesFromCache("vacancies_filter_" + hashMap)
	if err != nil {
		return nil, err
	}
	if vacanciesMap != nil {
		return convertVacanciesMap(vacanciesMap), nil
	}

	// Данных нет
	baseQuery := `
		SELECT 
			v.id AS "vacancyId",
			v.name AS "vacancyName",
			v."departmentCompany",
			v.description,
			cv.name AS "categoryName",
			r."experienceYears",
			e.name AS "educationName",
			e."placeEducation",
			wc."workMode",
			wc.salary,
			wc."workHoursPerDay",
			wc."workSchedule",
			wc."salaryTaxIncluded",
			v."geolocationCompany",
			array_agg(DISTINCT s.name) AS skills,
			array_agg(DISTINCT CASE WHEN st.type = 1 THEN st.name END) AS backend_stack,
			array_agg(DISTINCT CASE WHEN st.type = 2 THEN st.name END) AS frontend_stack,
			array_agg(DISTINCT CASE WHEN st.type = 3 THEN st.name END) AS database_stack
		FROM 
			vacancies v
		JOIN 
			"categoriesVacancies" cv ON v."categoryVacanciesID" = cv.id
		JOIN 
			requirements r ON v."requirementsID" = r.id
		JOIN 
			educations e ON r."educationID" = e.id
		JOIN 
			"WorkingConditions" wc ON v."workingConditionsID" = wc.id
		LEFT JOIN 
			"requirementsSkills" rs ON rs."requirementsID" = r.id
		LEFT JOIN 
			skills s ON rs."skillsID" = s.id
		LEFT JOIN 
			"vacanciesStack" vs ON vs."vacancyId" = v.id
		LEFT JOIN 
			stack st ON vs."stackId" = st.id
		WHERE `

	// Определяем условия фильтрации
	filterConditions := map[string]string{
		"departmentCompany":   `v."departmentCompany" = $1`,
		"categoryVacancies":   `cv.name = $2`,
		"experienceStartYear": `r."experienceYears" >= $3`,
		"experienceEndYear":   `r."experienceYears" <= $4`,
		"educationId":         `e.id = $5`,
		"salary":              `wc.salary >= $6`,
		"workHoursPerDay":     `wc."workHoursPerDay" = $7`,
		"workSchedule":        `wc."workSchedule" = $8`,
		"salaryTaxIncluded":   `wc."salaryTaxIncluded" = $9`,
		"geolocationCompany":  `v."geolocationCompany" = $10`,
	}

	// Формируем условия фильтрации
	var filters []string
	var values []interface{}
	for key, query := range filterConditions {
		if value, ok := vacanciesFilter[key]; ok && value != "" {
			filters = append(filters, query)
			values = append(values, value)
		}
	}

	// Добавляем условия фильтрации к базовому запросу
	if len(filters) > 0 {
		baseQuery += strings.Join(filters, " AND ")
	} else {
		// Если фильтры не заданы, добавляем "true", чтобы избежать ошибки синтаксиса
		baseQuery += "true"
	}

	baseQuery += `
		GROUP BY 
			v.id, cv.name, r."experienceYears", e.name, e."placeEducation", wc."workMode", wc.salary, wc."workHoursPerDay", wc."workSchedule", wc."salaryTaxIncluded";`

	vacanciesMap, err = fetchVacanciesFromDB(baseQuery, values...)
	if err != nil {
		return nil, err
	}

	err = cache.SaveCache("vacancies_categories_"+hashMap, vacanciesMap)
	if err != nil {
		return nil, err
	}

	return convertVacanciesMap(vacanciesMap), nil
}

func (s *Service) AddVacanciesService(vacanciesMap map[string]string) (int32, error) {
	/*t, err := time.Parse("2006-01-02 15:04:05", datetime)
	if err != nil {
		return 0, err
	}*/

	// Начало транзакции
	tx, err := db.Conn.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	var requirementsID int
	err = tx.QueryRow(`INSERT INTO requirements ("educationID", "experienceYears") 
		VALUES ($1, $2) 
		RETURNING id`, vacanciesMap["educationId"], vacanciesMap["experienceYears"]).Scan(&requirementsID)
	if err != nil {
		return 0, err
	}

	var workingConditionsID int
	err = tx.QueryRow(`INSERT INTO "WorkingConditions" ("workMode", salary, "workHoursPerDay", "workSchedule", "salaryTaxIncluded") 
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING id`, vacanciesMap["workMode"], vacanciesMap["salary"], vacanciesMap["workHoursPerDay"], vacanciesMap["workSchedule"], vacanciesMap["salaryTaxIncluded"]).Scan(&workingConditionsID)
	if err != nil {
		return 0, err
	}

	var vacancyId int32
	err = tx.QueryRow(`INSERT INTO vacancies (name, "departmentCompany", description, "categoryVacanciesID", "requirementsID", "workingConditionsID", "geolocationCompany") 
		VALUES ($1, $2, $3, $4, $5, $6, $7) 
		RETURNING id`, vacanciesMap["name"], vacanciesMap["departmentCompany"], vacanciesMap["description"], vacanciesMap["categoryVacancies"], requirementsID, workingConditionsID, vacanciesMap["geolocationCompany"]).Scan(&vacancyId)
	if err != nil {
		return 0, err
	}

	if vacanciesMap["backendStack"] != "" {
		backendStack := strings.Split(vacanciesMap["backendStack"], ",")
		for _, item := range backendStack {
			insertQuery := `INSERT INTO "vacanciesStack" ("vacancyId", "stackId") VALUES ($1, $2)`
			_, err = tx.Exec(insertQuery, vacancyId, item)
			if err != nil {
				return 0, err
			}
		}
	}

	if vacanciesMap["frontendStack"] != "" {
		frontendStack := strings.Split(vacanciesMap["frontendStack"], ",")
		for _, item := range frontendStack {
			insertQuery := `INSERT INTO "vacanciesStack" ("vacancyId", "stackId") VALUES ($1, $2)`
			_, err = tx.Exec(insertQuery, vacancyId, item)
			if err != nil {
				return 0, err
			}
		}
	}

	if vacanciesMap["databaseStack"] != "" {
		databaseStack := strings.Split(vacanciesMap["databaseStack"], ",")
		for _, item := range databaseStack {
			insertQuery := `INSERT INTO "vacanciesStack" ("vacancyId", "stackId") VALUES ($1, $2)`
			_, err = tx.Exec(insertQuery, vacancyId, item)
			if err != nil {
				return 0, err
			}
		}
	}

	if errTx := tx.Commit(); errTx != nil {
		return 0, errTx
	}

	return vacancyId, nil
}

func (s *Service) DelVacanciesService(newsID int32) error {
	// Начало транзакции
	tx, err := db.Conn.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Удаление связей новости с категориями
	_, err = tx.Exec("DELETE FROM \"categoriesNews\" WHERE \"newsID\" = $1", newsID)
	if err != nil {
		return err
	}

	// Удаление самой новости
	_, err = tx.Exec("DELETE FROM news WHERE id = $1", newsID)
	if err != nil {
		return err
	}

	// Фиксация транзакции
	if errTx := tx.Commit(); errTx != nil {
		return errTx
	}

	return nil
}

func getVacanciesFromCache(cacheKey string) (map[string]map[string]interface{}, error) {
	if exists, err := cache.IsExistInCache(cacheKey); err != nil {
		return nil, err
	} else if !exists {
		return nil, nil
	}

	vacanciesMap, err := cache.ReadCache(cacheKey)
	if err != nil {
		return nil, err
	}

	if _, notFound := vacanciesMap["notFound"]; notFound {
		return nil, status.Error(codes.NotFound, "нет значений в БД")
	}

	return vacanciesMap, nil
}

func convertVacanciesMap(VacanciesMap map[string]map[string]interface{}) map[string]*pb.GetVacanciesItem {
	contentVacanciesMap := make(map[string]*pb.GetVacanciesItem)
	for _, data := range VacanciesMap {
		id, err := strconv.ParseInt(strings.TrimSpace(data["id"].(string)), 10, 32)
		if err != nil {
			log.Fatalf("Error converting string to int32: %v", err)
		}

		content := &pb.GetVacanciesItem{
			Id:                 int32(id),
			Name:               data["name"].(string),
			DepartmentCompany:  data["departmentCompany"].(string),
			Description:        "",
			CategoryVacancies:  "",
			Requirements:       "",
			WorkingConditions:  "",
			GeolocationCompany: data["geolocationCompany"].(string),
		}
		contentVacanciesMap[data["id"].(string)] = content
	}
	return contentVacanciesMap
}

func convertVacanciesByIdMap(data map[string]interface{}) *pb.GetVacanciesByIdResponse {
	vacancyContent := &pb.GetVacanciesByIdResponse{
		Name:               data["vacancyName"].(string),
		DepartmentCompany:  data["departmentCompany"].(string),
		Description:        data["description"].(string),
		CategoryVacancies:  data["categoryName"].(string),
		ExperienceYears:    data["experienceYears"].(int32),
		EducationName:      data["educationName"].(string),
		WorkMode:           data["workMode"].(bool),
		Salary:             data["salary"].(int32),
		WorkHoursPerDay:    data["workHoursPerDay"].(int32),
		WorkSchedule:       data["workSchedule"].(string),
		SalaryTaxIncluded:  data["salaryTaxIncluded"].(bool),
		GeolocationCompany: data["geolocationCompany"].(string),
		Skills:             data["skills"].(string),
		BackendStack:       data["backendStack"].(string),
		FrontendStack:      data["frontendStack"].(string),
		DatabaseStack:      data["databaseStack"].(string),
	}

	return vacancyContent
}

func fetchVacanciesFromDB(query string, args ...interface{}) (map[string]map[string]interface{}, error) {
	rows, err := db.Conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	vacanciesMap := make(map[string]map[string]interface{})
	found := false
	for rows.Next() {
		var id int
		var name, departmentCompany, geolocationCompany string
		var salary float64
		var workMode, salaryTaxIncluded bool

		err := rows.Scan(
			&id,
			&name,
			&departmentCompany,
			&workMode,
			&salary,
			&salaryTaxIncluded,
			&geolocationCompany,
		)
		if err != nil {
			log.Fatal(err)
		}

		vacanciesMap[fmt.Sprint(id)] = map[string]interface{}{
			"id":                 fmt.Sprint(id),
			"name":               name,
			"departmentCompany":  departmentCompany,
			"workMode":           workMode,
			"salary":             fmt.Sprint(salary),
			"salaryTaxIncluded":  salaryTaxIncluded,
			"geolocationCompany": geolocationCompany,
		}

		found = true
	}

	if !found {
		vacanciesMap["notFound"] = map[string]interface{}{"message": "нет значений в БД"}
		return nil, status.Error(codes.NotFound, "нет значений в БД")
	}

	return vacanciesMap, nil
}

func fetchVacanciesByIdFromDB(query string, args ...interface{}) (map[string]map[string]interface{}, error) {
	rows, err := db.Conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vacancyName, departmentCompany, description, categoryName, educationName, workSchedule, geolocationCompany string
	var experienceYears, workHoursPerDay, salary int32
	var salaryTaxIncluded, workMode bool
	var skills, backendStack, frontendStack, databaseStack string

	vacanciesMap := make(map[string]map[string]interface{})

	found := false
	if rows.Next() {
		err := rows.Scan(
			&vacancyName,
			&departmentCompany,
			&description,
			&categoryName,
			&experienceYears,
			&educationName,
			&workMode,
			&salary,
			&workHoursPerDay,
			&workSchedule,
			&salaryTaxIncluded,
			&geolocationCompany,
			&skills,
			&backendStack,
			&frontendStack,
			&databaseStack,
		)
		if err != nil {
			log.Fatal(err)
		}

		found = true
	}

	if !found {
		vacanciesMap["notFound"] = map[string]interface{}{"message": "нет значений в БД"}
		return nil, status.Error(codes.NotFound, "нет значений в БД")
	}

	vacanciesMap["0"] = map[string]interface{}{
		"vacancyName":        vacancyName,
		"departmentCompany":  departmentCompany,
		"description":        description,
		"categoryName":       categoryName,
		"experienceYears":    experienceYears,
		"educationName":      educationName,
		"workMode":           workMode,
		"salary":             salary,
		"workHoursPerDay":    workHoursPerDay,
		"workSchedule":       workSchedule,
		"salaryTaxIncluded":  salaryTaxIncluded,
		"geolocationCompany": geolocationCompany,
		"skills":             skills,
		"backendStack":       backendStack,
		"frontendStack":      frontendStack,
		"databaseStack":      databaseStack,
	}

	return vacanciesMap, nil
}
