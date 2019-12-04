/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package routes

import (
	"../dbprovider"
	"../analysis"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func TestResults(ctx iris.Context) {

	testResults := dbprovider.GetDBManager().GetTestResults()

	ctx.ViewData("error", "")

	if len(testResults) < 1 {
		ctx.ViewData("error", "Error: No apps available. Add one!")
	}

	ctx.ViewData("testResultsList", testResults)
	ctx.View("testResults.html")
}

func ShowTestResult(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error parsing firmware Id!")
	}

	testResult := dbprovider.GetDBManager().GetTestResultInfo(i)
	testResult.SetResult( strings.Replace(testResult.Result(), "../../working/filesystem", "[testtarget] -> ", -1))

	fmt.Printf(testResult.Result())

	ctx.ViewData("testResult", testResult)
	ctx.View("showTestResult.html")
}


func RemoveTestResult(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)
	err = dbprovider.GetDBManager().RemoveTestResult(i)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error removing test Result!")
	}

	testResults := dbprovider.GetDBManager().GetTestResults()

	ctx.ViewData("testResultsList", testResults)
	ctx.View("testResults.html")
}

type TestResultMsg struct {
	Result string `json:"result"`
	Source string `json:"source"`
}

// POST
func AddResultSet(ctx iris.Context) {

	id := ctx.Params().Get("project_id")
	dt := time.Now()
	i, err := strconv.Atoi(id)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error parsing project Id!")
	}

	body, _ := ioutil.ReadAll(ctx.Request().Body)
	result := TestResultMsg{}
	json.Unmarshal([]byte(body), &result)

	dbprovider.GetDBManager().AddTestResult(result.Source, result.Result, dt, i)

	err = analysis.GetResultAnalysisDispatcher().DispatchResult(result.Source, result.Result, i)

	ctx.Writef("Result set received!")
}



