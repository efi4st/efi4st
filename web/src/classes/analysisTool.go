/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type AnalysisTool struct {
	analysisTool_id int `db:"analysisTool_id"`
	name string `db:"name"`
	executionString string `db:"executionString"`
}

func NewAnalysisTool(analysisTool_id int, name string, executionString string) *AnalysisTool {
	return &AnalysisTool{analysisTool_id: analysisTool_id, name: name, executionString: executionString}
}

func (a *AnalysisTool) Name() string {
	return a.name
}

func (a *AnalysisTool) SetName(name string) {
	a.name = name
}

func (a *AnalysisTool) ExecutionString() string {
	return a.executionString
}

func (a *AnalysisTool) SetExecutionString(executionString string) {
	a.executionString = executionString
}

func (a *AnalysisTool) AnalysisTool_id() int {
	return a.analysisTool_id
}

func (a *AnalysisTool) SetAnalysisTool_id(analysisTool_id int) {
	a.analysisTool_id = analysisTool_id
}

