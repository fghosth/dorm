package model

import (
	"reflect"
	"testing"

	"jvole.com/createProject/model/base"
)

func TestHsAuthApplicationDao_Select(t *testing.T) {
	type args struct {
		sql    string
		limit  int
		offset int
		value  []interface{}
	}
	tests := []struct {
		name    string
		dao     HsAuthApplicationDao
		args    args
		want    []interface{}
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := tt.dao.Select(tt.args.sql, tt.args.limit, tt.args.offset, tt.args.value...)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthApplicationDao.Select() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. HsAuthApplicationDao.Select() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestHsAuthApplicationDao_FindByID(t *testing.T) {
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		dao     *HsAuthApplicationDao
		args    args
		want    interface{}
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := tt.dao.FindByID(tt.args.id)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthApplicationDao.FindByID() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. HsAuthApplicationDao.FindByID() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestHsAuthApplicationDao_Add(t *testing.T) {
	tests := []struct {
		name    string
		dao     HsAuthApplicationDao
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := tt.dao.Add()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthApplicationDao.Add() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. HsAuthApplicationDao.Add() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestHsAuthApplicationDao_AddBatch(t *testing.T) {
	type args struct {
		obj []interface{}
	}
	tests := []struct {
		name    string
		dao     HsAuthApplicationDao
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := tt.dao.AddBatch(tt.args.obj); (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthApplicationDao.AddBatch() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestHsAuthApplicationDao_Update(t *testing.T) {
	tests := []struct {
		name    string
		dao     HsAuthApplicationDao
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := tt.dao.Update()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthApplicationDao.Update() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. HsAuthApplicationDao.Update() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestHsAuthApplicationDao_UpdateBatch(t *testing.T) {
	type args struct {
		obj []interface{}
	}
	tests := []struct {
		name    string
		dao     HsAuthApplicationDao
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := tt.dao.UpdateBatch(tt.args.obj); (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthApplicationDao.UpdateBatch() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestHsAuthApplicationDao_Delete(t *testing.T) {
	tests := []struct {
		name    string
		dao     HsAuthApplicationDao
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := tt.dao.Delete()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthApplicationDao.Delete() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. HsAuthApplicationDao.Delete() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestHsAuthApplicationDao_DeleteBatch(t *testing.T) {
	type args struct {
		obj []interface{}
	}
	tests := []struct {
		name    string
		dao     HsAuthApplicationDao
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := tt.dao.DeleteBatch(tt.args.obj); (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthApplicationDao.DeleteBatch() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestHsAuthApplicationDao_SDelete(t *testing.T) {
	tests := []struct {
		name    string
		dao     HsAuthApplicationDao
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := tt.dao.SDelete()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthApplicationDao.SDelete() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. HsAuthApplicationDao.SDelete() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestHsAuthApplicationDao_SDeleteBatch(t *testing.T) {
	type args struct {
		obj []interface{}
	}
	tests := []struct {
		name    string
		dao     HsAuthApplicationDao
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := tt.dao.SDeleteBatch(tt.args.obj); (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthApplicationDao.SDeleteBatch() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestHsAuthApplicationDao_Exec(t *testing.T) {
	type args struct {
		sql   string
		value []interface{}
	}
	tests := []struct {
		name    string
		dao     HsAuthApplicationDao
		args    args
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := tt.dao.Exec(tt.args.sql, tt.args.value...)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthApplicationDao.Exec() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. HsAuthApplicationDao.Exec() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestHsAuthApplicationDao_GetSql(t *testing.T) {
	tests := []struct {
		name  string
		dao   HsAuthApplicationDao
		want  string
		want1 []interface{}
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, got1 := tt.dao.GetSql()
		if got != tt.want {
			t.Errorf("%q. HsAuthApplicationDao.GetSql() got = %v, want %v", tt.name, got, tt.want)
		}
		if !reflect.DeepEqual(got1, tt.want1) {
			t.Errorf("%q. HsAuthApplicationDao.GetSql() got1 = %v, want %v", tt.name, got1, tt.want1)
		}
	}
}

func TestHsAuthApplicationDao_SetDBConn(t *testing.T) {
	type args struct {
		db  string
		str string
	}
	tests := []struct {
		name string
		dao  HsAuthApplicationDao
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt.dao.SetDBConn(tt.args.db, tt.args.str)
	}
}

func TestHsAuthApplicationDao_getObjWithValue(t *testing.T) {
	type args struct {
		dao HsAuthApplicationDao
	}
	tests := []struct {
		name string
		daoo HsAuthApplicationDao
		args args
		want base.HsAuthApplication
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.daoo.getObjWithValue(tt.args.dao); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. HsAuthApplicationDao.getObjWithValue() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestNewHsAuthApplicationDao(t *testing.T) {
	tests := []struct {
		name string
		want HsAuthApplicationDao
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := NewHsAuthApplicationDao(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. NewHsAuthApplicationDao() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
