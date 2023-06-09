module github.com/GoDriveApp/GoDriveApi/infra

go 1.19

replace (
	github.com/GoDriveApp/GoDriveApi/core => ../core
	github.com/GoDriveApp/GoDriveApi/domain => ../domain
)

require (
	github.com/GoDriveApp/GoDriveApi/core v0.0.0
	github.com/GoDriveApp/GoDriveApi/domain v0.0.0-00010101000000-000000000000
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.3.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/crypto v0.6.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	gorm.io/driver/postgres v1.5.0 // indirect
	gorm.io/gorm v1.25.0 // indirect
)
