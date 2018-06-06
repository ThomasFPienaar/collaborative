package httplayer

import (
	"net/http"
	"log"
)
var resolver = RequestResolver{}

func SetupHttpServer() RequestResolver {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

	return resolver

}

func AddEndpoint(s string, handler EndpointHandler) {
	resolver.AddEndpoint(s, handler)

}

func handler(w http.ResponseWriter, r *http.Request) {
	resolver.DoRequest(w, r)
}


