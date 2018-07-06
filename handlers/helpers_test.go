package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/micro-message/store"
)

var tmpMessages []store.Message

// This function is used for setup before executing the test functions
func TestMain(m *testing.M) {
	//Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Run the other tests
	os.Exit(m.Run())
}

// Helper function to create a router during testing
func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("../templates/*")
	}
	return r
}

// Helper function to process a request and test its response
func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {

	// Create a response recorder
	w := httptest.NewRecorder()

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

func getLoginUser() string {
    params := url.Values{}
    params.Add("username", "user1")
    params.Add("password", "pass1")
    params.Add("email", "user1@qlik.com")

    return params.Encode()
}

func getRegistrationUser() string {
	param := url.Values{}
	param.Add("username", "test123")
	param.Add("password", "haha")
	param.Add("email", "test123@qlik.com")
	return param.Encode()
}

func getDeleteMessage() string {
    param := url.Values{}
    param.Add("messageid", "1")
    return param.Encode()
}
