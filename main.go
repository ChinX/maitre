package main

import (
	"github.com/chinx/maitre/module/svrmux"
	"github.com/chinx/maitre/router"
	"log"
)

func main() {
	r := svrmux.NewRouter()
	router.SoftwareRouter(r)
	sources := r.Sources()

	for i := 0; i < len(sources); i++ {
		log.Println(*sources[i])
	}
}
