package services

import (
	"flag"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestMain(m *testing.M) {
	os.Setenv("DB_URL", "root:mysql123@/lademocracia?parseTime=True")

	flag.Parse()
	os.Exit(m.Run())
}

func TestFetchFromSQL(t *testing.T) {
	var ids []int64

	ids = append(ids, 14244)
	ids = append(ids, 14248)

	resp := fetchFromSQL(ids)

	if len(resp) < 2 {
		t.Logf("count %v", len(resp))
		t.Fail()
	}
}
