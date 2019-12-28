package zzexterntest

import (
	"ginRestfulBase/restserver"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServerFunc(t *testing.T) {

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
	body = body
	//if strings.{
	//	t.Errorf("response body was %s; expected %s", body, "world")
	//}


}
