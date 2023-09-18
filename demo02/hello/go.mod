module example.com/hello

go 1.21.0
// 重定向到本地
replace example.com/greetings => ../greetings

require example.com/greetings v0.0.0-00010101000000-000000000000
