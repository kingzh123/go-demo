package route

import (
	"fmt"
	"net/http"
)

type Route struct {
	ip string
}

func (route *Route) sayHallo(w http.ResponseWriter, r http.Request)  {
	r.ParseForm()
	fmt.Println(r.Form)
}