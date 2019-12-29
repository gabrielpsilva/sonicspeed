package sonicspeed

import (
	"github.com/gabrielpsilva/sonicspeed/restserver"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestRestServer_StandardServer SonicSpeed framework exposes the rest server components under in
// restserver.RestServer to provide flexibility
func TestRestServer_StandardServer(t *testing.T) {

	// CREATE ENDPOINT
	// Setup server
	s := restserver.StandardServer()

	// Function to request
	fn := func(c *gin.Context){
		c.JSON(http.StatusOK, "world")
	}
	// Set up request and function
	s.Gin.Group("/v1").GET("/hello", fn)

	// CONSUME ENDPOINT
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/hello", nil)
	req.Header.Add("User-Agent", "Test")
	s.Gin.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("response code was %d; expected %d", w.Code, 200)
	}

	body := string(w.Body.Bytes())
	if body != "\"world\"\n"{
		t.Errorf("response body was %s; expected %s", body, "world")
	}

}

// TestRestServer_Log In can use the standard logging formatters by using GlobalLogrusFormat functions or do it
// yourself by calling logrus.SetFormatter
func TestRestServer_Log(t *testing.T) {
	tt := [] struct{
		name 		string
	}{
		{name:"JSON"},
		{name:"TEXT"},
	}

	// CREATE ENDPOINT
	s := restserver.StandardServer()
	s.AddRequestIDHeader("X-Request-Id", "x-cloud-trace-context")

	// Function to request
	fn := func(c *gin.Context){
		ctx := restserver.ContextWrapper(c)
		ctx.Log().Debugf("Hello from Hello World test endpoint")
		c.JSON(http.StatusOK, "world")
	}
	// Set up request and function
	s.Gin.Group("/v1").GET("/hello", fn)


	for _, tc := range tt {

		if tc.name == "JSON" {
			restserver.GlobalLogrusFormatJSON(logrus.DebugLevel)
		}
		if tc.name == "TEXT" {
			restserver.GlobalLogrusFormatTEXT(logrus.DebugLevel)
		}

		// CONSUME ENDPOINT
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/v1/hello", nil)
		req.Header.Add("User-Agent", "Test")
		s.Gin.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("response code was %d; expected %d", w.Code, 200)
		}

		body := string(w.Body.Bytes())
		if body != "\"world\"\n"{
			t.Errorf("response body was %s; expected %s", body, "world")
		}
	}
}

// TestRestServer_RegisterFunc No hassle developing you application architecture. just write a ServerFunc
// type ServerFunc func(rs *RestServer, c *Context)
// SonicSpeed framework will take that and run it. Focus on Business logic
func TestRestServer_RegisterFunc(t *testing.T) {

	// Setup server
	s := restserver.StandardServer()

	s.MongoConnect("mongodb+srv://dummy:dummy@cluster0-huuez.gcp.mongodb.net", nil)
	if s.MongoDB.Client == nil {
		t.Fatalf("failed to connect to mongodb")
	}

	fn := func(rs *restserver.RestServer, ctx *restserver.Context){
		ctx.Log().Info("its 3PM")
		databases, err := rs.MongoDB.Client.ListDatabaseNames(ctx, bson.M{})
		if err != nil {
			ctx.Log().Errorf("can not fetch db list: %v", err)
			return
		}
		ctx.Rest.JSON(200, databases)
	}

	s.RegisterFunc("test", fn)

	// Set up Request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/test", nil)
	req.Header.Add("User-Agent", "Test")
	s.Gin.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("response code was %d; expected %d", w.Code, 200)
	}

	body := string(w.Body.Bytes())
	if strings.HasPrefix(body, "\"[" ) {
		t.Errorf("response body was %s; expected an array", body)
	}
}

