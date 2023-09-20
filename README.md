### databse migration

[golang-migrate] (https://github.com/golang-migrate/migrate/blob/v4.16.2/database/postgres/TUTORIAL.md)

## create PostgreSQL database called example, user postgres, password password and host is localhost

`psql -h localhost -U postgres -w -c "create database example;"`

## Postgres url

`export POSTGRESQL_URL='postgres://postgres:password@localhost:5432/example?sslmode=disable'`

## create a migration up and down files named <create_users_table>

`migrate create -ext sql -dir db/migrations -seq create_users_table`

## run migration

`migrate -database ${POSTGRESQL_URL} -path db/migrations up`
