# golang-microservice

# create db with docker
docker run --name postgres -e POSTGRES_USER=xxxx -e POSTGRES_PASSWORD=xxxx -e POSTGRES_DB=golang-microservice -p 5432:5432 -d postgres
