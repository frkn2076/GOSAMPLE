package environments

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	environment := os.Getenv("Environment")
	if environment != "PROD" && environment != "STAGE" {
		environment = "DEV"
	}

	fmt.Println("Environment selected as", environment)
	envFilePath := fmt.Sprintf("config/environments/%s.env", environment)
	godotenv.Load(envFilePath)
}
