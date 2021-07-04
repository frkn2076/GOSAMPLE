# GOSAMPLE

<br>

## To run docker compose
* **Modify .env file first. It has ENV=DEV as default.**
   - Set ENV variable as DEV to run Development environment (like ENV=DEV)
   - Set ENV variable as STAGE to run Staging environment (like ENV=STAGE)
   - Set ENV variable as PROD to run Production environment (like ENV=PROD)
* **docker-compose --env-file .env up**

<br>

## To run on local machine
* **You may want to have database instances on docker while running project locally.**
   - To start postgre docker instance
      - docker run -p 5432:5432 --name sampledb -e POSTGRES_USER=uSeR1 -e POSTGRES_PASSWORD=12345 -e POSTGRES_DB=SampleDB -d postgres:13-alpine
   - To start mongo docker instance
      - docker run -p 27017:27017 --name logsdb -e MONGO_INITDB_ROOT_USERNAME=uSeRrr -e MONGO_INITDB_ROOT_PASSWORD=PassWorD -d mongo
* **go run main.go {ENV}**
   - **go run main.go** (runs with default config, see ./config/environments/LOCAL-DEFAULT.env)
   - **go run main.go DEV** (runs Development environment locally)

<br>

## To run unit tests
* **go test**

<br>

## Technology Details
* **freecache** for caching, please visit for more detail https://github.com/coocood/freecache
* **godotenv** for env. variables initalization, please visit for more detail https://github.com/joho/godotenv
* **gorilla** for session, please visit for more detail github.com/gorilla/sessions
* **gin-gonic** as web framework, please visit for more detail github.com/gin-gonic/gin
* **gorm** as ORM tool, please visit for more detail https://gorm.io/
