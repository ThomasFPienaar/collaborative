package httplayer

import (
	"testing"
	"net/http/httptest"
	"net/http"
)

type testEndpointHandler struct {
	handler func(r *http.Request ) string
}

func (t testEndpointHandler) Handler(r *http.Request) string {
	return t.handler(r)
}

type MockWriter struct {

}

func (mw MockWriter) Header() http.Header {
 return nil
}
func (mw MockWriter) Write([]byte) (int, error) {
	return 0, nil
}
func (mw MockWriter)  WriteHeader(status int) {

}

func TestAddEndpoint(t *testing.T) {

	test1EndCalled := false
	test2EndCalled := false

	requestTest1End := httptest.NewRequest("GET", "/test1end", nil)
	requestTest2End := httptest.NewRequest("GET", "/test2end", nil)

	resolver := RequestResolver{}

	mockHandler1 := testEndpointHandler{func(r *http.Request) string {
		test1EndCalled = true
		return ""
	}}

	mockHandler2 := testEndpointHandler{func(r *http.Request) string {
		test2EndCalled = true
		return ""
	}}

	resolver.AddEndpoint("test1end", mockHandler1)
	resolver.AddEndpoint("test2end", mockHandler2)

	mw:=MockWriter{}

	resolver.DoRequest(mw, requestTest1End)
	resolver.DoRequest(mw, requestTest2End)

	if !test1EndCalled {
		t.Errorf("Handler not called")
	}

	if !test2EndCalled {
		t.Errorf("Handler not called")
	}


}
