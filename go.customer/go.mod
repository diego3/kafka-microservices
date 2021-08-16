module go.customer

go 1.16

require github.com/gorilla/mux v1.8.0

require (
	github.com/go-redis/redis/v8 v8.11.1
	github.com/google/uuid v1.3.0
	go.elastic.co/apm v1.13.1 // indirect
	go.elastic.co/apm/module/apmgorilla v1.13.1 // indirect
)
require github.com/google/uuid v1.3.0

replace go.customer/controller => ./
