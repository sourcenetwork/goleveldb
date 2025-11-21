package table

import (
	"testing"

	"github.com/sourcenetwork/goleveldb/leveldb/testutil"
)

func TestTable(t *testing.T) {
	testutil.RunSuite(t, "Table Suite")
}
