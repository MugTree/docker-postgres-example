### Create the database and add some data

createdb docker_test

### Print the status of all migrations

GOOSE_DRIVER=postgres GOOSE_DBSTRING="user=portable sslmode=disable" goose status

### Run the statements

GOOSE_DRIVER=postgres GOOSE_DBSTRING="user=portable sslmode=disable" goose up
