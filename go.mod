module IdentityProvider

go 1.12

require (
    api v0.0.0
	common v0.0.0
	service v0.0.0
	model v0.0.0
)

replace api v0.0.0 => ./api

replace common v0.0.0 => ./common

replace service v0.0.0 => ./service

replace model v0.0.0 => ./model
