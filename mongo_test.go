package sonicspeed

import (
	"context"
	"github.com/gabrielpsilva/sonicspeed/restserver"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

// TestRestServer_Mongo DBMongo component works independently. No Need to use StandardServer()
func TestRestServer_Mongo(t *testing.T) {

	//s := restserver.StandardServer()
	s := restserver.RestServer{}

	s.MongoConnect("mongodb+srv://dummy:dummy@cluster0-huuez.gcp.mongodb.net", nil)
	if s.MongoDB.Client == nil {
		t.Fatalf("failed to connect to mongodb")
	}

	dbs, err := s.MongoDB.Client.ListDatabaseNames(context.Background(), bson.M{} )
	if err != nil {
		t.Errorf("fail to list databases: %v", err)
		return
	}
	t.Log("listing databases: ", dbs)
}
