// Code generated by SQLBoiler 4.16.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testPets(t *testing.T) {
	t.Parallel()

	query := Pets()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testPetsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Pet{}
	if err = randomize.Struct(seed, o, petDBTypes, true, petColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pet struct: %s", err)
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

	count, err := Pets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPetsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Pet{}
	if err = randomize.Struct(seed, o, petDBTypes, true, petColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Pets().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Pets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPetsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Pet{}
	if err = randomize.Struct(seed, o, petDBTypes, true, petColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PetSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Pets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPetsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Pet{}
	if err = randomize.Struct(seed, o, petDBTypes, true, petColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := PetExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Pet exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PetExists to return true, but got false.")
	}
}

func testPetsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Pet{}
	if err = randomize.Struct(seed, o, petDBTypes, true, petColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	petFound, err := FindPet(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if petFound == nil {
		t.Error("want a record, got nil")
	}
}

func testPetsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Pet{}
	if err = randomize.Struct(seed, o, petDBTypes, true, petColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Pets().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testPetsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Pet{}
	if err = randomize.Struct(seed, o, petDBTypes, true, petColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Pets().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPetsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	petOne := &Pet{}
	petTwo := &Pet{}
	if err = randomize.Struct(seed, petOne, petDBTypes, false, petColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pet struct: %s", err)
	}
	if err = randomize.Struct(seed, petTwo, petDBTypes, false, petColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = petOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = petTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Pets().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPetsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	petOne := &Pet{}
	petTwo := &Pet{}
	if err = randomize.Struct(seed, petOne, petDBTypes, false, petColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pet struct: %s", err)
	}
	if err = randomize.Struct(seed, petTwo, petDBTypes, false, petColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = petOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = petTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Pets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func petBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Pet) error {
	*o = Pet{}
	return nil
}

func petAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Pet) error {
	*o = Pet{}
	return nil
}

func petAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Pet) error {
	*o = Pet{}
	return nil
}

func petBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Pet) error {
	*o = Pet{}
	return nil
}

func petAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Pet) error {
	*o = Pet{}
	return nil
}

func petBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Pet) error {
	*o = Pet{}
	return nil
}

func petAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Pet) error {
	*o = Pet{}
	return nil
}

func petBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Pet) error {
	*o = Pet{}
	return nil
}

func petAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Pet) error {
	*o = Pet{}
	return nil
}

func testPetsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Pet{}
	o := &Pet{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, petDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Pet object: %s", err)
	}

	AddPetHook(boil.BeforeInsertHook, petBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	petBeforeInsertHooks = []PetHook{}

	AddPetHook(boil.AfterInsertHook, petAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	petAfterInsertHooks = []PetHook{}

	AddPetHook(boil.AfterSelectHook, petAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	petAfterSelectHooks = []PetHook{}

	AddPetHook(boil.BeforeUpdateHook, petBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	petBeforeUpdateHooks = []PetHook{}

	AddPetHook(boil.AfterUpdateHook, petAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	petAfterUpdateHooks = []PetHook{}

	AddPetHook(boil.BeforeDeleteHook, petBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	petBeforeDeleteHooks = []PetHook{}

	AddPetHook(boil.AfterDeleteHook, petAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	petAfterDeleteHooks = []PetHook{}

	AddPetHook(boil.BeforeUpsertHook, petBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	petBeforeUpsertHooks = []PetHook{}

	AddPetHook(boil.AfterUpsertHook, petAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	petAfterUpsertHooks = []PetHook{}
}

func testPetsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Pet{}
	if err = randomize.Struct(seed, o, petDBTypes, true, petColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Pets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPetsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Pet{}
	if err = randomize.Struct(seed, o, petDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Pet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(petColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Pets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPetToOneCategoryUsingPetCategory(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Pet
	var foreign Category

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, petDBTypes, true, petColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pet struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, categoryDBTypes, false, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.Category, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.PetCategory().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddCategoryHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *Category) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := PetSlice{&local}
	if err = local.L.LoadPetCategory(ctx, tx, false, (*[]*Pet)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.PetCategory == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.PetCategory = nil
	if err = local.L.LoadPetCategory(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.PetCategory == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testPetToOneSetOpCategoryUsingPetCategory(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Pet
	var b, c Category

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, petDBTypes, false, strmangle.SetComplement(petPrimaryKeyColumns, petColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, categoryDBTypes, false, strmangle.SetComplement(categoryPrimaryKeyColumns, categoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, categoryDBTypes, false, strmangle.SetComplement(categoryPrimaryKeyColumns, categoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Category{&b, &c} {
		err = a.SetPetCategory(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.PetCategory != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Pets[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.Category, x.ID) {
			t.Error("foreign key was wrong value", a.Category)
		}

		zero := reflect.Zero(reflect.TypeOf(a.Category))
		reflect.Indirect(reflect.ValueOf(&a.Category)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.Category, x.ID) {
			t.Error("foreign key was wrong value", a.Category, x.ID)
		}
	}
}

func testPetToOneRemoveOpCategoryUsingPetCategory(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Pet
	var b Category

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, petDBTypes, false, strmangle.SetComplement(petPrimaryKeyColumns, petColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, categoryDBTypes, false, strmangle.SetComplement(categoryPrimaryKeyColumns, categoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetPetCategory(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemovePetCategory(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.PetCategory().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.PetCategory != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.Category) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.Pets) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testPetsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Pet{}
	if err = randomize.Struct(seed, o, petDBTypes, true, petColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pet struct: %s", err)
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

func testPetsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Pet{}
	if err = randomize.Struct(seed, o, petDBTypes, true, petColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PetSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPetsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Pet{}
	if err = randomize.Struct(seed, o, petDBTypes, true, petColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Pets().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	petDBTypes = map[string]string{`ID`: `integer`, `Name`: `text`, `Category`: `integer`}
	_          = bytes.MinRead
)

func testPetsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(petPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(petAllColumns) == len(petPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Pet{}
	if err = randomize.Struct(seed, o, petDBTypes, true, petColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Pets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, petDBTypes, true, petPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Pet struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testPetsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(petAllColumns) == len(petPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Pet{}
	if err = randomize.Struct(seed, o, petDBTypes, true, petColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Pets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, petDBTypes, true, petPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Pet struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(petAllColumns, petPrimaryKeyColumns) {
		fields = petAllColumns
	} else {
		fields = strmangle.SetComplement(
			petAllColumns,
			petPrimaryKeyColumns,
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

	slice := PetSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testPetsUpsert(t *testing.T) {
	t.Parallel()

	if len(petAllColumns) == len(petPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Pet{}
	if err = randomize.Struct(seed, &o, petDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Pet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Pet: %s", err)
	}

	count, err := Pets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, petDBTypes, false, petPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Pet struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Pet: %s", err)
	}

	count, err = Pets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
