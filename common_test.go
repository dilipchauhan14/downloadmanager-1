// Author : Dilip Chauhan
//date : 17/01/2020
package main
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	//Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Run the other tests
	os.Exit(m.Run())
}

// Helper function to create a router during testing
func getRouter() *gin.Engine {
	r := gin.Default()

	//r.LoadHTMLGlob("templates/*")

	return r
}

func processHttpRequest(r *gin.Engine, req *http.Request) *httptest.ResponseRecorder {
	// Create a response recorder
	w := httptest.NewRecorder()

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	return w
}
