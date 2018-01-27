package main

import (
	"log"
	"net/http"
	"io/ioutil"
	"flag"
	"fmt"
)

func main() {
	dirFlag := flag.String("dir", ".", "The directory you want to serve")
	portFlag := flag.Int("port", 8080, "The port on which you want to serve the files")

	flag.Parse()

	fs := http.FileServer(http.Dir(*dirFlag))
	log.Println("Starting server ...")
	http.Handle("/", loggingMiddleware(http.StripPrefix("/", fs)))

	num, names, err := info(*dirFlag)
	if err != nil {
		log.Fatal("Failed to read the directory")
	}
	log.Printf("Serving %v files/directories. First few include:", num)
	for _, name := range names {
		log.Printf("... %v", name)
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", *portFlag), nil))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Printf("Serving %v", r.URL)
			next.ServeHTTP(w, r)
		})
}

func info(dir string) (int, []string, error) {
	files, err := ioutil.ReadDir(dir)
	var names []string
	if err != nil {
		return 0, names, err
	}
	if len(files) > 3 {
		for _, file := range files[:3] {
			names = append(names, file.Name())
		}
	}
	return len(files), names, err
}
