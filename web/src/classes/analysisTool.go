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
	call string `db:"call"`
}

func NewAnalysisTool(analysisTool_id int, name string, call string) *AnalysisTool {
	return &AnalysisTool{analysisTool_id: analysisTool_id, name: name, call: call}
}

func (a *AnalysisTool) Name() string {
	return a.name
}

func (a *AnalysisTool) SetName(name string) {
	a.name = name
}

func (a *AnalysisTool) Call() string {
	return a.call
}

func (a *AnalysisTool) SetCall(call string) {
	a.call = call
}

func (a *AnalysisTool) AnalysisTool_id() int {
	return a.analysisTool_id
}

func (a *AnalysisTool) SetAnalysisTool_id(analysisTool_id int) {
	a.analysisTool_id = analysisTool_id
}

