package modules

import (
	"net/http"
	"github.com/thomas/collaborative/modules/syncdata"
)

type ModuleSyncData struct {
	syncList syncdata.SyncWorker
}

func (msd ModuleSyncData) Handler(r *http.Request) interface{} {
	return msd.syncList.GetDataList()
}