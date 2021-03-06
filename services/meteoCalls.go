package services

import (
	"github.com/deeper-x/goship/lib/ldb"
	"github.com/kataras/iris/v12"
)

// ActiveStations todo doc
func (objPortinformer Portinformer) ActiveStations(ctx iris.Context) {
	activeStations := ldb.GetActiveStations()

	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(activeStations)
}
