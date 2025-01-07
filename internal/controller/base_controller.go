package controller

type baseController struct {
	Menu menu
}

type menu struct {
	Name  string
	Route string
}
