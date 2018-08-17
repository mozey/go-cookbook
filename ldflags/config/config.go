package config

var Foo string
var fiz string

func GetFiz() string {
	if fiz == "" {
		return "bla"
	}
	return fiz
}