package syncdata

// Action for the actions that can be applied to the data
type Action int

// Actions for modifying the data
const (
	ADD Action = iota
	EDIT
	DELETE
)

// SyncActionImpl for syncing data
type SyncActionImpl struct {
	action Action
	data   interface{}
}

// GetAction the syncaction that should be applied
func (sa SyncActionImpl) GetAction() Action {
	return sa.action
}

// GetData returns the data
func (sa SyncActionImpl) GetData() interface{} {
	return sa.data
}
