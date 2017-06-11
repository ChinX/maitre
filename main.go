package main

import (
	"log"

	"github.com/chinx/maitre/router"
	"github.com/chinx/maitre/svrmux"
)

func main() {
	r := svrmux.NewRouter()
	router.SoftwareRouter(r)
	sources := r.Sources()

	for i := 0; i < len(sources); i++ {
		log.Println(*sources[i])
	}
}
