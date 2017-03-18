package golinedb

import (
	"github.com/carlosdavidepto/golinedb/iface"
	"testing"
)

var _ iface.DB = (*DB)(nil)

func getEmptyDB() *DB {
	return &DB{}
}

func getDB3() *DB {
	return &DB{[]string{
		"test1",
		"test2",
		"test3",
	}}
}

func TestLength(t *testing.T) {
	empty := getEmptyDB()
	three := getDB3()
	msg := "DB::Length: got %v, expected %v."

	if empty.Length() != 0 {
		t.Errorf(msg, empty.Length(), 0)
	}
	if three.Length() != 3 {
		t.Errorf(msg, three.Length(), 3)
	}
}

func TestUpdateNonexistentLine(t *testing.T) {
	db := getDB3()
	msg := "DB::Update: got %v, expected %v."

	err := db.Update(4, "new 4")

	if err == nil || err.Error() != "Index 4 does not exist." {
		t.Errorf(msg, err, "Index 4 does not exist.")
	}

	err = db.Update(5, "new 5")

	if err == nil || err.Error() != "Index 5 does not exist." {
		t.Errorf(msg, err, "Index 5 does not exist.")
	}

	err = db.Update(0, "new 0")

	if err == nil || err.Error() != "Index 0 does not exist." {
		t.Errorf(msg, err, "Index 0 does not exist.")
	}
}

func TestUpdate(t *testing.T) {
	db := getDB3()

	err := db.Update(2, "new 2")

	if db.Length() != 3 {
		t.Errorf("DB::Update does not maintain length.")
	}

	if err != nil {
		t.Errorf("DB::Update returned error %v", err)
	}

	if db.lines[1] != "new 2" {
		t.Errorf("DB::Update does not set line text properly")
	}
}

func TestDeleteNonExistentLine(t *testing.T) {
	db := getDB3()
	msg := "DB::Delete: got %v, expected %v."

	err := db.Delete(4)

	if err == nil || err.Error() != "Index 4 does not exist." {
		t.Errorf(msg, err, "Index 4 does not exist.")
	}

	err = db.Delete(5)

	if err == nil || err.Error() != "Index 5 does not exist." {
		t.Errorf(msg, err, "Index 5 does not exist.")
	}

	err = db.Delete(0)

	if err == nil || err.Error() != "Index 0 does not exist." {
		t.Errorf(msg, err, "Index 0 does not exist.")
	}
}

func TestDelete(t *testing.T) {
	db := getDB3()

	err := db.Delete(2)

	if db.Length() != 2 {
		t.Errorf("DB::Delete does not decrease length by 1.")
	}

	if err != nil {
		t.Errorf("DB::Delete returned error %v", err)
	}

	if db.lines[0] != "test1" || db.lines[1] != "test3" {
		t.Errorf("DB::Delete does not maintain subsequent lines")
		t.Errorf("%v\n", db)
	}
}
