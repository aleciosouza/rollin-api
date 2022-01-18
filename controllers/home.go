package controllers

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	//Page containing API documentation
	fmt.Fprint(w, "Home Page")
}
