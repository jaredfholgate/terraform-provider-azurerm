package frontdoor

import (
	"sort"
	"github.com/Azure/azure-sdk-for-go/services/frontdoor/mgmt/2020-01-01/frontdoor"
)

type ById []frontdoor.FrontendEndpoint

func (a ById) Len() int           { return len(a) }
func (a ById) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ById) Less(i, j int) bool { 
	firstId := *a[i].ID
	secondId := *a[j].ID
	return firstId < secondId
}

func Sort(endpoints *[]frontdoor.FrontendEndpoint) {
	sort.Sort(ById(*endpoints))
}

