package leveldb

import (
	"testing"

	"github.com/sourcenetwork/goleveldb/leveldb/testutil"
)

func TestLevelDB(t *testing.T) {
	testutil.RunSuite(t, "LevelDB Suite")
}
