package restserver

import (
	"ginRestfulBase/dbmongo"
	"ginRestfulBase/restserver/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type ServerFunc func(*RestServer, *Context)

type RestServer struct {
	Gin 	*gin.Engine 		// direct access to http server
	MongoDB *dbmongo.DBMongo	// direct access to mongo client

	reqIDHeaders []string
}



func  StandardServer() *RestServer {

	s := &RestServer{}

	s.Gin = gin.New()
	s.Gin.Use(
		gin.Recovery(),
		middleware.RequestIDWithOpts(s.reqIDHeaders),
		middleware.GinReturnLog(logrus.StandardLogger()),
		)
	return s
}

func (s *RestServer) AddRequestIDHeader(headers ...string){
	s.reqIDHeaders = append(s.reqIDHeaders, headers...)
}

func (s *RestServer) RegisterFunc(path string, serverFunc ServerFunc){

	// Function to request
	fn := func(c *gin.Context){
		serverFunc(s, ContextWrapper(c))
	}
	// Set up request and function
	s.Gin.Group("/v1").GET(path, fn)


}

func (s *RestServer) RunServer(){

	if s.Gin == nil {
		log.Fatal("server not configured")
	}
	s.Gin.Run(":8080")

}


func (s *RestServer) MongoConnect(connUrl string, clientOptions *options.ClientOptions) {

	s.MongoDB = dbmongo.Connect(connUrl, clientOptions)
	if s.MongoDB.Client == nil {
		logrus.Errorf("failed to connect to mongodb")
	}
}
