### To start local server development mode

First, install required packages `CompileDaemon` to add auto compile ability
```
go get github.com/githubnemo/CompileDaemon
```
And then,
```
go install github.com/githubnemo/CompileDaemon
```
> This should be required only the first time

Then, install packages from go.mod with this command
```
go mod download
```

After you have setup variables in an .env by following the .env.template and your database is up and running, don't forget to run a migration
```
go run migrations/migrations.go
```

Now you should ready to start development, simply run this command to start local server.
```
CompileDaemon -command="./go-app"
```