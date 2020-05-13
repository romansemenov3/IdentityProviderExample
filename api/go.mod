module api

go 1.12

require github.com/gorilla/mux v1.7.4

require service v0.0.0

replace service v0.0.0 => ../service

require model v0.0.0

replace model v0.0.0 => ../model
