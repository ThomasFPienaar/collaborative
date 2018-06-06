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
	Data []interface{}
}

// AddAction - action to add
func (s *SyncWorker) AddAction(sd SyncData, position int) {
	s.Data = append(s.Data, sd.GetData())
}

// ReplaceAction - action to replace
func (s *SyncWorker) ReplaceAction(sd SyncData, position int) {
	s.Data[position] = sd.GetData()
}

// DeleteAction - action to delete
func (s *SyncWorker) DeleteAction(position int) {
	s.Data = append(s.Data[:position], s.Data[position+1:]...)
}

// MoveAction - action to delete
func (s *SyncWorker) MoveAction(fromPosition int, toPosition int) {
	data := s.Data[fromPosition]
	copy(s.Data[fromPosition+1:], s.Data[fromPosition:])
	copy(s.Data[toPosition:], s.Data[toPosition+1:])

	s.Data[toPosition] = data
}

// GetDataList Return the data
func (s SyncWorker) GetDataList() []interface{} {
	fmt.Println("data: ", s.Data)
	return s.Data
}
