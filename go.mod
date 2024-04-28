module go-weather-app

replace retriever => ./retriever

replace location => ./location

replace weather => ./weather

go 1.21.6

require retriever v0.0.0-00010101000000-000000000000

require (
	location v0.0.0-00010101000000-000000000000 // indirect
	weather v0.0.0-00010101000000-000000000000 // indirect
)
