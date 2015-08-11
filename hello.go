package hello

import (
	"fmt"
	"net/http"

	"appengine"
	"appengine/user"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(rw http.ResponseWriter, req *http.Request) {
	// Crea un nuevo contexto
	c := appengine.NewContext(req)
	// Obtiene el usuario actual
	u := user.Current(c)
	// Si no hay usuario (o sea no se ha logueado)
	if u == nil {
		// Obtenemos la url de logueo
		url, err := user.LoginURL(c, req.URL.String())
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		// Mostramos la url en pantalla
		rw.Header().Set("Location", url)
		rw.WriteHeader(http.StatusFound)
		return
	}
	//
	fmt.Fprintf(rw, "Hello, %v!", u)
}
