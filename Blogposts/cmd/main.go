package main

import (
	"log"
	"os"

	blogposts "github.com/anicse37/TDD_Go"
)

func main() {
	posts, err := blogposts.NewPostsFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(posts)
}
