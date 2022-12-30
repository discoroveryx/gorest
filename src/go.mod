module gorest

go 1.19

replace app/user/actions => ./app/user/actions

replace app/user/models => ./app/user/models

replace app/user/exceptions => ./app/user/exceptions

replace app/user/handlers => ./app/user/handlers

replace app/user/repositories => ./app/user/repositories

replace app/user/transformers => ./app/user/transformers

replace app/user/controllers => ./app/user/controllers

replace app/auth/controllers => ./app/auth/controllers

replace app/auth/transformers => ./app/auth/transformers

replace app/auth/actions => ./app/auth/actions

replace app/auth/handlers => ./app/auth/handlers

replace app/auth/exceptions => ./app/auth/exceptions

replace app/auth/configs => ./app/auth/configs

replace app/user/myte => ./app/user/myte

replace dbstorage => ./dbstorage

replace config => ./config

replace helpers => ./helpers

replace transport => ./transport

replace tests => ./tests

replace tests/fixtures => ./tests/fixtures

require (
	dbstorage v0.0.0-00010101000000-000000000000
	transport v0.0.0-00010101000000-000000000000
)

require (
	app/auth/actions v0.0.0-00010101000000-000000000000 // indirect
	app/auth/configs v0.0.0-00010101000000-000000000000 // indirect
	app/auth/controllers v0.0.0-00010101000000-000000000000 // indirect
	app/auth/exceptions v0.0.0-00010101000000-000000000000 // indirect
	app/auth/handlers v0.0.0-00010101000000-000000000000 // indirect
	app/auth/transformers v0.0.0-00010101000000-000000000000 // indirect
	app/user/actions v0.0.0-00010101000000-000000000000 // indirect
	app/user/controllers v0.0.0-00010101000000-000000000000 // indirect
	app/user/exceptions v0.0.0-00010101000000-000000000000 // indirect
	app/user/handlers v0.0.0-00010101000000-000000000000 // indirect
	app/user/models v0.0.0-00010101000000-000000000000 // indirect
	app/user/repositories v0.0.0-00010101000000-000000000000 // indirect
	app/user/transformers v0.0.0-00010101000000-000000000000 // indirect
	config v0.0.0-00010101000000-000000000000 // indirect
	github.com/gin-contrib/cors v1.4.0 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.8.1 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.10.0 // indirect
	github.com/goccy/go-json v0.9.7 // indirect
	github.com/golang-jwt/jwt/v4 v4.4.3 // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.13.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.1 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.12.0 // indirect
	github.com/jackc/pgx/v4 v4.17.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mattn/go-sqlite3 v1.14.15 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.1 // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa // indirect
	golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2 // indirect
	golang.org/x/sys v0.0.0-20210806184541-e5e7981a1069 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/driver/postgres v1.4.5 // indirect
	gorm.io/driver/sqlite v1.4.3 // indirect
	gorm.io/gorm v1.24.2 // indirect
	helpers v0.0.0-00010101000000-000000000000 // indirect
	tests/fixtures v0.0.0-00010101000000-000000000000 // indirect
)
