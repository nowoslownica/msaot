// Code generated by SQLBoiler 4.3.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testLemmas(t *testing.T) {
	t.Parallel()

	query := Lemmas()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testLemmasDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Lemma{}
	if err = randomize.Struct(seed, o, lemmaDBTypes, true, lemmaColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Lemma struct: %s", err)
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

	count, err := Lemmas().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLemmasQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Lemma{}
	if err = randomize.Struct(seed, o, lemmaDBTypes, true, lemmaColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Lemma struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Lemmas().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Lemmas().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLemmasSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Lemma{}
	if err = randomize.Struct(seed, o, lemmaDBTypes, true, lemmaColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Lemma struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := LemmaSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Lemmas().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLemmasExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Lemma{}
	if err = randomize.Struct(seed, o, lemmaDBTypes, true, lemmaColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Lemma struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := LemmaExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Lemma exists: %s", err)
	}
	if !e {
		t.Errorf("Expected LemmaExists to return true, but got false.")
	}
}

func testLemmasFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Lemma{}
	if err = randomize.Struct(seed, o, lemmaDBTypes, true, lemmaColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Lemma struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	lemmaFound, err := FindLemma(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if lemmaFound == nil {
		t.Error("want a record, got nil")
	}
}

func testLemmasBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Lemma{}
	if err = randomize.Struct(seed, o, lemmaDBTypes, true, lemmaColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Lemma struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Lemmas().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testLemmasOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Lemma{}
	if err = randomize.Struct(seed, o, lemmaDBTypes, true, lemmaColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Lemma struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Lemmas().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testLemmasAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	lemmaOne := &Lemma{}
	lemmaTwo := &Lemma{}
	if err = randomize.Struct(seed, lemmaOne, lemmaDBTypes, false, lemmaColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Lemma struct: %s", err)
	}
	if err = randomize.Struct(seed, lemmaTwo, lemmaDBTypes, false, lemmaColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Lemma struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = lemmaOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = lemmaTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Lemmas().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testLemmasCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	lemmaOne := &Lemma{}
	lemmaTwo := &Lemma{}
	if err = randomize.Struct(seed, lemmaOne, lemmaDBTypes, false, lemmaColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Lemma struct: %s", err)
	}
	if err = randomize.Struct(seed, lemmaTwo, lemmaDBTypes, false, lemmaColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Lemma struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = lemmaOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = lemmaTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Lemmas().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testLemmasInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Lemma{}
	if err = randomize.Struct(seed, o, lemmaDBTypes, true, lemmaColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Lemma struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Lemmas().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testLemmasInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Lemma{}
	if err = randomize.Struct(seed, o, lemmaDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Lemma struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(lemmaColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Lemmas().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testLemmaToManyLemmaIdFlexies(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Lemma
	var b, c Flexy

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, lemmaDBTypes, true, lemmaColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Lemma struct: %s", err)
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

	queries.Assign(&b.LemmaId, a.ID)
	queries.Assign(&c.LemmaId, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.LemmaIdFlexies().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.LemmaId, b.LemmaId) {
			bFound = true
		}
		if queries.Equal(v.LemmaId, c.LemmaId) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := LemmaSlice{&a}
	if err = a.L.LoadLemmaIdFlexies(ctx, tx, false, (*[]*Lemma)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.LemmaIdFlexies); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.LemmaIdFlexies = nil
	if err = a.L.LoadLemmaIdFlexies(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.LemmaIdFlexies); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testLemmaToManyAddOpLemmaIdFlexies(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Lemma
	var b, c, d, e Flexy

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, lemmaDBTypes, false, strmangle.SetComplement(lemmaPrimaryKeyColumns, lemmaColumnsWithoutDefault)...); err != nil {
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
		err = a.AddLemmaIdFlexies(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.LemmaId) {
			t.Error("foreign key was wrong value", a.ID, first.LemmaId)
		}
		if !queries.Equal(a.ID, second.LemmaId) {
			t.Error("foreign key was wrong value", a.ID, second.LemmaId)
		}

		if first.R.LemmaIdLemmas != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.LemmaIdLemmas != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.LemmaIdFlexies[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.LemmaIdFlexies[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.LemmaIdFlexies().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testLemmaToManySetOpLemmaIdFlexies(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Lemma
	var b, c, d, e Flexy

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, lemmaDBTypes, false, strmangle.SetComplement(lemmaPrimaryKeyColumns, lemmaColumnsWithoutDefault)...); err != nil {
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

	err = a.SetLemmaIdFlexies(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.LemmaIdFlexies().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetLemmaIdFlexies(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.LemmaIdFlexies().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.LemmaId) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.LemmaId) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.LemmaId) {
		t.Error("foreign key was wrong value", a.ID, d.LemmaId)
	}
	if !queries.Equal(a.ID, e.LemmaId) {
		t.Error("foreign key was wrong value", a.ID, e.LemmaId)
	}

	if b.R.LemmaIdLemmas != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.LemmaIdLemmas != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.LemmaIdLemmas != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.LemmaIdLemmas != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.LemmaIdFlexies[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.LemmaIdFlexies[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testLemmaToManyRemoveOpLemmaIdFlexies(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Lemma
	var b, c, d, e Flexy

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, lemmaDBTypes, false, strmangle.SetComplement(lemmaPrimaryKeyColumns, lemmaColumnsWithoutDefault)...); err != nil {
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

	err = a.AddLemmaIdFlexies(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.LemmaIdFlexies().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveLemmaIdFlexies(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.LemmaIdFlexies().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.LemmaId) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.LemmaId) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.LemmaIdLemmas != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.LemmaIdLemmas != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.LemmaIdLemmas != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.LemmaIdLemmas != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.LemmaIdFlexies) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.LemmaIdFlexies[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.LemmaIdFlexies[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testLemmasReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Lemma{}
	if err = randomize.Struct(seed, o, lemmaDBTypes, true, lemmaColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Lemma struct: %s", err)
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

func testLemmasReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Lemma{}
	if err = randomize.Struct(seed, o, lemmaDBTypes, true, lemmaColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Lemma struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := LemmaSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testLemmasSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Lemma{}
	if err = randomize.Struct(seed, o, lemmaDBTypes, true, lemmaColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Lemma struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Lemmas().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	lemmaDBTypes = map[string]string{`ID`: `INT`, `Value`: `TEXT`, `Pos`: `TEXT`, `ChangeSchema`: `INT`}
	_            = bytes.MinRead
)

func testLemmasUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(lemmaPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(lemmaAllColumns) == len(lemmaPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Lemma{}
	if err = randomize.Struct(seed, o, lemmaDBTypes, true, lemmaColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Lemma struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Lemmas().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, lemmaDBTypes, true, lemmaPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Lemma struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testLemmasSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(lemmaAllColumns) == len(lemmaPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Lemma{}
	if err = randomize.Struct(seed, o, lemmaDBTypes, true, lemmaColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Lemma struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Lemmas().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, lemmaDBTypes, true, lemmaPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Lemma struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(lemmaAllColumns, lemmaPrimaryKeyColumns) {
		fields = lemmaAllColumns
	} else {
		fields = strmangle.SetComplement(
			lemmaAllColumns,
			lemmaPrimaryKeyColumns,
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

	slice := LemmaSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}