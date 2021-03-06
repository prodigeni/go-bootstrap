package helpers

import (
	"os"
	"strings"
	"testing"
)

func TestBashEscape(t *testing.T) {
	os.Setenv("PGPASSWORD", "boo")
	dbName := DefaultPGDSN("default")
	escapedDbName := BashEscape(dbName)
	if !strings.Contains(escapedDbName, `\&`) {
		t.Errorf("Incorrectly escape dbName. Output: %v", escapedDbName)
	}
}

func TestRandString(t *testing.T) {
	str := RandString(12)
	if len(str) != 12 {
		t.Errorf("Generated string with incorrect length. Output: %v", str)
	}
}

func TestDefaultPGDSN(t *testing.T) {
	os.Setenv("PGPASSWORD", "boo")
	dbName := DefaultPGDSN("default")
	if !strings.HasPrefix(dbName, `postgres://`) {
		t.Errorf("Incorrectly generate DSN. Output: %v", dbName)
	}
	if !strings.Contains(dbName, "password=boo") {
		t.Errorf("Incorrectly generate DSN. Output: %v", dbName)
	}
}
