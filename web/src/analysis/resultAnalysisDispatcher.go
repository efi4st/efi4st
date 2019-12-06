/**
 * Author:    Admiral Helmut
 * Created:   28.11.2019
 *
 * (C)
 **/

package analysis

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"strings"
	"../dbprovider"
	"time"
)

type ResultAnalysisDispatcher interface {
	DispatchResult(source string, result string, firmwareId int) error
	AddRelevantApp(source string, result string, firmwareId int) error
}

type resultAnalysisDispatcher struct {
	version int
}

var rAD ResultAnalysisDispatcher
func GetResultAnalysisDispatcher() ResultAnalysisDispatcher { return rAD }

func init() {
	version := 1
	rAD = &resultAnalysisDispatcher{version: version}
}

func (rAD *resultAnalysisDispatcher) DispatchResult(source string, result string, firmwareId int) (err error) {

	switch source {
	case "CronJobSearch":
		err = rAD.analyzeCronJobSearch(result, firmwareId)
	case "ExecutableFinder":
		err = rAD.analyzeExecutableFinder(result, firmwareId)
	case "BashInitAnalysis":
		err = rAD.analyzeBashSearch(result, firmwareId)
	default:

	}

	return err
}

func (rAD *resultAnalysisDispatcher) AddRelevantApp(source string, result string, firmwareId int) (err error) {

	return err
}

func (rAD *resultAnalysisDispatcher) analyzeBashSearch(result string, firmwareId int) (err error) {
	for _, line := range strings.Split(strings.TrimSuffix(result, "\n"), "\n") {
		if(len(string(line))>2 && string(line[0]) != "#") {
				path := strings.Split(line, " ")
				for _, v := range path {
					if(len(string(v)) > 3){
						if (string(v[0]) == "/" && len(string(v)) > 3) {
							if (!strings.Contains(string(v), "bash.bashrc")) {
								id := dbprovider.GetDBManager().GetRelevantAppByPath(v, firmwareId)
								if (id == 0) {
									lastIndex := strings.LastIndex(path[1], "/")
									name := path[1][lastIndex+1 : len(path[1])]
									dbprovider.GetDBManager().AddRelevantApp(name, v, 0, "", "", firmwareId)
									id := dbprovider.GetDBManager().GetRelevantAppByPath(v, firmwareId)
									dbprovider.GetDBManager().UpdateRelevantApp("moduleBash", strconv.Itoa(id))
								}
							}
						}
					}
				}

		}
	}
	return err
}

func (rAD *resultAnalysisDispatcher) analyzeCronJobSearch(result string, firmwareId int) (err error) {
	for _, line := range strings.Split(strings.TrimSuffix(result, "\n"), "\n") {
		if(len(string(line))>2) {
			i, err := strconv.Atoi(string(line[0]))
			i=i+1
			if (err == nil) {
				path := strings.Split(line, " ")
				for _, v := range path {
					if(len(string(v)) > 3){
					if (string(v[0]) == "/" && len(string(v)) > 3) {
						if(!strings.Contains(string(v), "cron")){
						id := dbprovider.GetDBManager().GetRelevantAppByPath(v, firmwareId)
							if (id == 0) {
								lastIndex := strings.LastIndex(path[1],"/")
								name := path[1][lastIndex+1:len(path[1])]
								dbprovider.GetDBManager().AddRelevantApp(name, v, 0, "", "", firmwareId)
								id := dbprovider.GetDBManager().GetRelevantAppByPath(v, firmwareId)
								dbprovider.GetDBManager().UpdateRelevantApp("moduleCronJob",strconv.Itoa(id))
							}
						}
					}
				  }
				}
			}
		}
	}
	return err
}

func (rAD *resultAnalysisDispatcher) analyzeExecutableFinder(result string, firmwareId int) (err error) {
	i := 0
	for _, line := range strings.Split(strings.TrimSuffix(result, "\n"), "\n") {
		path := strings.Split(line, "../../working/filesystem")

		id := dbprovider.GetDBManager().GetRelevantAppByPath(path[1], firmwareId)
		fmt.Printf(strconv.Itoa(id))
		if(id == 0){
			lastIndex := strings.LastIndex(path[1],"/")
			name := path[1][lastIndex+1:len(path[1])]
			dbprovider.GetDBManager().AddRelevantApp(name, path[1], 0, "", "", firmwareId)
			id := dbprovider.GetDBManager().GetRelevantAppByPath(path[1], firmwareId)
			dbprovider.GetDBManager().UpdateRelevantApp("moduleDefault",strconv.Itoa(id))
		}
		i = i + 1
		if(i%15==0){
			time.Sleep(2 * time.Second)
		}

	}
	return err
}
