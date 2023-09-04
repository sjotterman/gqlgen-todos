# gqlgen-todos

following the gqlgen tutorial

## Tasks

### sql

Generate sqlc files

directory: ./backend/sqlc

```
sqlc generate
```

### generate-graphql

Generate graphql with graphqlgen

directory: ./backend
requires: sql

```
go run github.com/99designs/gqlgen generate
```

### generate

Generate types from graphql

directory: ./frontend
requires: generate-graphql

```
npm run codegen
```

### serve-dev

directory: ./backend
Run graphql server in development mode, automatically restarting on changes
Requires air - https://github.com/cosmtrek/air/

```
air -c .air.toml
```

### serve

directory: ./backend
Run graphql server

```
go run server.go
```

### start-db

Run database container

```
docker-compose up -d
```

### stop-db

Stop database container

```
docker-compose down
```

### connect-db-shell

Connect to shell in database container

```
docker exec -it local_pgdb /bin/bash
```

### db-dump
Dump the current state of the database so it can be used as a restore state.
```
docker exec local_pgdb pg_dump --clean --if-exists -cC -U user --dbname="postgres" > dump.sql
```

### db-check
Check that db is set up and accepting connections
```
docker exec -i local_pgdb pg_isready -d postgres -U user
```

### db-reset
Reset the local database to a fresh state

TODO: remove the sleep and replace it with something like the above
```
docker-compose down --volumes
xc start-db
sleep 5
cat dump.sql | docker exec -i local_pgdb psql -U user -d postgres
```


# Dependencies

https://docs.tea.xyz/getting-started/install-tea

| Project    | Version |
| ---------- | ------- |
| go.dev     | ~1.19   |
| xcfile.dev | ~0.4.1  |
| npmjs.com  | ~9      |
| nodejs.org | ~20     |
