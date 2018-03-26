package syncdata

import (
	"fmt"
)

// SyncAction actions and data
type SyncAction interface {
	GetData() interface{}
	GetAction() Action
}

//SyncWorker synchs some stuff
type SyncWorker struct {
	data []interface{}
}

func (s *SyncWorker) processAction(sa SyncAction) {
	if sa.GetAction() == ADD {
		s.data = append(s.data, sa.GetData())
	}
	// var data = make([]interface{}, 20)
}

// GetDataList Return the data
func (s SyncWorker) GetDataList() []interface{} {
	fmt.Println("data: ", s.data)
	return s.data
}
