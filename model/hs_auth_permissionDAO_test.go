package model

import (
	"reflect"
	"testing"

	"jvole.com/createProject/model/base"
)

func TestHsAuthPermissionDao_Select(t *testing.T) {
	type args struct {
		sql    string
		limit  int
		offset int
		value  []interface{}
	}
	tests := []struct {
		name    string
		dao     HsAuthPermissionDao
		args    args
		want    []interface{}
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := tt.dao.Select(tt.args.sql, tt.args.limit, tt.args.offset, tt.args.value...)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthPermissionDao.Select() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. HsAuthPermissionDao.Select() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestHsAuthPermissionDao_FindByID(t *testing.T) {
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		dao     *HsAuthPermissionDao
		args    args
		want    interface{}
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := tt.dao.FindByID(tt.args.id)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthPermissionDao.FindByID() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. HsAuthPermissionDao.FindByID() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestHsAuthPermissionDao_Add(t *testing.T) {
	tests := []struct {
		name    string
		dao     HsAuthPermissionDao
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := tt.dao.Add()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthPermissionDao.Add() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. HsAuthPermissionDao.Add() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestHsAuthPermissionDao_AddBatch(t *testing.T) {
	type args struct {
		obj []interface{}
	}
	tests := []struct {
		name    string
		dao     HsAuthPermissionDao
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := tt.dao.AddBatch(tt.args.obj); (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthPermissionDao.AddBatch() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestHsAuthPermissionDao_Update(t *testing.T) {
	tests := []struct {
		name    string
		dao     HsAuthPermissionDao
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := tt.dao.Update()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthPermissionDao.Update() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. HsAuthPermissionDao.Update() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestHsAuthPermissionDao_UpdateBatch(t *testing.T) {
	type args struct {
		obj []interface{}
	}
	tests := []struct {
		name    string
		dao     HsAuthPermissionDao
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := tt.dao.UpdateBatch(tt.args.obj); (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthPermissionDao.UpdateBatch() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestHsAuthPermissionDao_Delete(t *testing.T) {
	tests := []struct {
		name    string
		dao     HsAuthPermissionDao
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := tt.dao.Delete()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthPermissionDao.Delete() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. HsAuthPermissionDao.Delete() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestHsAuthPermissionDao_DeleteBatch(t *testing.T) {
	type args struct {
		obj []interface{}
	}
	tests := []struct {
		name    string
		dao     HsAuthPermissionDao
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := tt.dao.DeleteBatch(tt.args.obj); (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthPermissionDao.DeleteBatch() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestHsAuthPermissionDao_SDelete(t *testing.T) {
	tests := []struct {
		name    string
		dao     HsAuthPermissionDao
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := tt.dao.SDelete()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthPermissionDao.SDelete() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. HsAuthPermissionDao.SDelete() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestHsAuthPermissionDao_SDeleteBatch(t *testing.T) {
	type args struct {
		obj []interface{}
	}
	tests := []struct {
		name    string
		dao     HsAuthPermissionDao
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := tt.dao.SDeleteBatch(tt.args.obj); (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthPermissionDao.SDeleteBatch() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestHsAuthPermissionDao_Exec(t *testing.T) {
	type args struct {
		sql   string
		value []interface{}
	}
	tests := []struct {
		name    string
		dao     HsAuthPermissionDao
		args    args
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := tt.dao.Exec(tt.args.sql, tt.args.value...)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. HsAuthPermissionDao.Exec() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. HsAuthPermissionDao.Exec() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestHsAuthPermissionDao_GetSql(t *testing.T) {
	tests := []struct {
		name  string
		dao   HsAuthPermissionDao
		want  string
		want1 []interface{}
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, got1 := tt.dao.GetSql()
		if got != tt.want {
			t.Errorf("%q. HsAuthPermissionDao.GetSql() got = %v, want %v", tt.name, got, tt.want)
		}
		if !reflect.DeepEqual(got1, tt.want1) {
			t.Errorf("%q. HsAuthPermissionDao.GetSql() got1 = %v, want %v", tt.name, got1, tt.want1)
		}
	}
}

func TestHsAuthPermissionDao_SetDBConn(t *testing.T) {
	type args struct {
		db  string
		str string
	}
	tests := []struct {
		name string
		dao  HsAuthPermissionDao
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt.dao.SetDBConn(tt.args.db, tt.args.str)
	}
}

func TestHsAuthPermissionDao_getObjWithValue(t *testing.T) {
	type args struct {
		dao HsAuthPermissionDao
	}
	tests := []struct {
		name string
		daoo HsAuthPermissionDao
		args args
		want base.HsAuthPermission
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.daoo.getObjWithValue(tt.args.dao); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. HsAuthPermissionDao.getObjWithValue() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestNewHsAuthPermissionDao(t *testing.T) {
	tests := []struct {
		name string
		want HsAuthPermissionDao
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := NewHsAuthPermissionDao(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. NewHsAuthPermissionDao() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
