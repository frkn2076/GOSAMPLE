package environments

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	environment := os.Args[1]
	if environment != "PROD" && environment != "STAGE" && environment != "DEV" {
		environment = "LOCAL-DEFAULT"
	}

	fmt.Println("Environment selected as", environment)
	envFilePath := fmt.Sprintf("config/environments/%s.env", environment)
	godotenv.Load(envFilePath)
}
