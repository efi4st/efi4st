/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package routes

import (
	"encoding/json"
	"github.com/efi4st/efi4st/dbprovider"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/v12"
)

// GET



// POST
func ShowStatistics(ctx iris.Context) {
	stats, err := dbprovider.GetDBManager().GetSystemVersionStatistics()
	if err != nil {
		ctx.StatusCode(500)
		ctx.WriteString("Fehler beim Abrufen der Statistik")
		return
	}

	var systemVersions []string
	var projectCounts []int

	for _, stat := range stats {
		systemVersions = append(systemVersions, stat.SystemVersion)
		projectCounts = append(projectCounts, stat.ProjectCount)
	}

	// Daten als JSON für das Template formatieren
	jsonSystemVersions, _ := json.Marshal(systemVersions)
	jsonProjectCounts, _ := json.Marshal(projectCounts)

	ctx.ViewData("system_versions", string(jsonSystemVersions))  // JSON-String
	ctx.ViewData("project_counts", string(jsonProjectCounts))    // JSON-String
	ctx.View("sms_showStatistics.html")
}

