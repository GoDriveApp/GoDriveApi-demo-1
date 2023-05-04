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
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	gorm.io/gorm v1.25.0 // indirect
)
