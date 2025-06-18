$ migrate create -seq -ext=.sql -dir=./migrations {name_table}

$ migrate -path=./migrations -database="postgres://greenlight:123456@localhost/greenlight" up
