/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package routes

import (
	"github.com/efi4st/efi4st/dbprovider"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/v12"
	"strings"
)

func ShowElementSearchPage(ctx iris.Context) {
	ctx.View("sms_elementSearch.html")
}

func SearchElementsAPI(ctx iris.Context) {
	query := strings.ToLower(strings.TrimSpace(ctx.URLParam("q")))

	if len(query) < 2 {
		ctx.StatusCode(400)
		ctx.JSON(iris.Map{"error": "Bitte mindestens 2 Zeichen eingeben."})
		return
	}

	results := dbprovider.GetDBManager().GetSMSElementSearchLike(query)
	ctx.JSON(results)
}