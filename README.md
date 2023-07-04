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


# Dependencies

https://docs.tea.xyz/getting-started/install-tea

| Project    | Version |
| ---------- | ------- |
| go.dev     | ~1.19   |
| xcfile.dev | ~0.4.1  |
| npmjs.com  | ~9      |
| nodejs.org | ~20     |
