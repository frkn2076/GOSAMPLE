package environments

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Load() {
	environment := "LOCAL-DEFAULT"
	if len(os.Args) > 1 {
		if os.Args[1] == "DEV" || os.Args[1] == "STAGE" || os.Args[1] == "PROD" {
			environment = os.Args[1]
		}
	} 

	fmt.Println("Environment selected as", environment)
	envFilePath := fmt.Sprintf("config/environments/%s.env", environment)
	godotenv.Load(envFilePath)
}
