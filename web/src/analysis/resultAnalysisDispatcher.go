/**
 * Author:    Admiral Helmut
 * Created:   28.11.2019
 *
 * (C)
 **/

package analysis

import (
	"../dbprovider"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"strings"
	"time"
)

type ResultAnalysisDispatcher interface {
	DispatchResult(source string, result string, firmwareId int, relevantApp_path string) error
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

func (rAD *resultAnalysisDispatcher) DispatchResult(source string, result string, firmwareId int, relevantApp_path string) (err error) {

	switch source {
	case "CronJobSearch":
		err = rAD.analyzeCronJobSearch(result, firmwareId)
	case "ExecutableFinder":
		err = rAD.analyzeExecutableFinder(result, firmwareId)
	case "BashInitAnalysis":
		err = rAD.analyzeBashSearch(result, firmwareId)
	case "AppChainAnalysis":
		err = rAD.analyzeAppChain(result, firmwareId, relevantApp_path)
	case "BinWalkAnalysis":
		err = rAD.analyzeBinWalk(result, firmwareId, relevantApp_path)
	case "InitSystemAnalysis":
		err = rAD.analyzeInitSystem(result, firmwareId)
	case "PSLocalProcesses":
		err = rAD.analyzeLocalProcesses(result, firmwareId)
	case "NetstatLocal":
		err = rAD.analyzeLocalInterfaces(result, firmwareId)
	case "LocalSystemServices":
		err = rAD.analyzeLocalSystemServices(result, firmwareId)
	case "Binary4EmuAnalysisreadelf":
		err = rAD.analyzeBinary4Emureadelf(result, firmwareId, relevantApp_path)
	case "Binary4EmuAnalysisldd":
		err = rAD.analyzeBinary4Emuldd(result, firmwareId, relevantApp_path)
	case "Binary4EmuAnalysisstrace":
		err = rAD.analyzeBinary4Emustrace(result, firmwareId, relevantApp_path)
	case "NMAPNetworkInterfaces":
		err = rAD.analyzeNMAPNetworkInterfaces(result, firmwareId)
	case "SimpleHTTPTest":
		err = rAD.analyzeSimpleHTTPTest(result, firmwareId)
	case "SystemCTL":
		err = rAD.analyzeSystemCtl(result, firmwareId)
	case "ListProc":
		err = rAD.analyzeProc(result, firmwareId)
	case "NetstatLocalInterfaces":
		err = rAD.analyzeNetstatInterfaces(result, firmwareId)
	case "LSOF":
		err = rAD.analyzeLSOF(result, firmwareId)
	case "NetworkHandles":
		err = rAD.analyzeNetworkHandles(result, firmwareId)
	default:

	}

	return err
}

func (rAD *resultAnalysisDispatcher) analyzeNetworkHandles(result string, firmwareId int) (err error) {

	return err
}

func (rAD *resultAnalysisDispatcher) analyzeLSOF(result string, firmwareId int) (err error) {

	return err
}

func (rAD *resultAnalysisDispatcher) analyzeNetstatInterfaces(result string, firmwareId int) (err error) {

	return err
}

func (rAD *resultAnalysisDispatcher) analyzeProc(result string, firmwareId int) (err error) {

	return err
}

func (rAD *resultAnalysisDispatcher) analyzeSystemCtl(result string, firmwareId int) (err error) {

	return err
}

func (rAD *resultAnalysisDispatcher) analyzeSimpleHTTPTest(result string, firmwareId int) (err error) {
	fmt.Printf("UUUUUUUUUUUuuu")
	protocol := ""
	port := 0
	app := ""
	target := false
	for _, line := range strings.Split(strings.TrimSuffix(result, "\n"), "\n") {
		if(len(line)>3){
			if(strings.Contains(line, "HTTP") && strings.Contains(line, "200") && strings.Contains(line, "OK") ){
				protocol = "HTTP"
				port = 80
				target=true
			}
			if(strings.Contains(line, "Server")){
				serverParts := strings.Split(line, ":")
				s := strings.TrimSpace(serverParts[1])
				sP := strings.Split(s, "/")
				app = sP[0]
				fmt.Printf("LLLLLLLLLLLLLLLL"+app+"LLLLLLLLLLLll")
				target = true
			}
			if(target== true){
				id := dbprovider.GetDBManager().GetRelevantAppByName(app, firmwareId)
				if (id == 0) {
					fmt.Printf("AAAAAAA")
					dbprovider.GetDBManager().AddRelevantApp(app, "", 0, protocol, "", firmwareId)
				}
				id = dbprovider.GetDBManager().GetRelevantAppByName(app, firmwareId)
				fmt.Printf("OOOOOOOOOOOOOO"+strconv.Itoa(id))
				dbprovider.GetDBManager().UpdateRelevantAppForInterface("moduleProtocolls",strconv.Itoa(id), port, protocol)
			}
			target=false

		}
	}
	return err
}

func (rAD *resultAnalysisDispatcher) analyzeNMAPNetworkInterfaces(result string, firmwareId int) (err error) {
	for _, line := range strings.Split(strings.TrimSuffix(result, "\n"), "\n") {
		if(len(line)>3){
			i, err := strconv.Atoi(string(line[0]))
			i=i+1
			if (err == nil) {
				path := strings.Split(line, " ")

				for _, v := range path {
					if(len(string(v)) > 3){
							id := dbprovider.GetDBManager().GetRelevantAppByName(v, firmwareId)
							if (id == 0) {
								name := string(v)
								dbprovider.GetDBManager().AddRelevantApp(name, v, 0, "", "", firmwareId)
							}
							dbprovider.GetDBManager().UpdateRelevantApp("modulePortscanner",strconv.Itoa(id))


					}
				}
			}
		}
	}
	return err
}

func (rAD *resultAnalysisDispatcher) AddRelevantApp(source string, result string, firmwareId int) (err error) {

	return err
}

func (rAD *resultAnalysisDispatcher) analyzeBinWalk(result string, firmwareId int, relevantApp_path string) (err error) {

	orgPathIndex := strings.Index(relevantApp_path, "working/filesystem")
	orgPath := relevantApp_path[orgPathIndex+len("working/filesystem"):len(relevantApp_path)]

	appcontent := dbprovider.GetDBManager().GetAppContentForRelevantAppByPath(orgPath)
	if (appcontent == nil) {
		dbprovider.GetDBManager().AddAppContent("", result, "", "", "",  orgPath)
	}else{
		dbprovider.GetDBManager().UpdateAppContent(appcontent.AppContent_id(), "binwalk", result)
	}

	return err
}

func (rAD *resultAnalysisDispatcher) analyzeAppChain(result string, firmwareId int, relevantApp_path string) (err error) {
	i := strings.Index(result, "[")

	orgPathIndex := strings.Index(relevantApp_path, "working/filesystem")
	orgPath := relevantApp_path[orgPathIndex+len("working/filesystem"):len(relevantApp_path)]

	dbprovider.GetDBManager().RemoveAppContentByRelevantAppPath(orgPath)
	dbprovider.GetDBManager().AddAppContent(result[i : len(result)], "", "", "", "", orgPath)

	return err
}

func (rAD *resultAnalysisDispatcher) analyzeBashSearch(result string, firmwareId int) (err error) {
	for _, line := range strings.Split(strings.TrimSuffix(result, "\n"), "\n") {
		line =  strings.TrimSpace(line)
		if(len(string(line))>2 && string(line[0]) != "#") {
				path := strings.Split(line, " ")
				for _, v := range path {
					if(len(string(v)) > 3){
						fmt.Printf("F")
						if (string(v[0]) == "/" && len(string(v)) > 3) {
							if (!strings.Contains(string(v), "bash.bashrc")) {
								if(string(v[len(v)-1]) == ")"){
									v = v[0:len(v)-1]
								}
								id := dbprovider.GetDBManager().GetRelevantAppByPath(v, firmwareId)
								if (id == 0) {
									lastIndex := strings.LastIndex(v, "/")
									name := v[lastIndex+1 : len(v)]
									dbprovider.GetDBManager().AddRelevantApp(name, v, 0, "", "", firmwareId)
								}
								id = dbprovider.GetDBManager().GetRelevantAppByPath(v, firmwareId)
								dbprovider.GetDBManager().UpdateRelevantApp("moduleBash", strconv.Itoa(id))

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
							}
							id = dbprovider.GetDBManager().GetRelevantAppByPath(v, firmwareId)
							dbprovider.GetDBManager().UpdateRelevantApp("moduleCronJob",strconv.Itoa(id))
						}
					}
				  }
				}
			}
		}
	}
	return err
}

func (rAD *resultAnalysisDispatcher) analyzeInitSystem(result string, firmwareId int) (err error) {
	for _, line := range strings.Split(strings.TrimSuffix(result, "\n"), "\n") {
		if(len(line)>3){
			if(string(line[0])=="+"){
				id := dbprovider.GetDBManager().GetRelevantAppByPath(line[1:len(line)], firmwareId)
				if (id == 0) {
					lastIndex := strings.LastIndex(line[1:len(line)],"/")
					name := line[1:len(line)][lastIndex+1:len(line[1:len(line)])]
					dbprovider.GetDBManager().AddRelevantApp(name, line[1:len(line)], 0, "", "", firmwareId)
				}
				id = dbprovider.GetDBManager().GetRelevantAppByPath(line[1:len(line)], firmwareId)
				dbprovider.GetDBManager().UpdateRelevantApp("moduleInitSystem",strconv.Itoa(id))
			}
		}
	}
	return err
}

func (rAD *resultAnalysisDispatcher) analyzeLocalProcesses(result string, firmwareId int) (err error) {
	for _, line := range strings.Split(strings.TrimSuffix(result, "\n"), "\n") {
		if(len(line)>3){
			path := strings.Split(line, " ")

			for _, v := range path {
				if(len(string(v)) > 3){
					if (string(v[0]) == "/" && len(string(v)) > 3) {
							id := dbprovider.GetDBManager().GetRelevantAppByPath(v, firmwareId)
							if (id == 0) {
								lastIndex := strings.LastIndex(path[1],"/")
								name := path[1][lastIndex+1:len(path[1])]
								dbprovider.GetDBManager().AddRelevantApp(name, v, 0, "", "", firmwareId)
							}
							id = dbprovider.GetDBManager().GetRelevantAppByPath(v, firmwareId)
							dbprovider.GetDBManager().UpdateRelevantApp("moduleProcesses",strconv.Itoa(id))

					}
				}
			}
		}
	}
	return err
}

func (rAD *resultAnalysisDispatcher) analyzeLocalInterfaces(result string, firmwareId int) (err error) {
	for _, line := range strings.Split(strings.TrimSuffix(result, "\n"), "\n") {
		if(len(line)>3){
			if(string(line[0])=="+"){
				id := dbprovider.GetDBManager().GetRelevantAppByPath(line[1:len(line)], firmwareId)
				if (id == 0) {
					lastIndex := strings.LastIndex(line[1:len(line)],"/")
					name := line[1:len(line)][lastIndex+1:len(line[1:len(line)])]
					dbprovider.GetDBManager().AddRelevantApp(name, line[1:len(line)], 0, "", "", firmwareId)
				}
				id = dbprovider.GetDBManager().GetRelevantAppByPath(line[1:len(line)], firmwareId)
				dbprovider.GetDBManager().UpdateRelevantApp("moduleInterfaces",strconv.Itoa(id))
			}
		}
	}
	return err
}

func (rAD *resultAnalysisDispatcher) analyzeLocalSystemServices(result string, firmwareId int) (err error) {

	return err
}


func (rAD *resultAnalysisDispatcher) analyzeBinary4Emureadelf(result string, firmwareId int, relevantApp_path string) (err error) {
	orgPathIndex := strings.Index(relevantApp_path, "working/filesystem")
	orgPath := relevantApp_path[orgPathIndex+len("working/filesystem"):len(relevantApp_path)]

	appcontent := dbprovider.GetDBManager().GetAppContentForRelevantAppByPath(orgPath)
	if (appcontent == nil) {
		dbprovider.GetDBManager().AddAppContent("", "", result, "", "",  orgPath)
	}else{
		dbprovider.GetDBManager().UpdateAppContent(appcontent.AppContent_id(), "readelf", result)
	}

	return err
}

func (rAD *resultAnalysisDispatcher) analyzeBinary4Emuldd(result string, firmwareId int, relevantApp_path string) (err error) {
	orgPathIndex := strings.Index(relevantApp_path, "working/filesystem")
	orgPath := relevantApp_path[orgPathIndex+len("working/filesystem"):len(relevantApp_path)]

	appcontent := dbprovider.GetDBManager().GetAppContentForRelevantAppByPath(orgPath)
	if (appcontent == nil) {
		dbprovider.GetDBManager().AddAppContent("", "", "", result, "",  orgPath)
	}else{
		dbprovider.GetDBManager().UpdateAppContent(appcontent.AppContent_id(), "ldd", result)
	}

	return err
}

func (rAD *resultAnalysisDispatcher) analyzeBinary4Emustrace(result string, firmwareId int, relevantApp_path string) (err error) {
	orgPathIndex := strings.Index(relevantApp_path, "working/filesystem")
	orgPath := relevantApp_path[orgPathIndex+len("working/filesystem"):len(relevantApp_path)]

	appcontent := dbprovider.GetDBManager().GetAppContentForRelevantAppByPath(orgPath)
	if (appcontent == nil) {
		dbprovider.GetDBManager().AddAppContent("", "", "", "", result,  orgPath)
	}else{
		dbprovider.GetDBManager().UpdateAppContent(appcontent.AppContent_id(), "strace", result)
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
		}
		id = dbprovider.GetDBManager().GetRelevantAppByPath(path[1], firmwareId)
		dbprovider.GetDBManager().UpdateRelevantApp("moduleDefault",strconv.Itoa(id))
		i = i + 1
		if(i%15==0){
			time.Sleep(2 * time.Second)
		}

	}
	return err
}
