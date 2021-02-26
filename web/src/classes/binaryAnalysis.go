/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type BinaryAnalysis struct {
	binaryAnalysis_id int `db:"binaryAnalysis_id"`
	toolOutput string `db:"toolOutput"`
	analysisTool_id int `db:"analysisTool_id"`
	relevantApps_id int `db:"relevantApps_id"`
}

func (b *BinaryAnalysis) AnalysisTool_id() int {
	return b.analysisTool_id
}

func (b *BinaryAnalysis) SetAnalysisTool_id(analysisTool_id int) {
	b.analysisTool_id = analysisTool_id
}

func (b *BinaryAnalysis) ToolOutput() string {
	return b.toolOutput
}

func (b *BinaryAnalysis) SetToolOutput(toolOutput string) {
	b.toolOutput = toolOutput
}

func (b *BinaryAnalysis) BinaryAnalysis_id() int {
	return b.binaryAnalysis_id
}

func (b *BinaryAnalysis) SetBinaryAnalysis_id(binaryAnalysis_id int) {
	b.binaryAnalysis_id = binaryAnalysis_id
}

func (b *BinaryAnalysis) RelevantApps_id() int {
	return b.relevantApps_id
}

func (b *BinaryAnalysis) SetRelevantApps_id(relevantApps_id int) {
	b.relevantApps_id = relevantApps_id
}

func NewBinaryAnalysis(binaryAnalysis_id int, toolOutput string, analysisTool_id int, relevantApps_id int) *BinaryAnalysis {
	return &BinaryAnalysis{binaryAnalysis_id: binaryAnalysis_id, toolOutput: toolOutput, analysisTool_id: analysisTool_id, relevantApps_id: relevantApps_id}
}

