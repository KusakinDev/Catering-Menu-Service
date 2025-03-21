/*
 * Catering service
 *
 * Menu service.
 *
 * API version: 1.0.0
 */

package main

import (
	"log"

	sw "github.com/KusakinDev/Catering-Menu-Service/internal/routes"
)

func main() {
	routes := sw.ApiHandleFunctions{}

	log.Printf("Server started")

	router := sw.NewRouter(routes)

	log.Fatal(router.Run(":8080"))
}
