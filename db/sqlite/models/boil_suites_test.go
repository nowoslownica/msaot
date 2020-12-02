// Code generated by SQLBoiler 4.2.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Flexies", testFlexies)
	t.Run("GrammarPositions", testGrammarPositions)
}

func TestDelete(t *testing.T) {
	t.Run("Flexies", testFlexiesDelete)
	t.Run("GrammarPositions", testGrammarPositionsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Flexies", testFlexiesQueryDeleteAll)
	t.Run("GrammarPositions", testGrammarPositionsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Flexies", testFlexiesSliceDeleteAll)
	t.Run("GrammarPositions", testGrammarPositionsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Flexies", testFlexiesExists)
	t.Run("GrammarPositions", testGrammarPositionsExists)
}

func TestFind(t *testing.T) {
	t.Run("Flexies", testFlexiesFind)
	t.Run("GrammarPositions", testGrammarPositionsFind)
}

func TestBind(t *testing.T) {
	t.Run("Flexies", testFlexiesBind)
	t.Run("GrammarPositions", testGrammarPositionsBind)
}

func TestOne(t *testing.T) {
	t.Run("Flexies", testFlexiesOne)
	t.Run("GrammarPositions", testGrammarPositionsOne)
}

func TestAll(t *testing.T) {
	t.Run("Flexies", testFlexiesAll)
	t.Run("GrammarPositions", testGrammarPositionsAll)
}

func TestCount(t *testing.T) {
	t.Run("Flexies", testFlexiesCount)
	t.Run("GrammarPositions", testGrammarPositionsCount)
}

func TestInsert(t *testing.T) {
	t.Run("Flexies", testFlexiesInsert)
	t.Run("Flexies", testFlexiesInsertWhitelist)
	t.Run("GrammarPositions", testGrammarPositionsInsert)
	t.Run("GrammarPositions", testGrammarPositionsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("FlexyToGrammarPositionUsingGPosition", testFlexyToOneGrammarPositionUsingGPosition)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("GrammarPositionToGPositionFlexies", testGrammarPositionToManyGPositionFlexies)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("FlexyToGrammarPositionUsingGPositionFlexies", testFlexyToOneSetOpGrammarPositionUsingGPosition)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("FlexyToGrammarPositionUsingGPositionFlexies", testFlexyToOneRemoveOpGrammarPositionUsingGPosition)
}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("GrammarPositionToGPositionFlexies", testGrammarPositionToManyAddOpGPositionFlexies)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("GrammarPositionToGPositionFlexies", testGrammarPositionToManySetOpGPositionFlexies)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("GrammarPositionToGPositionFlexies", testGrammarPositionToManyRemoveOpGPositionFlexies)
}

func TestReload(t *testing.T) {
	t.Run("Flexies", testFlexiesReload)
	t.Run("GrammarPositions", testGrammarPositionsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Flexies", testFlexiesReloadAll)
	t.Run("GrammarPositions", testGrammarPositionsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Flexies", testFlexiesSelect)
	t.Run("GrammarPositions", testGrammarPositionsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Flexies", testFlexiesUpdate)
	t.Run("GrammarPositions", testGrammarPositionsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Flexies", testFlexiesSliceUpdateAll)
	t.Run("GrammarPositions", testGrammarPositionsSliceUpdateAll)
}