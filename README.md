# You owe me

API for keeping track of who owes me what.

## Environment variables
```Environment
GO_USER
GO_PASS
GO_HOST
GO_PORT
GO_DBNAME
GO_SSLMODE
GO_TZ
JWT_SECRET
```

## DB
gorm adapter using postgresql database.

## Authentication and Authorization
Classic JWT stuff w/o refresh token and casbin for authorization.