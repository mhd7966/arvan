# arvan -> Code , Wallet
Do the following for each program

 Download swag by using:
```sh
$ go get -u github.com/swaggo/swag/cmd/swag
```

 Swagger init changes, run in main project directory : 
```sh
$ swag init
```

  To run program : 
```sh
$ docker-compose up
```

  Download sql-migrate by using:
```
$ go get -v github.com/rubenv/sql-migrate/...
```
For running migration, run this command (Note: The database must be up)
```
$ sql-migrate up
```
