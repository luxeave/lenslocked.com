package main

import (
	"fmt"
	"net/http"

	//"github.com/julienschmidt/httprouter"
	"github.com/gorilla/mux"
)

/*func someHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	//fmt.Fprint(w, r.URL.Path)
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to my awesome site</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch, please send an email to <h href=\"mailtoLsupport@lenslocked.com\">support@lenslocked.com</a>.")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>We could not find the page you were looking for :(</h1><p>Please email us if you keep being sent to an invalid page.</p>")
	}
}*/

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my awesome site</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch, please send an email to <h href=\"mailtoLsupport@lenslocked.com\">support@lenslocked.com</a>.")
}

// Hello is
/*func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}*/

func main() {
	// MUX
	//mux := &http.ServeMux{}
	//mux.HandleFunc("/", someHandlerFunc) // SLASH IS IMPORTANT
	//http.ListenAndServe(":3000", mux)

	// HTTP
	//http.HandleFunc("/", someHandlerFunc)
	//http.ListenAndServe(":3000", nil) // :3000 or localhost:3000 the same

	// JULIEN
	//router := httprouter.New()
	//router.GET("/hello/:name", Hello)
	//http.ListenAndServe(":3000", router) // :3000 or localhost:3000 the same

	// GORILLA
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":3000", r)
}
