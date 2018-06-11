package modules

import (
	"github.com/ThomasFPienaar/collaborative/application/httplayer"
	"github.com/ThomasFPienaar/collaborative/modules/syncdata"
)

func LoadModules() {
	httplayer.AddEndpoint("synclist", ModuleSyncData{syncdata.SyncWorker{make([]interface{}, 0)}})
}