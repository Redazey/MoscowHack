package resume_test

import (
	"fmt"
	"log"
	pb "moscowhack/gen/go/resume"
	"moscowhack/tests/suite"
	"os"
	"testing"
)

func ParseResumeTest(t *testing.T) {
	ctx, st := suite.New(t)

	file, err := os.ReadFile(".../.tests/resume/resumes/BA Senior.doc")
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
	}

	req := &pb.ResumeRequest{
		ResumeDoc: file,
	}

	t.Run("UserLogin Test", func(t *testing.T) {
		response, err := st.ResumeClient.ParseResume(ctx, req)
		if err != nil {
			log.Fatalf("Error when calling Login: %v", err)
		}
		fmt.Printf("Вывод: %v", response.ResumeMap)

		/*assert.Equal(t, exceptedKey, response.Key)*/
	})
}
