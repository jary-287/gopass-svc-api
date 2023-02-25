package v1

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/jary-287/gopass-svc-api/pkg"
	"github.com/jary-287/gopass-svc-api/pkg/e"
	"github.com/jary-287/gopass-svc-api/service"
	"github.com/jary-287/gopass-svc/proto/svc"
)

func GetAllSvc(c *gin.Context) {
	app := pkg.Gin{C: c}
	svcs, err := service.Client.FindAllSvc(context.Background(), &svc.FindAll{})
	if err != nil {
		log.Println(err)
		app.Response(http.StatusOK, e.ERROR, nil)
		return
	}

	app.Response(http.StatusOK, e.SUCCESS, svcs.SvcInfo)
}

func UpdateSvc(c *gin.Context) {
	var svcInfo svc.SvcInfo
	app := pkg.Gin{C: c}
	if err := jsonpb.Unmarshal(c.Request.Body, &svcInfo); err != nil {
		log.Println("解析失败：", err)
		app.Response(http.StatusOK, e.INVALID_PARAMS, err)
		return
	}
	rsp, err := service.Client.UpdateSvc(context.Background(), &svcInfo)
	if err != nil {
		log.Println(err)
		app.Response(http.StatusOK, e.ERROR, err)
		return
	}
	app.Response(http.StatusOK, e.SUCCESS, rsp.GetMsg())
}
func AddSvc(c *gin.Context) {
	var svcInfo svc.SvcInfo
	app := pkg.Gin{C: c}
	if err := jsonpb.Unmarshal(c.Request.Body, &svcInfo); err != nil {
		log.Println("解析失败：", err)
		app.Response(http.StatusOK, e.INVALID_PARAMS, err)
		return
	}
	rsp, err := service.Client.AddSvc(context.Background(), &svcInfo)
	if err != nil {
		log.Println(err)
		app.Response(http.StatusOK, e.INVALID_PARAMS, err)
		return
	}
	app.Response(http.StatusOK, e.SUCCESS, rsp.GetMsg())
}
func DeleteSvc(c *gin.Context) {
	app := pkg.Gin{C: c}
	var svcInfo svc.SvcInfo
	if err := jsonpb.Unmarshal(app.C.Request.Body, &svcInfo); err != nil {
		log.Println(err)
		app.Response(http.StatusOK, e.INVALID_PARAMS, err)
	}

	if rsp, err := service.Client.DeleteSvc(context.Background(), &svcInfo); err != nil {
		log.Println(err)
		app.Response(http.StatusOK, e.ERROR_NOT_EXIST_SERVICE, err)
	} else {
		app.Response(http.StatusOK, e.SUCCESS, rsp.Msg)
	}
}
func GetSvcByID(c *gin.Context) {
	app := pkg.Gin{C: c}
	id, err := strconv.Atoi(app.C.Param("id"))
	if err != nil {
		log.Println(err)
		app.Response(http.StatusOK, e.INVALID_PARAMS, e.GetMsg(e.INVALID_PARAMS))
	}
	svcInfo, err := service.Client.FindSvcById(context.Background(), &svc.SvcId{Id: uint64(id)})
	if err != nil {
		log.Println(err)
		app.Response(http.StatusOK, e.ERROR, err)
		return
	}
	app.Response(http.StatusOK, e.SUCCESS, svcInfo)

}
