module example.com/hello

go 1.22.0

replace example.com/greetings => ../greetings
replace global/utils => ../utils

require (
	example.com/greetings v0.0.0-00010101000000-000000000000
	global/utils v0.0.0-00010101000000-000000000000
)
