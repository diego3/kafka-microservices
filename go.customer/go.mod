module go.customer

go 1.16

require github.com/gorilla/mux v1.8.0

require (
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/google/uuid v1.3.0
	go.elastic.co/apm/module/apmgorilla v1.13.1
	golang.org/x/net v0.0.0-20201202161906-c7110b5ffcbb // indirect
	golang.org/x/sys v0.0.0-20210112080510-489259a85091 // indirect
	golang.org/x/tools v0.0.0-20201224043029-2b0845dc783e // indirect
)

replace go.customer/controller => ./
