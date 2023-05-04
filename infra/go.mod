module github.com/GoDriveApp/GoDriveApi/infra

go 1.19

replace (
	github.com/GoDriveApp/GoDriveApi/core => ../core
	github.com/GoDriveApp/GoDriveApi/domain => ../domain
)

require (
	github.com/GoDriveApp/GoDriveApi/core v0.0.0-00010101000000-000000000000
	github.com/GoDriveApp/GoDriveApi/domain v0.0.0-00010101000000-000000000000
)