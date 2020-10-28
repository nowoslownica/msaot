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

func testFlexies(t *testing.T) {
	t.Parallel()

	query := Flexies()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testFlexiesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Flexy{}
	if err = randomize.Struct(seed, o, flexyDBTypes, true, flexyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Flexy struct: %s", err)
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

	count, err := Flexies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFlexiesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Flexy{}
	if err = randomize.Struct(seed, o, flexyDBTypes, true, flexyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Flexy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Flexies().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Flexies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFlexiesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Flexy{}
	if err = randomize.Struct(seed, o, flexyDBTypes, true, flexyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Flexy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := FlexySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Flexies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFlexiesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Flexy{}
	if err = randomize.Struct(seed, o, flexyDBTypes, true, flexyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Flexy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := FlexyExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Flexy exists: %s", err)
	}
	if !e {
		t.Errorf("Expected FlexyExists to return true, but got false.")
	}
}

func testFlexiesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Flexy{}
	if err = randomize.Struct(seed, o, flexyDBTypes, true, flexyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Flexy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	flexyFound, err := FindFlexy(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if flexyFound == nil {
		t.Error("want a record, got nil")
	}
}

func testFlexiesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Flexy{}
	if err = randomize.Struct(seed, o, flexyDBTypes, true, flexyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Flexy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Flexies().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testFlexiesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Flexy{}
	if err = randomize.Struct(seed, o, flexyDBTypes, true, flexyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Flexy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Flexies().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testFlexiesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	flexyOne := &Flexy{}
	flexyTwo := &Flexy{}
	if err = randomize.Struct(seed, flexyOne, flexyDBTypes, false, flexyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Flexy struct: %s", err)
	}
	if err = randomize.Struct(seed, flexyTwo, flexyDBTypes, false, flexyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Flexy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = flexyOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = flexyTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Flexies().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testFlexiesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	flexyOne := &Flexy{}
	flexyTwo := &Flexy{}
	if err = randomize.Struct(seed, flexyOne, flexyDBTypes, false, flexyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Flexy struct: %s", err)
	}
	if err = randomize.Struct(seed, flexyTwo, flexyDBTypes, false, flexyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Flexy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = flexyOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = flexyTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Flexies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testFlexiesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Flexy{}
	if err = randomize.Struct(seed, o, flexyDBTypes, true, flexyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Flexy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Flexies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFlexiesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Flexy{}
	if err = randomize.Struct(seed, o, flexyDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Flexy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(flexyColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Flexies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFlexyToOneGrammarPositionUsingGPosition(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Flexy
	var foreign GrammarPosition

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, flexyDBTypes, true, flexyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Flexy struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, grammarPositionDBTypes, true, grammarPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GrammarPosition struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.GPosition, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.GPosition().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := FlexySlice{&local}
	if err = local.L.LoadGPosition(ctx, tx, false, (*[]*Flexy)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.GPosition == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.GPosition = nil
	if err = local.L.LoadGPosition(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.GPosition == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFlexyToOneSetOpGrammarPositionUsingGPosition(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Flexy
	var b, c GrammarPosition

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, flexyDBTypes, false, strmangle.SetComplement(flexyPrimaryKeyColumns, flexyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, grammarPositionDBTypes, false, strmangle.SetComplement(grammarPositionPrimaryKeyColumns, grammarPositionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, grammarPositionDBTypes, false, strmangle.SetComplement(grammarPositionPrimaryKeyColumns, grammarPositionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*GrammarPosition{&b, &c} {
		err = a.SetGPosition(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.GPosition != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.GPositionFlexies[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.GPosition, x.ID) {
			t.Error("foreign key was wrong value", a.GPosition)
		}

		zero := reflect.Zero(reflect.TypeOf(a.GPosition))
		reflect.Indirect(reflect.ValueOf(&a.GPosition)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.GPosition, x.ID) {
			t.Error("foreign key was wrong value", a.GPosition, x.ID)
		}
	}
}

func testFlexyToOneRemoveOpGrammarPositionUsingGPosition(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Flexy
	var b GrammarPosition

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, flexyDBTypes, false, strmangle.SetComplement(flexyPrimaryKeyColumns, flexyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, grammarPositionDBTypes, false, strmangle.SetComplement(grammarPositionPrimaryKeyColumns, grammarPositionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetGPosition(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveGPosition(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.GPosition().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.GPosition != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.GPosition) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.GPositionFlexies) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testFlexiesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Flexy{}
	if err = randomize.Struct(seed, o, flexyDBTypes, true, flexyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Flexy struct: %s", err)
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

func testFlexiesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Flexy{}
	if err = randomize.Struct(seed, o, flexyDBTypes, true, flexyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Flexy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := FlexySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testFlexiesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Flexy{}
	if err = randomize.Struct(seed, o, flexyDBTypes, true, flexyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Flexy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Flexies().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	flexyDBTypes = map[string]string{`ID`: `INT`, `Value`: `TEXT`, `Normal`: `TEXT`, `Pos`: `TEXT`, `GPosition`: `INT`}
	_            = bytes.MinRead
)

func testFlexiesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(flexyPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(flexyAllColumns) == len(flexyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Flexy{}
	if err = randomize.Struct(seed, o, flexyDBTypes, true, flexyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Flexy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Flexies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, flexyDBTypes, true, flexyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Flexy struct: %s", err)
	}

	whitelist := strmangle.SetComplement(flexyAllColumns, flexyPrimaryKeyColumns)
	if rowsAff, err := o.Update(ctx, tx, boil.Whitelist(whitelist...)); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testFlexiesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(flexyAllColumns) == len(flexyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Flexy{}
	if err = randomize.Struct(seed, o, flexyDBTypes, true, flexyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Flexy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Flexies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, flexyDBTypes, true, flexyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Flexy struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(flexyAllColumns, flexyPrimaryKeyColumns) {
		fields = flexyAllColumns
	} else {
		fields = strmangle.SetComplement(
			flexyAllColumns,
			flexyPrimaryKeyColumns,
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

	slice := FlexySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}
