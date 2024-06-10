package vacancies

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/lib/pq"
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

func (s *Service) GetVacanciesService(ctx context.Context) (map[string]*pb.VacanciesItem, error) {
	// Initialize newsSlice
	vacanciesMap := make(map[string]map[string]interface{})

	newsCheck, err := cache.IsExistInCache("vacancies")
	if newsCheck && err == nil {
		vacanciesMap, err = cache.ReadCache("vacancies")
		if err != nil {
			return nil, err
		}

		vacanciesContentMap := createVacanciesContentMap(vacanciesMap)

		return vacanciesContentMap, nil
	}

	// Данных нет
	rows, err := db.Conn.Query(`
		SELECT 
			v.id AS "vacancyId",
			v.name AS "vacancyName",
			r."experienceYears",
			wc."workMode",
			wc.salary,
			wc."salaryTaxIncluded",
			v."geolocationCompany"
		FROM 
			vacancies v
		JOIN 
			requirements r ON v."requirementsID" = r.id
		JOIN 
			"WorkingConditions" wc ON v."workingConditionsID" = wc.id
		GROUP BY 
			v.id, r."experienceYears", wc."workMode", wc.salary, wc."workHoursPerDay", wc."workSchedule", wc."salaryTaxIncluded";
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name, geolocationCompany string
		var experienceYears int
		var salary float64
		var workMode, salaryTaxIncluded bool

		err := rows.Scan(
			&id,
			&name,
			&experienceYears,
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
			"experienceYears":    experienceYears,
			"workMode":           workMode,
			"salary":             fmt.Sprint(salary),
			"salaryTaxIncluded":  salaryTaxIncluded,
			"geolocationCompany": geolocationCompany,
		}
	}

	err = cache.SaveCache("vacancies", vacanciesMap)
	if err != nil {
		return nil, err
	}

	vacanciesContentMap := createVacanciesContentMap(vacanciesMap)

	return vacanciesContentMap, nil
}

func (s *Service) GetVacanciesByIdService(ctx context.Context, id int32) (map[string]*pb.VacanciesIdItem, error) {
	// Initialize newsSlice
	vacanciesMap := make(map[string]map[string]interface{})

	vacanciesCheck, err := cache.IsExistInCache("vacancies_" + fmt.Sprint(id))
	if vacanciesCheck && err == nil {
		vacanciesMap, err = cache.ReadCache("vacancies_" + fmt.Sprint(id))
		if err != nil {
			return nil, err
		}

		vacanciesContentMap := createVacanciesIdContentMap(vacanciesMap)

		return vacanciesContentMap, nil
	}

	// Данных нет
	rows, err := db.Conn.Query(`
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
		WHERE 
			v.id = $1
		GROUP BY 
			v.id, cv.name, r."experienceYears", e.name, e."placeEducation", wc."workMode", wc.salary, wc."workHoursPerDay", wc."workSchedule", wc."salaryTaxIncluded";
	`, id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

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

		vacanciesMap[fmt.Sprint(vacancyId)] = map[string]interface{}{
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
	}

	err = cache.SaveCache("vacancies_"+fmt.Sprint(id), vacanciesMap)
	if err != nil {
		return nil, err
	}

	vacanciesContentMap := createVacanciesIdContentMap(vacanciesMap)

	return vacanciesContentMap, nil
}

func (s *Service) GetVacanciesByFilterService(ctx context.Context, vacanciesFilter map[string]string) (map[string]*pb.VacanciesItem, error) {
	// Initialize newsSlice
	vacanciesMap, hashMap := cache.ConvertMap(vacanciesFilter)

	vacanciesCheck, err := cache.IsExistInCache("vacancies_filter_" + hashMap)
	if vacanciesCheck && err == nil {
		vacanciesMap, err = cache.ReadCache("vacancies_filter_" + hashMap)
		if err != nil {
			return nil, err
		}

		vacanciesContentMap := createVacanciesContentMap(vacanciesMap)

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
		if value, ok := vacanciesFilter[condition.key]; ok && value != "" {
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

		vacanciesMap[fmt.Sprint(vacancyId)] = map[string]interface{}{
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
	}

	err = cache.SaveCache("vacancies_categories_"+hashMap, vacanciesMap)
	if err != nil {
		return nil, err
	}

	vacanciesContentMap := createVacanciesContentMap(vacanciesMap)

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

	var categoryVacanciesID int
	err = tx.QueryRow(`INSERT INTO "categoriesVacancies" (name) 
		VALUES ($1) 
		ON CONFLICT (name) DO NOTHING 
		RETURNING id`, vacanciesMap["categoryVacancies"]).Scan(&categoryVacanciesID)
	if err != nil && err != sql.ErrNoRows {
		tx.Rollback()
		return 0, err
	}

	var requirementsID int
	err = tx.QueryRow(`INSERT INTO requirements ("educationID", "experienceYears") 
		VALUES ($1, $2) 
		RETURNING id`, vacanciesMap["educationId"], vacanciesMap["experienceYears"]).Scan(&requirementsID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	var workingConditionsID int
	err = tx.QueryRow(`INSERT INTO "WorkingConditions" ("workMode", salary, "workHoursPerDay", "workSchedule", "salaryTaxIncluded") 
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING id`, vacanciesMap["workMode"], vacanciesMap["salary"], vacanciesMap["workHoursPerDay"], vacanciesMap["workSchedule"], vacanciesMap["salaryTaxIncluded"]).Scan(&workingConditionsID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	var vacancyId int32
	err = tx.QueryRow(`INSERT INTO vacancies (name, "departmentCompany", description, "categoryVacanciesID", "requirementsID", "workingConditionsID", "geolocationCompany") 
		VALUES ($1, $2, $3, $4, $5, $6, $7) 
		RETURNING id`, vacanciesMap["name"], vacanciesMap["departmentCompany"], vacanciesMap["description"], categoryVacanciesID, requirementsID, workingConditionsID, vacanciesMap["geolocationCompany"]).Scan(&vacancyId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	if len(vacanciesMap["skills"]) > 0 {
		query := `INSERT INTO "requirementsSkills" ("requirementsID", "skillsID") 
			SELECT $1, id FROM skills WHERE name = ANY($2)`
		_, err = tx.Exec(query, requirementsID, pq.Array(vacanciesMap["skills"]))
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	if vacanciesMap["backendStack"] != "" {
		backendStack := strings.Split(vacanciesMap["backendStack"], ",")
		for _, item := range backendStack {
			var stackID int
			query := `SELECT id FROM stack WHERE name = $1`
			err := tx.QueryRow(query, item).Scan(&stackID)
			if err != nil {
				tx.Rollback()
				return 0, err
			}

			insertQuery := `INSERT INTO "vacanciesStack" ("vacancyId", "stackId") VALUES ($1, $2)`
			_, err = tx.Exec(insertQuery, vacancyId, stackID)
			if err != nil {
				tx.Rollback()
				return 0, err
			}
		}
	}

	if vacanciesMap["frontendStack"] != "" {
		frontendStack := strings.Split(vacanciesMap["frontendStack"], ",")
		for _, item := range frontendStack {
			var stackID int
			query := `SELECT id FROM stack WHERE name = $1`
			err := tx.QueryRow(query, item).Scan(&stackID)
			if err != nil {
				tx.Rollback()
				return 0, err
			}

			insertQuery := `INSERT INTO "vacanciesStack" ("vacancyId", "stackId") VALUES ($1, $2)`
			_, err = tx.Exec(insertQuery, vacancyId, stackID)
			if err != nil {
				tx.Rollback()
				return 0, err
			}
		}
	}

	if vacanciesMap["databaseStack"] != "" {
		databaseStack := strings.Split(vacanciesMap["databaseStack"], ",")
		for _, item := range databaseStack {
			var stackID int
			query := `SELECT id FROM stack WHERE name = $1`
			err := tx.QueryRow(query, item).Scan(&stackID)
			if err != nil {
				tx.Rollback()
				return 0, err
			}

			insertQuery := `INSERT INTO "vacanciesStack" ("vacancyId", "stackId") VALUES ($1, $2)`
			_, err = tx.Exec(insertQuery, vacancyId, stackID)
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

func createVacanciesContentMap(VacanciesMap map[string]map[string]interface{}) map[string]*pb.VacanciesItem {
	contentVacanciesMap := make(map[string]*pb.VacanciesItem)
	for _, data := range VacanciesMap {
		id, err := strconv.ParseInt(strings.TrimSpace(data["id"].(string)), 10, 32)
		if err != nil {
			log.Fatalf("Error converting string to int32: %v", err)
		}

		content := &pb.VacanciesItem{
			Id:         int32(id),
			Title:      data["title"].(string),
			Text:       data["text"].(string),
			Datetime:   data["datetime"].(string),
			Categories: data["categories"].(string),
		}
		contentVacanciesMap[data["id"].(string)] = content
	}
	return contentVacanciesMap
}

func createVacanciesIdContentMap(VacanciesMap map[string]map[string]interface{}) map[string]*pb.VacanciesIdItem {
	contentVacanciesMap := make(map[string]*pb.VacanciesIdItem)
	for _, data := range VacanciesMap {
		id, err := strconv.ParseInt(strings.TrimSpace(data["id"].(string)), 10, 32)
		if err != nil {
			log.Fatalf("Error converting string to int32: %v", err)
		}

		content := &pb.VacanciesIdItem{
			Id:         int32(id),
			Title:      data["title"].(string),
			Text:       data["text"].(string),
			Datetime:   data["datetime"].(string),
			Categories: data["categories"].(string),
		}
		contentVacanciesMap[data["id"].(string)] = content
	}
	return contentVacanciesMap
}
