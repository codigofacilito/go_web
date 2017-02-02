package main

import(
	"net/http"
	"log"
	"fmt"
)
type customeHandler func(http.ResponseWriter, *http.Request)

type MuxFacilito struct{
	rutasFacilitas map[string] customeHandler //handlers
}

func (this *MuxFacilito) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fn := this.rutasFacilitas[r.URL.Path]
	fn(w,r)
}

func (this *MuxFacilito) AddMux(ruta string, fn customeHandler){
	this.rutasFacilitas[ruta] = fn
}

func main() {

	newMap := make(map[string]customeHandler)
	mux := &MuxFacilito{ rutasFacilitas: newMap }
	
	mux.AddMux("/hola", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hola desde una función anonima")
	})

	log.Fatal(http.ListenAndServe("localhost:3000",  mux ))
}