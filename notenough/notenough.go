package notenough

type MyApp struct {
	name string
}

func(a *MyApp) setName(name string){
	a.name = name
}

func (a *MyApp) SimpleMethod() string {
	// comment here
	a.setName("Golang At Speed")
	return a.name + " the API of notenough only comprises a really simple exported method"
}
