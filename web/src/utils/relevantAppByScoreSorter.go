package utils

/**
 * Author:    Admiral Helmut
 * Created:   04.12.2019
 *
 * (C)
 **/

import (
	"github.com/efi4st/efi4st/classes"
)

// RelevantAppsByScoreSorter implements sort.Interface based on the Age field.
type RelevantAppsByScoreSorter []classes.RelevantApps

func (a RelevantAppsByScoreSorter) Len() int           { return len(a) }
func (a RelevantAppsByScoreSorter) Less(i, j int) bool { return a[i].Msg() > a[j].Msg() }
func (a RelevantAppsByScoreSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
