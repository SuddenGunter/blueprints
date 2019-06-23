package main

import (
	"log"
	"net/http"
)

func main() {
	r := newRoom()
	http.Handle("/room", r)
	http.Handle("/avatars/",
		http.StripPrefix("/avatars/",
			http.FileServer(http.Dir("./avatars"))))
	// get the room going
	go r.run()
	// start the web server
	log.Println("Starting web server on", ":8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
