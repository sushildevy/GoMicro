package bootstrap

import (
	"net/http"
	"webapp/controller"
)

func Bootapplication() {

	http.HandleFunc("/employees", controller.GetEmployee)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
