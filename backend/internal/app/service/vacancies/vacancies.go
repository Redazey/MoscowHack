package vacancies

import (
	"context"
	"database/sql"
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

func (s *Service) GetVacanciesService(ctx context.Context) (map[string]*pb.GetVacanciesItem, error) {
	// Initialize newsSlice
	vacanciesMap := make(map[string]map[string]interface{})

	newsCheck, err := cache.IsExistInCache("vacancies")
	if newsCheck && err == nil {
		vacanciesMap, err = cache.ReadCache("vacancies")
		if err != nil {
			return nil, err
		}

		if _, notFound := vacanciesMap["notFound"]; notFound {
			return nil, status.Error(codes.NotFound, "нет значений в БД")
		}

		vacanciesContentMap := createVacanciesContentMap(vacanciesMap)

		return vacanciesContentMap, nil
	}

	// Данных нет
	rows, err := db.Conn.Query(`
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
			v.id, wc."workMode", wc.salary, wc."workHoursPerDay", wc."workSchedule", wc."salaryTaxIncluded";
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	fmt.Println("1123")

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

		fmt.Println("1123344")

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

	err = cache.SaveCache("vacancies", vacanciesMap)
	if err != nil {
		return nil, err
	}

	vacanciesContentMap := createVacanciesContentMap(vacanciesMap)

	return vacanciesContentMap, nil
}

func (s *Service) GetVacanciesByIdService(ctx context.Context, id int32) (*pb.GetVacanciesByIdResponse, error) {
	// Initialize newsSlice
	vacanciesMap := make(map[string]map[string]interface{})

	vacanciesCheck, err := cache.IsExistInCache("vacancies_" + fmt.Sprint(id))
	if vacanciesCheck && err == nil {
		vacanciesMap, err = cache.ReadCache("vacancies_" + fmt.Sprint(id))
		if err != nil {
			return nil, err
		}

		if _, notFound := vacanciesMap["notFound"]; notFound {
			return nil, status.Error(codes.NotFound, "нет значений в БД")
		}

		vacanciesContent := &pb.GetVacanciesByIdResponse{
			Name:               vacanciesMap[fmt.Sprint(id)]["vacancyName"].(string),
			DepartmentCompany:  vacanciesMap[fmt.Sprint(id)]["departmentCompany"].(string),
			Description:        vacanciesMap[fmt.Sprint(id)]["description"].(string),
			CategoryVacancies:  vacanciesMap[fmt.Sprint(id)]["categoryName"].(string),
			ExperienceYears:    vacanciesMap[fmt.Sprint(id)]["experienceYears"].(int32),
			EducationName:      vacanciesMap[fmt.Sprint(id)]["educationName"].(string),
			WorkMode:           vacanciesMap[fmt.Sprint(id)]["workMode"].(bool),
			Salary:             vacanciesMap[fmt.Sprint(id)]["salary"].(int32),
			WorkHoursPerDay:    vacanciesMap[fmt.Sprint(id)]["workHoursPerDay"].(int32),
			WorkSchedule:       vacanciesMap[fmt.Sprint(id)]["workSchedule"].(string),
			SalaryTaxIncluded:  vacanciesMap[fmt.Sprint(id)]["salaryTaxIncluded"].(bool),
			GeolocationCompany: vacanciesMap[fmt.Sprint(id)]["geolocationCompany"].(string),
			Skills:             vacanciesMap[fmt.Sprint(id)]["skills"].(string),
			BackendStack:       vacanciesMap[fmt.Sprint(id)]["backendStack"].(string),
			FrontendStack:      vacanciesMap[fmt.Sprint(id)]["frontendStack"].(string),
			DatabaseStack:      vacanciesMap[fmt.Sprint(id)]["databaseStack"].(string),
		}

		return vacanciesContent, nil
	}

	// Данных нет
	rows, err := db.Conn.Query(`
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
			v.id, cv.name, r."experienceYears", e.name, wc."workMode", wc.salary, wc."workHoursPerDay", wc."workSchedule", wc."salaryTaxIncluded";
	`, id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	found := false

	var vacancyName, departmentCompany, description, categoryName, educationName, workSchedule, geolocationCompany string
	var experienceYears, workHoursPerDay, salary int32
	var salaryTaxIncluded, workMode bool
	var skills, backendStack, frontendStack, databaseStack string
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

	vacanciesMap[fmt.Sprint(id)] = map[string]interface{}{
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

	err = cache.SaveCache("vacancies_"+fmt.Sprint(id), vacanciesMap)
	if err != nil {
		return nil, err
	}

	vacanciesContent := &pb.GetVacanciesByIdResponse{
		Name:               vacanciesMap[fmt.Sprint(id)]["vacancyName"].(string),
		DepartmentCompany:  vacanciesMap[fmt.Sprint(id)]["departmentCompany"].(string),
		Description:        vacanciesMap[fmt.Sprint(id)]["description"].(string),
		CategoryVacancies:  vacanciesMap[fmt.Sprint(id)]["categoryName"].(string),
		ExperienceYears:    vacanciesMap[fmt.Sprint(id)]["experienceYears"].(int32),
		EducationName:      vacanciesMap[fmt.Sprint(id)]["educationName"].(string),
		WorkMode:           vacanciesMap[fmt.Sprint(id)]["workMode"].(bool),
		Salary:             vacanciesMap[fmt.Sprint(id)]["salary"].(int32),
		WorkHoursPerDay:    vacanciesMap[fmt.Sprint(id)]["workHoursPerDay"].(int32),
		WorkSchedule:       vacanciesMap[fmt.Sprint(id)]["workSchedule"].(string),
		SalaryTaxIncluded:  vacanciesMap[fmt.Sprint(id)]["salaryTaxIncluded"].(bool),
		GeolocationCompany: vacanciesMap[fmt.Sprint(id)]["geolocationCompany"].(string),
		Skills:             vacanciesMap[fmt.Sprint(id)]["skills"].(string),
		BackendStack:       vacanciesMap[fmt.Sprint(id)]["backendStack"].(string),
		FrontendStack:      vacanciesMap[fmt.Sprint(id)]["frontendStack"].(string),
		DatabaseStack:      vacanciesMap[fmt.Sprint(id)]["databaseStack"].(string),
	}

	return vacanciesContent, nil
}

func (s *Service) GetVacanciesByFilterService(ctx context.Context, vacanciesMap map[string]string) (map[string]*pb.GetVacanciesItem, error) {
	// Initialize newsSlice
	vacanciesFilter, hashMap := cache.ConvertMap(vacanciesMap)

	vacanciesCheck, err := cache.IsExistInCache("vacancies_filter_" + hashMap)
	if vacanciesCheck && err == nil {
		vacanciesFilter, err = cache.ReadCache("vacancies_filter_" + hashMap)
		if err != nil {
			return nil, err
		}

		if _, notFound := vacanciesFilter["notFound"]; notFound {
			return nil, status.Error(codes.NotFound, "нет значений в БД")
		}

		vacanciesContentMap := createVacanciesContentMap(vacanciesFilter)

		return vacanciesContentMap, nil
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
			stack st ON vs."stackId" = st.id`

	var filters []string
	var values []interface{}
	i := 1

	// Define a list of filter conditions
	filterConditions := []struct {
		key   string
		query string
	}{
		{"departmentCompany", `v."departmentCompany" = $%d`},
		{"categoryVacancies", `cv.name = $%d`},
		{"experienceStartYear", `r."experienceYears" >= $%d`},
		{"experienceEndYear", `r."experienceYears" <= $%d`},
		{"educationId", `e.id = $%d`},
		{"salary", `wc.salary >= $%d`},
		{"workHoursPerDay", `wc."workHoursPerDay" = $%d`},
		{"workSchedule", `wc."workSchedule" = $%d`},
		{"salaryTaxIncluded", `wc."salaryTaxIncluded" = $%d`},
		{"geolocationCompany", `v."geolocationCompany" = $%d`},
	}

	// Add filters based on parameters
	for _, condition := range filterConditions {
		if value, ok := vacanciesMap[condition.key]; ok && value != "" {
			filters = append(filters, fmt.Sprintf(condition.query, i))
			values = append(values, value)
			i++
		}
	}

	// Append filters to the base query if any
	if len(filters) > 0 {
		baseQuery += " WHERE " + strings.Join(filters, " AND ")
	}

	baseQuery += `
		GROUP BY 
			v.id, cv.name, r."experienceYears", e.name, e."placeEducation", wc."workMode", wc.salary, wc."workHoursPerDay", wc."workSchedule", wc."salaryTaxIncluded";`
	rows, err := db.Conn.Queryx(baseQuery, values...)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	found := false
	for rows.Next() {
		var vacancyId int
		var vacancyName, departmentCompany, description, categoryName, educationName, placeEducation, workSchedule, geolocationCompany string
		var experienceYears, workHoursPerDay int
		var salary float64
		var salaryTaxIncluded bool
		var workMode, skills, backendStack, frontendStack, databaseStack sql.NullString

		err := rows.Scan(
			&vacancyId,
			&vacancyName,
			&departmentCompany,
			&description,
			&categoryName,
			&experienceYears,
			&educationName,
			&placeEducation,
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

		vacanciesFilter[fmt.Sprint(vacancyId)] = map[string]interface{}{
			"vacancyId":          fmt.Sprint(vacancyId),
			"vacancyName":        vacancyName,
			"departmentCompany":  departmentCompany,
			"description":        description,
			"categoryName":       categoryName,
			"experienceYears":    experienceYears,
			"educationName":      educationName,
			"placeEducation":     placeEducation,
			"workMode":           workMode,
			"salary":             fmt.Sprint(salary),
			"workHoursPerDay":    workHoursPerDay,
			"workSchedule":       workSchedule,
			"salaryTaxIncluded":  salaryTaxIncluded,
			"geolocationCompany": geolocationCompany,
			"skills":             skills.String,
			"backendStack":       backendStack.String,
			"frontendStack":      frontendStack.String,
			"databaseStack":      databaseStack.String,
		}

		found = true
	}

	if !found {
		vacanciesFilter["notFound"] = map[string]interface{}{"message": "нет значений в БД"}
		return nil, status.Error(codes.NotFound, "нет значений в БД")
	}

	err = cache.SaveCache("vacancies_categories_"+hashMap, vacanciesFilter)
	if err != nil {
		return nil, err
	}

	vacanciesContentMap := createVacanciesContentMap(vacanciesFilter)

	return vacanciesContentMap, nil
}

func (s *Service) AddVacanciesService(ctx context.Context, vacanciesMap map[string]string) (int32, error) {
	/*t, err := time.Parse("2006-01-02 15:04:05", datetime)
	if err != nil {
		return 0, err
	}*/

	// Начало транзакции
	tx, err := db.Conn.Begin()
	if err != nil {
		return 0, err
	}

	var requirementsID int
	err = tx.QueryRow(`INSERT INTO requirements ("educationID", "experienceYears") 
		VALUES ($1, $2) 
		RETURNING id`, vacanciesMap["educationId"], vacanciesMap["experienceYears"]).Scan(&requirementsID)
	if err != nil {
		tx.Rollback()
		fmt.Println("2")
		return 0, err
	}

	var workingConditionsID int
	err = tx.QueryRow(`INSERT INTO "WorkingConditions" ("workMode", salary, "workHoursPerDay", "workSchedule", "salaryTaxIncluded") 
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING id`, vacanciesMap["workMode"], vacanciesMap["salary"], vacanciesMap["workHoursPerDay"], vacanciesMap["workSchedule"], vacanciesMap["salaryTaxIncluded"]).Scan(&workingConditionsID)
	if err != nil {
		tx.Rollback()
		fmt.Println("3")
		return 0, err
	}

	var vacancyId int32
	err = tx.QueryRow(`INSERT INTO vacancies (name, "departmentCompany", description, "categoryVacanciesID", "requirementsID", "workingConditionsID", "geolocationCompany") 
		VALUES ($1, $2, $3, $4, $5, $6, $7) 
		RETURNING id`, vacanciesMap["name"], vacanciesMap["departmentCompany"], vacanciesMap["description"], vacanciesMap["categoryVacancies"], requirementsID, workingConditionsID, vacanciesMap["geolocationCompany"]).Scan(&vacancyId)
	if err != nil {
		tx.Rollback()
		fmt.Println("4")
		return 0, err
	}

	if vacanciesMap["backendStack"] != "" {
		backendStack := strings.Split(vacanciesMap["backendStack"], ",")
		for _, item := range backendStack {
			insertQuery := `INSERT INTO "vacanciesStack" ("vacancyId", "stackId") VALUES ($1, $2)`
			_, err = tx.Exec(insertQuery, vacancyId, item)
			if err != nil {
				tx.Rollback()
				fmt.Println("7")
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
				tx.Rollback()
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
				tx.Rollback()
				return 0, err
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return vacancyId, nil
}

/*func (s *Service) DelVacanciesService(ctx context.Context, newsID int32) error {
	// Начало транзакции
	tx, err := db.Conn.Beginx()
	if err != nil {
		return err
	}

	// Удаление связей новости с категориями
	_, err = tx.Exec("DELETE FROM \"categoriesNews\" WHERE \"newsID\" = $1", newsID)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			return errRollback
		}
		return err
	}

	// Удаление самой новости
	_, err = tx.Exec("DELETE FROM news WHERE id = $1", newsID)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			return errRollback
		}
		return err
	}

	// Фиксация транзакции
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}*/

func createVacanciesContentMap(VacanciesMap map[string]map[string]interface{}) map[string]*pb.GetVacanciesItem {
	contentVacanciesMap := make(map[string]*pb.GetVacanciesItem)
	for _, data := range VacanciesMap {
		id, err := strconv.ParseInt(strings.TrimSpace(data["id"].(string)), 10, 32)
		if err != nil {
			log.Fatalf("Error converting string to int32: %v", err)
		}

		fmt.Println(data)

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
