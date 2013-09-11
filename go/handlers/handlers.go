package handlers

import (
	"fmt"
	"net/http"
	"appserver"
	"appserver/config"
	"encoding/json"
)

type AngryCat struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Image_path string `json:"image_path"`
	Rank int `json:"rank"`
}

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
	http.HandleFunc( "/AngryCats/", makeJsonHandler(angryCatsHandler) )
}


func makeJsonHandler( fn func(w http.ResponseWriter, r *http.Request) ) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type","application/json; charset=utf-8")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		fn(w,r)
	}
}

func angryCatsHandler( w http.ResponseWriter, r *http.Request ) {
	if len(r.URL.Path) > len("/AngryCats/") {
		angryCatHandler( w , r )
		return
	}
	angryCats := make([]AngryCat,0)
	angryCats = append(angryCats,AngryCat{1, "Wet Cat", "assets/images/cat2.jpg", 1 })
	angryCats = append(angryCats,AngryCat{2, "Bitey Cat", "assets/images/cat1.jpg",2 })
	angryCats = append(angryCats,AngryCat{3, "Surprised Cat", "assets/images/cat3.jpg", 3 })

	b, err := json.Marshal(angryCats)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	w.Write( b )
}

func angryCatHandler( w http.ResponseWriter, r *http.Request ) {
	/*
	var angryCat []AngryCat
	err := json.Unmarshal(r., &angryCat)
	if err != nil {
		fmt.Println("error:", err)
	}
	*/
	fmt.Printf("path %v\n",  r.URL.Path)
	//fmt.Printf("form data %v\n",  r.Form.Encode())
	w.Write([]byte("[]") )
}


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

