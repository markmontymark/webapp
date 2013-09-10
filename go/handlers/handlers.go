package handlers

import (
	"fmt"
	"net/http"
	"appserver"
	"appserver/config"
)

type MyHandlers struct {} // implements appserver/handlers/handlers

type AppContext struct {
	appserver.Context
}

type SessionContext struct { // embeds appserver/context.go
	appserver.Context
}

type RequestContext struct {
	appserver.Context
}

func NewMyHandlers() *MyHandlers {
	return &MyHandlers{}
}

func (this *MyHandlers) Load ( cfg *config.Config ) {
	http.HandleFunc( "/create/", createHandler )	
	http.HandleFunc( "/read/", readHandler )	
	http.HandleFunc( "/update/", updateHandler )	
	http.HandleFunc( "/delete/", deleteHandler )	
	http.HandleFunc( "/AngryCats", makeJsonHandler(angryCatsHandler) )
	http.HandleFunc( "/AngryCat", makeJsonHandler(angryCatHandler) )
}


func makeJsonHandler( fn func(w http.ResponseWriter, r *http.Request) ) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type","application/json; charset=utf-8")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		fn(w,r)
	}
}

func angryCatsHandler( w http.ResponseWriter, r *http.Request ) {
	w.Write([]byte(`[
{"id":1, "name": "Wet Cat", "image_path": "assets/images/cat2.jpg", "rank":1 }, 
{"id":2, "name": "Bitey Cat","image_path": "assets/images/cat1.jpg","rank":2 },
{"id":3, "name": "Surprised Cat", "image_path": "assets/images/cat3.jpg", "rank":3 }]`))
}

func angryCatHandler( w http.ResponseWriter, r *http.Request ) {
	w.Header().Add("content-type","application/json; charset=utf-8")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Write([]byte(`[
{"id":1, "name": "Wet Cat", "image_path": "assets/images/cat2.jpg", "rank":1 }, 
{"id":2, "name": "Bitey Cat","image_path": "assets/images/cat1.jpg","rank":2 },
{"id":3, "name": "Surprised Cat", "image_path": "assets/images/cat3.jpg", "rank":3 }]`))
}


//new AngryCat({ name: 'Wet Cat', image_path: 'assets/images/cat2.jpg' }),
      //new AngryCat({ name: 'Bitey Cat', image_path: 'assets/images/cat1.jpg' }),
      //new AngryCat({ name: 'Surprised Cat', image_path: 'assets/images/cat3.jpg' })
	


func badInputHandler(w http.ResponseWriter, r *http.Request, msgs []string) {
	for msg := range msgs {
		fmt.Printf("msg %s\n",msg)
		w.Write([]byte("{msg:\"msg\"}") )
	}
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	w.Write( []byte("{\"derp\":howdy}"))
}

func readHandler(w http.ResponseWriter, r *http.Request) {
	w.Write( []byte("{\"derp\":howdy}"))
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	w.Write( []byte("{\"derp\":howdy}"))
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Write( []byte("{\"derp\":howdy}"))
}

