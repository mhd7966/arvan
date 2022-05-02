# arvan -> Code , Wallet
Do the following for each program

 Swagger init changes, run in main project directory : 
```
  swag init
```

  To run program : 
```
    docker-compose up
```

  For running migration, run this command (Note: The database must be up):
```
    go get -v github.com/rubenv/sql-migrate/...
```
  then:
```
    sql-migrate up
```

