package full

import (
	"testing"
)

func TestMigrateSchema(t *testing.T) {
	err := MigrateSchema("../../../../test-db.db")
	if err != nil {
		t.Error(err)
	}
}
