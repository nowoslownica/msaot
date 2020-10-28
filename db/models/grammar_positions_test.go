// Code generated by SQLBoiler 4.1.2 (https://gitlab.etecs.ru/forks/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/strmangle"
	"gitlab.etecs.ru/forks/sqlboiler/v4/boil"
	"gitlab.etecs.ru/forks/sqlboiler/v4/queries"
	"gitlab.etecs.ru/forks/sqlboiler/v4/queries/qm"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
	_ = qm.Where
)

func testGrammarPositions(t *testing.T) {
	t.Parallel()

	query := GrammarPositions()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testGrammarPositionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GrammarPosition{}
	if err = randomize.Struct(seed, o, grammarPositionDBTypes, true, grammarPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GrammarPosition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := GrammarPositions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGrammarPositionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GrammarPosition{}
	if err = randomize.Struct(seed, o, grammarPositionDBTypes, true, grammarPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GrammarPosition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := GrammarPositions().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := GrammarPositions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGrammarPositionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GrammarPosition{}
	if err = randomize.Struct(seed, o, grammarPositionDBTypes, true, grammarPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GrammarPosition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := GrammarPositionSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := GrammarPositions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGrammarPositionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GrammarPosition{}
	if err = randomize.Struct(seed, o, grammarPositionDBTypes, true, grammarPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GrammarPosition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := GrammarPositionExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if GrammarPosition exists: %s", err)
	}
	if !e {
		t.Errorf("Expected GrammarPositionExists to return true, but got false.")
	}
}

func testGrammarPositionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GrammarPosition{}
	if err = randomize.Struct(seed, o, grammarPositionDBTypes, true, grammarPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GrammarPosition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	grammarPositionFound, err := FindGrammarPosition(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if grammarPositionFound == nil {
		t.Error("want a record, got nil")
	}
}

func testGrammarPositionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GrammarPosition{}
	if err = randomize.Struct(seed, o, grammarPositionDBTypes, true, grammarPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GrammarPosition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = GrammarPositions().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testGrammarPositionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GrammarPosition{}
	if err = randomize.Struct(seed, o, grammarPositionDBTypes, true, grammarPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GrammarPosition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := GrammarPositions().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testGrammarPositionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	grammarPositionOne := &GrammarPosition{}
	grammarPositionTwo := &GrammarPosition{}
	if err = randomize.Struct(seed, grammarPositionOne, grammarPositionDBTypes, false, grammarPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GrammarPosition struct: %s", err)
	}
	if err = randomize.Struct(seed, grammarPositionTwo, grammarPositionDBTypes, false, grammarPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GrammarPosition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = grammarPositionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = grammarPositionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := GrammarPositions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testGrammarPositionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	grammarPositionOne := &GrammarPosition{}
	grammarPositionTwo := &GrammarPosition{}
	if err = randomize.Struct(seed, grammarPositionOne, grammarPositionDBTypes, false, grammarPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GrammarPosition struct: %s", err)
	}
	if err = randomize.Struct(seed, grammarPositionTwo, grammarPositionDBTypes, false, grammarPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GrammarPosition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = grammarPositionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = grammarPositionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := GrammarPositions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testGrammarPositionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GrammarPosition{}
	if err = randomize.Struct(seed, o, grammarPositionDBTypes, true, grammarPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GrammarPosition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := GrammarPositions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGrammarPositionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GrammarPosition{}
	if err = randomize.Struct(seed, o, grammarPositionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize GrammarPosition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(grammarPositionColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := GrammarPositions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGrammarPositionToManyGPositionFlexies(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a GrammarPosition
	var b, c Flexy

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, grammarPositionDBTypes, true, grammarPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GrammarPosition struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, flexyDBTypes, false, flexyColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, flexyDBTypes, false, flexyColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.GPosition, a.ID)
	queries.Assign(&c.GPosition, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.GPositionFlexies().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.GPosition, b.GPosition) {
			bFound = true
		}
		if queries.Equal(v.GPosition, c.GPosition) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := GrammarPositionSlice{&a}
	if err = a.L.LoadGPositionFlexies(ctx, tx, false, (*[]*GrammarPosition)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.GPositionFlexies); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.GPositionFlexies = nil
	if err = a.L.LoadGPositionFlexies(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.GPositionFlexies); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testGrammarPositionToManyAddOpGPositionFlexies(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a GrammarPosition
	var b, c, d, e Flexy

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, grammarPositionDBTypes, false, strmangle.SetComplement(grammarPositionPrimaryKeyColumns, grammarPositionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Flexy{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, flexyDBTypes, false, strmangle.SetComplement(flexyPrimaryKeyColumns, flexyColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Flexy{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddGPositionFlexies(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.GPosition) {
			t.Error("foreign key was wrong value", a.ID, first.GPosition)
		}
		if !queries.Equal(a.ID, second.GPosition) {
			t.Error("foreign key was wrong value", a.ID, second.GPosition)
		}

		if first.R.GPosition != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.GPosition != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.GPositionFlexies[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.GPositionFlexies[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.GPositionFlexies().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testGrammarPositionToManySetOpGPositionFlexies(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a GrammarPosition
	var b, c, d, e Flexy

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, grammarPositionDBTypes, false, strmangle.SetComplement(grammarPositionPrimaryKeyColumns, grammarPositionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Flexy{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, flexyDBTypes, false, strmangle.SetComplement(flexyPrimaryKeyColumns, flexyColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetGPositionFlexies(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.GPositionFlexies().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetGPositionFlexies(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.GPositionFlexies().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.GPosition) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.GPosition) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.GPosition) {
		t.Error("foreign key was wrong value", a.ID, d.GPosition)
	}
	if !queries.Equal(a.ID, e.GPosition) {
		t.Error("foreign key was wrong value", a.ID, e.GPosition)
	}

	if b.R.GPosition != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.GPosition != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.GPosition != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.GPosition != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.GPositionFlexies[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.GPositionFlexies[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testGrammarPositionToManyRemoveOpGPositionFlexies(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a GrammarPosition
	var b, c, d, e Flexy

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, grammarPositionDBTypes, false, strmangle.SetComplement(grammarPositionPrimaryKeyColumns, grammarPositionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Flexy{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, flexyDBTypes, false, strmangle.SetComplement(flexyPrimaryKeyColumns, flexyColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddGPositionFlexies(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.GPositionFlexies().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveGPositionFlexies(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.GPositionFlexies().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.GPosition) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.GPosition) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.GPosition != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.GPosition != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.GPosition != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.GPosition != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.GPositionFlexies) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.GPositionFlexies[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.GPositionFlexies[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testGrammarPositionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GrammarPosition{}
	if err = randomize.Struct(seed, o, grammarPositionDBTypes, true, grammarPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GrammarPosition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testGrammarPositionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GrammarPosition{}
	if err = randomize.Struct(seed, o, grammarPositionDBTypes, true, grammarPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GrammarPosition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := GrammarPositionSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testGrammarPositionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GrammarPosition{}
	if err = randomize.Struct(seed, o, grammarPositionDBTypes, true, grammarPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GrammarPosition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := GrammarPositions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	grammarPositionDBTypes = map[string]string{`ID`: `INT`, `GCase`: `INT`, `GPerson`: `INT`, `GNumber`: `INT`, `GTense`: `INT`, `GGender`: `INT`, `Declension`: `INT`, `Conjugation`: `INT`}
	_                      = bytes.MinRead
)

func testGrammarPositionsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(grammarPositionPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(grammarPositionAllColumns) == len(grammarPositionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &GrammarPosition{}
	if err = randomize.Struct(seed, o, grammarPositionDBTypes, true, grammarPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GrammarPosition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := GrammarPositions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, grammarPositionDBTypes, true, grammarPositionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize GrammarPosition struct: %s", err)
	}

	whitelist := strmangle.SetComplement(grammarPositionAllColumns, grammarPositionPrimaryKeyColumns)
	if rowsAff, err := o.Update(ctx, tx, boil.Whitelist(whitelist...)); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testGrammarPositionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(grammarPositionAllColumns) == len(grammarPositionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &GrammarPosition{}
	if err = randomize.Struct(seed, o, grammarPositionDBTypes, true, grammarPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GrammarPosition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := GrammarPositions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, grammarPositionDBTypes, true, grammarPositionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize GrammarPosition struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(grammarPositionAllColumns, grammarPositionPrimaryKeyColumns) {
		fields = grammarPositionAllColumns
	} else {
		fields = strmangle.SetComplement(
			grammarPositionAllColumns,
			grammarPositionPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := GrammarPositionSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}
