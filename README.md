### To start local server development mode

1. Install required packages `CompileDaemon` to add auto compile ability
```
go get github.com/githubnemo/CompileDaemon
```
And then,
```
go install github.com/githubnemo/CompileDaemon
```
> This should be required only the first time

2. install packages from go.mod with this command
```
go mod download
```

3. Create `.env` from `.env.template` and change the following value to your preferences
- `{ANY_RANDOM_SECRET_KEY_TO_GENERATE_JWT}` = JWT secret
- `DB_USER` = username of your database. This value also used in Docker
- `DB_PASSWORD` = password of your database. This value also used in Docker
- `DB_NAME` = database name of your database. This value also used in Docker

Leave the rest as they are.

4. After you have setup variables in an .env by following the .env.template and your database is up and running, don't forget to run a migration
```
go run migrations/migrations.go
```

5. Now you should ready to start development, simply run this command to start local server.
```
CompileDaemon -command="./go-app"
```
