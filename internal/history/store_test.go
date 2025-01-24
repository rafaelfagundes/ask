package history

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestStore(t *testing.T) {
	tmpDB := "test_store.db"
	defer os.Remove(tmpDB)

	db, err := sql.Open("sqlite3", tmpDB)
	if err != nil {
		t.Fatalf("Failed to open DB: %v", err)
	}
	defer db.Close()

	s, err := NewStore(db)
	if err != nil {
		t.Fatalf("Failed to create store: %v", err)
	}
	if err := s.Save("testQ", "testR"); err != nil {
		t.Errorf("Save failed: %v", err)
	}
	entries, err := s.List()
	if err != nil || len(entries) == 0 {
		t.Errorf("Expected at least one history entry, got %v, err: %v", len(entries), err)
	}
}

func TestStoreMultipleEntries(t *testing.T) {
	tmpDB := "test_store_multiple.db"
	defer os.Remove(tmpDB)

	db, err := sql.Open("sqlite3", tmpDB)
	if err != nil {
		t.Fatalf("Failed to open DB: %v", err)
	}
	defer db.Close()

	s, err := NewStore(db)
	if err != nil {
		t.Fatalf("Failed to create store: %v", err)
	}

	testData := []struct {
		q, r string
	}{
		{"q1", "r1"},
		{"q2", "r2"},
		{"", "r3"}, // empty question
		{"q4", ""}, // empty response
	}

	for _, td := range testData {
		if err := s.Save(td.q, td.r); err != nil {
			t.Errorf("Save failed for %v: %v", td, err)
		}
	}

	entries, err := s.List()
	if err != nil {
		t.Errorf("List failed: %v", err)
	}
	if len(entries) != len(testData) {
		t.Errorf("Expected %d entries, got %d", len(testData), len(entries))
	}
}

func TestStoreInitFailure(t *testing.T) {
	db, err := sql.Open("sqlite3", "/nonexistent/path/db.sqlite")
	if err != nil {
		t.Fatalf("Failed to open DB: %v", err)
	}
	defer db.Close()

	_, err = NewStore(db)
	if err == nil {
		t.Error("Expected error when initializing store with invalid DB")
	}
}

func TestListEmptyStore(t *testing.T) {
	tmpDB := "test_store_empty.db"
	defer os.Remove(tmpDB)

	db, err := sql.Open("sqlite3", tmpDB)
	if err != nil {
		t.Fatalf("Failed to open DB: %v", err)
	}
	defer db.Close()

	s, err := NewStore(db)
	if err != nil {
		t.Fatalf("Failed to create store: %v", err)
	}

	entries, err := s.List()
	if err != nil {
		t.Errorf("List failed: %v", err)
	}
	if len(entries) != 0 {
		t.Errorf("Expected empty list, got %d entries", len(entries))
	}
}
