package restserver

import (
	"github.com/gabrielpsilva/sonicspeed/dbmongo"
	"github.com/gabrielpsilva/sonicspeed/restserver/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"regexp"
	"strings"
)

type ServerFunc func(rs *RestServer, c *Context)

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

func (s *RestServer) RegisterFunc(group, path, httpMethod string, serverFunc ServerFunc) {
	fn := func(c *gin.Context){
		serverFunc(s, ContextWrapper(c))
	}
	method := strings.ToUpper(httpMethod)
	reg, _ := regexp.Compile("[^A-Z]+")
	method = reg.ReplaceAllString(method, "")
	switch method {
	case "GET":
		s.Gin.Group(group).GET(path, fn)
	case "POST":
		s.Gin.Group(group).POST(path, fn)
	case "HEAD":
		s.Gin.Group(group).HEAD(path, fn)
	case "OPTIONS":
		s.Gin.Group(group).OPTIONS(path, fn)
	case "PATCH":
		s.Gin.Group(group).PATCH(path, fn)
	case "DELETE":
		s.Gin.Group(group).DELETE(path, fn)
	case "PUT":
		s.Gin.Group(group).PUT(path, fn)
	case "ANY":
		s.Gin.Group(group).Any(path, fn)
	}
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
