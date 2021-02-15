module github.com/mpieczaba/nimbus

go 1.15

replace github.com/99designs/gqlgen v0.13.0 => github.com/arsmn/gqlgen v0.13.2

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/go-playground/validator/v10 v10.4.1
	github.com/gofiber/fiber/v2 v2.5.0
	github.com/joho/godotenv v1.3.0
	github.com/rs/xid v1.2.1
	github.com/vektah/gqlparser/v2 v2.1.0
	gorm.io/driver/mysql v1.0.4
	gorm.io/driver/postgres v1.0.8
	gorm.io/gorm v1.20.12
)
