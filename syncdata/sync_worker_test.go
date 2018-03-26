package syncdata

import (
	"fmt"
	"testing"
)

func TestAddData(t *testing.T) {

	syncWorker := SyncWorker{make([]interface{}, 0)}
	syncAction := SyncActionImpl{ADD, "test"}

	syncWorker.processAction(syncAction)
	dataList := syncWorker.GetDataList()

	fmt.Println("v1: ", dataList)

	if len(dataList) != 1 {
		t.Fail()
	}

	if dataList[0] != "test" {
		t.Fail()
	}

}
