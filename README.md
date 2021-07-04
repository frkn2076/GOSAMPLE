# GOSAMPLE

<br>

## To run docker compose
* **Modify root.env file first. It has ENV=DEV as default.**
   - Set ENV variable as DEV to run Development environment (like ENV=DEV)
   - Set ENV variable as STAGE to run Staging environment (like ENV=STAGE)
   - Set ENV variable as PROD to run Production environment (like ENV=PROD)
* **docker-compose --env-file root.env up**

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
* **Freecache** for caching, please visit for more detail https://github.com/coocood/freecache
* **Gorilla** for session, please visit for more detail https://github.com/gorilla/sessions
* **Gin-gonic** as web framework, please visit for more detail https://github.com/gin-gonic/gin
* **Gorm** as ORM tool, please visit for more detail https://gorm.io
* **Mongo** as NoSQL database, please visit for more detail https://www.mongodb.com
* **PostgreSQL** as SQL database, please visit for more detail https://www.postgresql.org
* **JWT-go** as authorization, please visit for more detail https://github.com/dgrijalva/jwt-go
* **Testify** for unit tests, please visit for more detail https://github.com/stretchr/testify
* **Docker** for containerization, please visit for more detail https://www.docker.com

<br>

**PS:** 
* Exported Postman document for all services, is attached to root folder of project.
* Also deployed services are open for a while like Register Post service: http://37.148.212.195:5000/account/register
<br>
