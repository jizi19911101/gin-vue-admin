package sync

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestParseCase(t *testing.T) {
	dir, _ := ioutil.TempDir("./", "temp_*")
	defer os.RemoveAll(dir)
	syncService := SyncService{}
	syncService.ParseApiTestcase(dir)

}
