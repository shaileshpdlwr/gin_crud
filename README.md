# Local path : /home/shailesh/learn_go

# gin_crud
To understand crud operation using golang with gin 
# JSON CRUD USING gin and postgresql



# Installation 
>> go get -u package-name
The -u flag tells Go to update both the specified package and any of its dependencies to their latest versions.


>> go get package-name
This will download and install the specified package without updating it to the latest version.


# go.mod vs go.sum ?
In Go, go.mod and go.sum files are essential for managing dependencies in a Go project.

go.mod: The go.mod file is the module definition file for a Go project. It describes the module's name, its dependencies, and the minimum required versions of those dependencies. This file allows Go to maintain a clear and precise record of the project's dependencies, ensuring that builds are reproducible.

go.sum: The go.sum file is also created and maintained by Go. It contains *cryptographic* checksums of the specific versions of each module and its dependencies. These checksums are used to ensure the integrity of the downloaded modules, protecting against tampering during transit.


#Connection to DB: postgresql
```import (
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
```
# db migration command
go run migrate/migrate.go
