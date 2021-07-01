* To run docker compose
docker-compose --env-file ./config/environments/DEV.env up

* To start postgre docker instance explicitly 
docker run -p 5432:5432 --name sampledb -e POSTGRES_USER=uSeR1 -e POSTGRES_PASSWORD=12345 -e POSTGRES_DB=SampleDB -d postgres:13-alpine

* To start mongo docker instance explicitly 
docker run -p 27017:27017 --name logsdb -e MONGO_INITDB_ROOT_USERNAME=uSeRrr -e MONGO_INITDB_ROOT_PASSWORD=PassWorD -d mongo
