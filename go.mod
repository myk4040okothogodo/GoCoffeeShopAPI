module github.com/myk4040okothogodo/GoMicroserve

go 1.17

replace github.com/myk4040okothogodo/GoMicroserve/handlers => ./handlers

require (
	github.com/gorilla/mux v1.8.0
	github.com/myk4040okothogodo/GoMicroserve/handlers v0.0.0-00010101000000-000000000000
)

require github.com/myk4040okothogodo/GoMicroserve/data v0.0.0-00010101000000-000000000000 // indirect

replace github.com/myk4040okothogodo/GoMicroserve/data => ./data
