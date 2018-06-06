package modules

import (
	"github.com/thomas/collaborative/application/httplayer"
	"github.com/thomas/collaborative/modules/syncdata"
)

func LoadModules() {
	httplayer.AddEndpoint("synclist", ModuleSyncData{syncdata.SyncWorker{make([]interface{}, 0)}})
}