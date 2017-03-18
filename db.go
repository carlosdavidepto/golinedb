package golinedb

import (
	"fmt"
	"github.com/carlosdavidepto/golinedb/iface"
)

type DB struct {
	lines []string
}

func (db *DB) Search(str string) (error, []iface.Line) {
	return nil, []iface.Line{}
}

func (db *DB) RegExp(str string) (error, []iface.Line) {
	return nil, []iface.Line{}
}

func (db *DB) Insert(idx int, text string) error {
	return nil
}

func (db *DB) Update(idx int, text string) error {

	realIdx := idx - 1

	if realIdx < 0 || realIdx >= db.Length() {
		return fmt.Errorf("Index %d does not exist.", idx)
	}

	db.lines[realIdx] = text

	return nil
}

func (db *DB) Delete(idx int) error {

	realIdx := idx - 1

	if realIdx < 0 || realIdx >= db.Length() {
		return fmt.Errorf("Index %d does not exist.", idx)
	}

	db.lines = append(db.lines[:realIdx], db.lines[(realIdx+1):]...)

	return nil
}

func (db *DB) Length() int {
	return len(db.lines)
}

func (db *DB) Unwrap(str string) {

}

func (db *DB) Wrap() string {
	return ""
}
