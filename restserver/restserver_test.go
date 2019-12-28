package restserver

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRestServer_StandardServer(t *testing.T) {

	// Setup server
	s := StandardServer()


	// Function to request
	fn := func(c *gin.Context){
		c.JSON(http.StatusOK, "world")
	}
	// Set up request and function
	s.Gin.Group("/v1").GET("/hello", fn)

	// Set up Request
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

func TestRestServer_Log(t *testing.T) {
	tt := [] struct{
		name 		string
	}{
		{name:"JSON"},
		{name:"TEXT"},
	}

	// Setup server
	s := StandardServer()
	s.AddRequestIDHeader("X-Request-Id", "x-cloud-trace-context")

	// Function to request
	fn := func(c *gin.Context){
		ctx := ContextWrapper(c)
		ctx.Log().Debugf("Hello from Hello World test endpoint")
		c.JSON(http.StatusOK, "world")
	}
	// Set up request and function
	s.Gin.Group("/v1").GET("/hello", fn)


	for _, tc := range tt {

		if tc.name == "JSON" {
			GlobalLogrusFormatJSON(logrus.DebugLevel)
		}
		if tc.name == "TEXT" {
			GlobalLogrusFormatTEXT(logrus.DebugLevel)
		}


		// Set up Request
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


func TestRestServer_Mongo(t *testing.T) {

	// Setup server
	s := StandardServer()

	s.MongoConnect("mongodb+srv://dummy:dummy@cluster0-huuez.gcp.mongodb.net", nil)
	if s.MongoDB.Client == nil {
		t.Fatalf("failed to connect to mongodb")
	}

	var filter map[string]interface{}
	dbs, err := s.MongoDB.Client.ListDatabaseNames(context.Background(), filter )
	if err != nil {
		t.Errorf("fail to list databases: %v", err)
		return
	}
	t.Log("listing databases: ", dbs)
}
