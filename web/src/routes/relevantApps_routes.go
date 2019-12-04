/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package routes

import (
	"../dbprovider"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris"
	"strconv"
	"strings"
)

func RelevantApps(ctx iris.Context) {

	relevantApps := dbprovider.GetDBManager().GetRelevantApps()

	ctx.ViewData("error", "")

	if len(relevantApps) < 1 {
		ctx.ViewData("error", "Error: No apps available. Add one!")
	}

	ctx.ViewData("relevantAppsList", relevantApps)
	ctx.View("relevantApps.html")
}

func ShowRelevantApp(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error parsing firmware Id!")
	}

	relevantApp := dbprovider.GetDBManager().GetRelevantAppInfo(i)

	ctx.ViewData("relevantApp", relevantApp)
	ctx.View("showRelevantApp.html")
}


func RemoveRelevantApp(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)
	err = dbprovider.GetDBManager().RemoveRelevantApp(i)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error removing project!")
	}

	relevantApps := dbprovider.GetDBManager().GetRelevantApps()

	ctx.ViewData("relevantAppsList", relevantApps)
	ctx.View("relevantApps.html")
}

// POST
func AddRelevantApp(ctx iris.Context) {

	id := ctx.Params().Get("project_id")
	relevantAppByHand := ctx.PostValue("RelevantAppByHand")
	i, err := strconv.Atoi(id)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error parsing project Id!")
	}

	lastIndex := strings.LastIndex(relevantAppByHand,"/")
	name := relevantAppByHand[lastIndex+1:len(relevantAppByHand)]
	dbprovider.GetDBManager().AddRelevantApp(name, relevantAppByHand, 0, "", "", i)

	ctx.Redirect("/firmware/show/"+id)
}





