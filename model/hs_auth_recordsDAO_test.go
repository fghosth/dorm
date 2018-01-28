package model

import (
	"reflect"
	"testing"

	"jvole.com/createProject/model/base"
)

func TestHsAuthRecordsDao_Select(t *testing.T) {
	type args struct {
		sql    string
		limit  int
		offset int
		value  []interface{}
	}
	tests := []struct {
		name    string
		dao     HsAuthRecordsDao
		args    args
		want    []interface{}
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := tt.dao.Select(tt.args.sql, tt.args.limit, tt.args.offset, tt.args.value...)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthRecordsDao.Select() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. HsAuthRecordsDao.Select() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestHsAuthRecordsDao_FindByID(t *testing.T) {
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		dao     *HsAuthRecordsDao
		args    args
		want    interface{}
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := tt.dao.FindByID(tt.args.id)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthRecordsDao.FindByID() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. HsAuthRecordsDao.FindByID() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestHsAuthRecordsDao_Add(t *testing.T) {
	tests := []struct {
		name    string
		dao     HsAuthRecordsDao
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := tt.dao.Add()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthRecordsDao.Add() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. HsAuthRecordsDao.Add() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestHsAuthRecordsDao_AddBatch(t *testing.T) {
	type args struct {
		obj []interface{}
	}
	tests := []struct {
		name    string
		dao     HsAuthRecordsDao
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := tt.dao.AddBatch(tt.args.obj); (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthRecordsDao.AddBatch() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestHsAuthRecordsDao_Update(t *testing.T) {
	tests := []struct {
		name    string
		dao     HsAuthRecordsDao
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := tt.dao.Update()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthRecordsDao.Update() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. HsAuthRecordsDao.Update() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestHsAuthRecordsDao_UpdateBatch(t *testing.T) {
	type args struct {
		obj []interface{}
	}
	tests := []struct {
		name    string
		dao     HsAuthRecordsDao
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := tt.dao.UpdateBatch(tt.args.obj); (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthRecordsDao.UpdateBatch() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestHsAuthRecordsDao_Delete(t *testing.T) {
	tests := []struct {
		name    string
		dao     HsAuthRecordsDao
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := tt.dao.Delete()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthRecordsDao.Delete() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. HsAuthRecordsDao.Delete() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestHsAuthRecordsDao_DeleteBatch(t *testing.T) {
	type args struct {
		obj []interface{}
	}
	tests := []struct {
		name    string
		dao     HsAuthRecordsDao
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := tt.dao.DeleteBatch(tt.args.obj); (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthRecordsDao.DeleteBatch() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestHsAuthRecordsDao_SDelete(t *testing.T) {
	tests := []struct {
		name    string
		dao     HsAuthRecordsDao
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := tt.dao.SDelete()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthRecordsDao.SDelete() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. HsAuthRecordsDao.SDelete() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestHsAuthRecordsDao_SDeleteBatch(t *testing.T) {
	type args struct {
		obj []interface{}
	}
	tests := []struct {
		name    string
		dao     HsAuthRecordsDao
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := tt.dao.SDeleteBatch(tt.args.obj); (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthRecordsDao.SDeleteBatch() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestHsAuthRecordsDao_Exec(t *testing.T) {
	type args struct {
		sql   string
		value []interface{}
	}
	tests := []struct {
		name    string
		dao     HsAuthRecordsDao
		args    args
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := tt.dao.Exec(tt.args.sql, tt.args.value...)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthRecordsDao.Exec() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. HsAuthRecordsDao.Exec() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestHsAuthRecordsDao_GetSql(t *testing.T) {
	tests := []struct {
		name  string
		dao   HsAuthRecordsDao
		want  string
		want1 []interface{}
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, got1 := tt.dao.GetSql()
		if got != tt.want {
			t.Errorf("%q. HsAuthRecordsDao.GetSql() got = %v, want %v", tt.name, got, tt.want)
		}
		if !reflect.DeepEqual(got1, tt.want1) {
			t.Errorf("%q. HsAuthRecordsDao.GetSql() got1 = %v, want %v", tt.name, got1, tt.want1)
		}
	}
}

func TestHsAuthRecordsDao_SetDBConn(t *testing.T) {
	type args struct {
		db  string
		str string
	}
	tests := []struct {
		name string
		dao  HsAuthRecordsDao
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt.dao.SetDBConn(tt.args.db, tt.args.str)
	}
}

func TestHsAuthRecordsDao_getObjWithValue(t *testing.T) {
	type args struct {
		dao HsAuthRecordsDao
	}
	tests := []struct {
		name string
		daoo HsAuthRecordsDao
		args args
		want base.HsAuthRecords
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.daoo.getObjWithValue(tt.args.dao); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. HsAuthRecordsDao.getObjWithValue() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestNewHsAuthRecordsDao(t *testing.T) {
	tests := []struct {
		name string
		want HsAuthRecordsDao
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := NewHsAuthRecordsDao(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. NewHsAuthRecordsDao() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
