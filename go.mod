module github.com/mpieczaba/nimbus

go 1.15

replace github.com/99designs/gqlgen v0.13.0 => github.com/arsmn/gqlgen v0.13.2

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/form3tech-oss/jwt-go v3.2.2+incompatible
	github.com/go-playground/validator/v10 v10.4.1
	github.com/gofiber/fiber/v2 v2.5.0
	github.com/joho/godotenv v1.3.0
	github.com/rs/xid v1.2.1
	github.com/stretchr/testify v1.5.1 // indirect
	github.com/vektah/gqlparser/v2 v2.1.0
	gocv.io/x/gocv v0.26.0
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	gorm.io/driver/mysql v1.0.4
	gorm.io/gorm v1.20.12
)
