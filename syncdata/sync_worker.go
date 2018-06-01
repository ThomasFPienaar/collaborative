package syncdata

import (
	"fmt"
)

// SyncAction actions and data
type SyncData interface {
	GetData() interface{}
}

type SyncActions interface {
	AddAction(sd SyncData, position int)
	ReplaceAction(sd SyncData, position int)
	DeleteAction(position int)
	MoveAction(fromPosition int, toPosition int)
}

//SyncWorker synchs some stuff
type SyncWorker struct {
	data []interface{}
}

// AddAction - action to add
func (s *SyncWorker) AddAction(sd SyncData, position int) {
	s.data = append(s.data, sd.GetData())
}

// ReplaceAction - action to replace
func (s *SyncWorker) ReplaceAction(sd SyncData, position int) {
	s.data[position] = sd.GetData()
}

// DeleteAction - action to delete
func (s *SyncWorker) DeleteAction(position int) {
	s.data = append(s.data[:position], s.data[position+1:]...)
}

// MoveAction - action to delete
func (s *SyncWorker) MoveAction(fromPosition int, toPosition int) {
	data := s.data[fromPosition]
	copy(s.data[fromPosition+1:], s.data[fromPosition:])
	copy(s.data[toPosition:], s.data[toPosition+1:])

	s.data[toPosition] = data
}

// GetDataList Return the data
func (s SyncWorker) GetDataList() []interface{} {
	fmt.Println("data: ", s.data)
	return s.data
}
