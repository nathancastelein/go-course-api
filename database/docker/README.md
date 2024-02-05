# Database

## Run database with docker

```
docker run -d --name gocourse-postgres \
    -v ./init.sql:/docker-entrypoint-initdb.d/init.sql \
    -e POSTGRES_PASSWORD=test \
    -e POSTGRES_USER=gocourse \
    -e POSTGRES_DB=gocourse \
    -p 5432:5432 \
    postgres:latest
```

## Exec PSQL

```
docker exec -it gocourse-postgres psql -U gocourse
```