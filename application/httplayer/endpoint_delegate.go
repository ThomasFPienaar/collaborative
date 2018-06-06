package httplayer

import (
	"net/http"
	"fmt"
	"encoding/json"
)


type EndpointHandler interface {
	Handler(r *http.Request) interface{}
}
type RequestResolver struct {
	handlerList [] endpointKeyValue
}

type endpointKeyValue struct {
	name string
	handler EndpointHandler
}
func (resolver RequestResolver) DoRequest(w http.ResponseWriter, request *http.Request) {
	if w == nil || request == nil {
		return
	}

	for _, v := range resolver.handlerList {
		if v.name == request.URL.Path[1:] {
			jsonOutput, _  := json.Marshal(v.handler.Handler(request))
			fmt.Fprintf(w,string(jsonOutput) )
			break
		}
	}
}
func (resolver *RequestResolver) AddEndpoint(s string, handler EndpointHandler) {
	endpoint := endpointKeyValue{s, handler}
	resolver.handlerList = append(resolver.handlerList, endpoint)
}
