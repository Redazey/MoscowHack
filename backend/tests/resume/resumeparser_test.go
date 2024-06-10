package resume_test

import (
	"fmt"
	"log"
	pb "moscowhack/gen/go/resume"
	"moscowhack/tests/suite"
	"os"
	"testing"
)

func TestParseResume(t *testing.T) {
	ctx, st := suite.New(t)

	file, err := os.ReadFile("./resumes/BA Senior.doc")
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
	}

	req := &pb.ResumeRequest{
		ResumeDoc: file,
	}

	t.Run("ResumeParser Test", func(t *testing.T) {
		response, err := st.ResumeClient.ParseResume(ctx, req)
		if err != nil {
			log.Fatalf("Error when calling ResumeParser: %v", err)
		}
		fmt.Printf("Вывод: %v", response.ResumeMap)

		/*assert.Equal(t, exceptedKey, response.Key)*/
	})
}
