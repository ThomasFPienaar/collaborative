package syncdata

import (
	"fmt"
	"testing"
)

type StringData struct {
	data string
}

func (sd StringData) GetData() interface{} {
	return sd.data
}

func TestAddData(t *testing.T) {

	syncWorker := SyncWorker{make([]interface{}, 0)}

	syncWorker.AddAction(StringData{"test"}, 1)

	dataList := syncWorker.GetDataList()

	fmt.Println("v1: ", dataList)

	if len(dataList) != 1 {
		t.Fail()
	}

	if dataList[0] != "test" {
		t.Fail()
	}

}

func TestEditData(t *testing.T) {

	syncWorker := SyncWorker{make([]interface{}, 0)}

	syncWorker.AddAction(StringData{"test"}, 0)

	syncWorker.ReplaceAction(StringData{"edit"}, 0)
	dataList := syncWorker.GetDataList()

	fmt.Println("v1: ", dataList)

	if len(dataList) != 1 {
		t.Fail()
	}

	if dataList[0] != "edit" {
		t.Fail()
	}

}

func TestDeleteData(t *testing.T) {

	syncWorker := SyncWorker{make([]interface{}, 0)}

	syncWorker.AddAction(StringData{"test"}, 0)
	syncWorker.DeleteAction(0)

	dataList := syncWorker.GetDataList()

	fmt.Println("v1: ", dataList)

	if len(dataList) != 0 {
		t.Fail()
	}
}

func TestDeleteDataIndexOne(t *testing.T) {

	syncWorker := SyncWorker{make([]interface{}, 0)}

	syncWorker.AddAction(StringData{"test1"}, 0)
	syncWorker.AddAction(StringData{"test2"}, 0)

	syncWorker.DeleteAction(1)
	dataList := syncWorker.GetDataList()

	fmt.Println("v1: ", dataList)

	if len(dataList) != 1 {
		t.Fail()
	}

	if dataList[0] != "test1" {
		t.Fail()
	}

}

func TestMoveData(t *testing.T) {

	syncWorker := SyncWorker{make([]interface{}, 0)}

	syncWorker.AddAction(StringData{"test1"}, 0)
	syncWorker.AddAction(StringData{"test2"}, 0)
	syncWorker.AddAction(StringData{"test3"}, 0)

	syncWorker.MoveAction(1, 2)
	dataList := syncWorker.GetDataList()

	fmt.Println("v1: ", dataList)

	if len(dataList) != 3 {
		t.Fail()
	}

	if dataList[0] != "test1" && dataList[1] != "test3" && dataList[2] != "test2" {
		t.Fail()
	}

}
