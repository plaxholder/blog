package main

import (
	"fmt"
	"io"
	"os"
	"net/http"
	"html/template"
	"log"

	"github.com/plaxholder/blog"
)

func main(){
	err := run(os.Args, os.Stdout)
	if err != nil{
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(args []string, stdout io.Writer) error{
	mux := http.NewServeMux()

	postTemplate := template.Must(template.ParseFiles("post.gohtml"))
	indexTemplate := template.Must(template.ParseFiles("index.gohtml"))

	fr := blog.FileReader{
		Dir : "posts",
	}
	mux.HandleFunc("GET /posts/{slug}", blog.PostHandler(fr,postTemplate))
	mux.HandleFunc("GET /", blog.IndexHandler(fr, indexTemplate))

	err := http.ListenAndServe(":3030", mux)
	if err != nil{
		log.Fatal(err)
	}
	return nil
}
