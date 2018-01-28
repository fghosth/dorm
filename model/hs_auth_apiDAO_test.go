package model

import (
	"reflect"
	"testing"

	"jvole.com/createProject/model/base"
)

func TestHsAuthApiDao_Select(t *testing.T) {
	type args struct {
		sql    string
		limit  int
		offset int
		value  []interface{}
	}
	tests := []struct {
		name    string
		dao     HsAuthApiDao
		args    args
		want    []interface{}
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := tt.dao.Select(tt.args.sql, tt.args.limit, tt.args.offset, tt.args.value...)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthApiDao.Select() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. HsAuthApiDao.Select() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestHsAuthApiDao_FindByID(t *testing.T) {
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		dao     *HsAuthApiDao
		args    args
		want    interface{}
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := tt.dao.FindByID(tt.args.id)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthApiDao.FindByID() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. HsAuthApiDao.FindByID() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestHsAuthApiDao_Add(t *testing.T) {
	tests := []struct {
		name    string
		dao     HsAuthApiDao
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := tt.dao.Add()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthApiDao.Add() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. HsAuthApiDao.Add() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestHsAuthApiDao_AddBatch(t *testing.T) {
	type args struct {
		obj []interface{}
	}
	tests := []struct {
		name    string
		dao     HsAuthApiDao
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := tt.dao.AddBatch(tt.args.obj); (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthApiDao.AddBatch() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestHsAuthApiDao_Update(t *testing.T) {
	tests := []struct {
		name    string
		dao     HsAuthApiDao
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := tt.dao.Update()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthApiDao.Update() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. HsAuthApiDao.Update() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestHsAuthApiDao_UpdateBatch(t *testing.T) {
	type args struct {
		obj []interface{}
	}
	tests := []struct {
		name    string
		dao     HsAuthApiDao
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := tt.dao.UpdateBatch(tt.args.obj); (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthApiDao.UpdateBatch() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestHsAuthApiDao_Delete(t *testing.T) {
	tests := []struct {
		name    string
		dao     HsAuthApiDao
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := tt.dao.Delete()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthApiDao.Delete() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. HsAuthApiDao.Delete() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestHsAuthApiDao_DeleteBatch(t *testing.T) {
	type args struct {
		obj []interface{}
	}
	tests := []struct {
		name    string
		dao     HsAuthApiDao
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := tt.dao.DeleteBatch(tt.args.obj); (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthApiDao.DeleteBatch() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestHsAuthApiDao_SDelete(t *testing.T) {
	tests := []struct {
		name    string
		dao     HsAuthApiDao
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := tt.dao.SDelete()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthApiDao.SDelete() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. HsAuthApiDao.SDelete() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestHsAuthApiDao_SDeleteBatch(t *testing.T) {
	type args struct {
		obj []interface{}
	}
	tests := []struct {
		name    string
		dao     HsAuthApiDao
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := tt.dao.SDeleteBatch(tt.args.obj); (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthApiDao.SDeleteBatch() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestHsAuthApiDao_Exec(t *testing.T) {
	type args struct {
		sql   string
		value []interface{}
	}
	tests := []struct {
		name    string
		dao     HsAuthApiDao
		args    args
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := tt.dao.Exec(tt.args.sql, tt.args.value...)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthApiDao.Exec() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. HsAuthApiDao.Exec() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestHsAuthApiDao_GetSql(t *testing.T) {
	tests := []struct {
		name  string
		dao   HsAuthApiDao
		want  string
		want1 []interface{}
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, got1 := tt.dao.GetSql()
		if got != tt.want {
			t.Errorf("%q. HsAuthApiDao.GetSql() got = %v, want %v", tt.name, got, tt.want)
		}
		if !reflect.DeepEqual(got1, tt.want1) {
			t.Errorf("%q. HsAuthApiDao.GetSql() got1 = %v, want %v", tt.name, got1, tt.want1)
		}
	}
}

func TestHsAuthApiDao_SetDBConn(t *testing.T) {
	type args struct {
		db  string
		str string
	}
	tests := []struct {
		name string
		dao  HsAuthApiDao
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt.dao.SetDBConn(tt.args.db, tt.args.str)
	}
}

func TestHsAuthApiDao_getObjWithValue(t *testing.T) {
	type args struct {
		dao HsAuthApiDao
	}
	tests := []struct {
		name string
		daoo HsAuthApiDao
		args args
		want base.HsAuthApi
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.daoo.getObjWithValue(tt.args.dao); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. HsAuthApiDao.getObjWithValue() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestNewHsAuthApiDao(t *testing.T) {
	tests := []struct {
		name string
		want HsAuthApiDao
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := NewHsAuthApiDao(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. NewHsAuthApiDao() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
