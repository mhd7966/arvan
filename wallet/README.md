# VulEQ

For running migration, run this command (Note: The database must be up):
```
go get -v github.com/rubenv/sql-migrate/...
```
then:
```
sql-migrate up
```
 Swagger init changes, run in main project directory : 
```
  swag init -g cmd/api/main.go
```

  To run program : 
```
    docker-compose up
```