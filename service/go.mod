module service

go 1.12

require (
	common v0.0.0
	github.com/coreos/go-systemd v0.0.0-20190719114852-fd7a80b32e1f // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gofrs/uuid v3.2.0+incompatible
	github.com/jackc/pgx/v4 v4.6.0
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/kr/pty v1.1.8 // indirect
	github.com/stretchr/objx v0.2.0 // indirect
	golang.org/x/net v0.0.0-20190813141303-74dc4d7220e7 // indirect
	golang.org/x/tools v0.0.0-20190823170909-c4a336ef6a2f // indirect
	model v0.0.0
)

replace common v0.0.0 => ../common

replace model v0.0.0 => ../model
