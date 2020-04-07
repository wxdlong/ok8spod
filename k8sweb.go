package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		log.Printf("Access ok8spod...%s%s",r.Host,r.URL)
		envs := os.Environ()
		env := "Request from " + r.Host + r.URL.String() + "\n"
		for _, en := range envs {
			env = env + en + "\n"
		}
		w.Write([]byte(env))
	})
	log.Println("Starting golang ok8spod server ...")
	log.Fatal(http.ListenAndServe(":80", nil))
}
