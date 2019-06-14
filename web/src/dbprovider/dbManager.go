/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package dbprovider

import (
	"../classes"
)

type Manager interface {
	AddProject(project *classes.Project) error
	// Add other methods
}