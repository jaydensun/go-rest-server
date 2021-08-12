package main

import (
	"fmt"
	restful "github.com/emicklei/go-restful/v3"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	ws := new(restful.WebService)
	//ws.
	//	Consumes(restful.MIME_XML, restful.MIME_JSON).
	//	Produces(restful.MIME_JSON, restful.MIME_XML)

	ws.Route(ws.GET("/plus/{a}/{b}").To(findUser).
		Doc("plus two number").
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")))

	ws.Route(ws.POST("/echo").To(echo))

	restful.Add(ws)
	args := os.Args
	var port = "8080"
	if len(args) > 1 {
		port = args[1]
	}
	log.Fatal(http.ListenAndServe(":" + port, nil))

}

func echo(request *restful.Request, response *restful.Response) {
	body, _ := ioutil.ReadAll(request.Request.Body)
	fmt.Println(string(body))
	response.Write(body)
}

func findUser(request *restful.Request, response *restful.Response) {
	a, _ := strconv.Atoi(request.PathParameter("a"))
	b, _ := strconv.Atoi(request.PathParameter("b"))
	log.Println(time.Now(), a , b)
	response.WriteAsJson(a + b)
}