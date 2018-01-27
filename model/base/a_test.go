package base

import (
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func TestHsAuthPermission_checkAddCache(t *testing.T) {
	type fields struct {
		Id        int64
		AppKey    string
		ApiKey    string
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthPermission := HsAuthPermission{
				Id:        tt.fields.Id,
				AppKey:    tt.fields.AppKey,
				ApiKey:    tt.fields.ApiKey,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			hsAuthPermission.checkAddCache()
		})
	}
}

func TestHsAuthPermission_StartAddCache(t *testing.T) {
	type fields struct {
		Id        int64
		AppKey    string
		ApiKey    string
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthPermission := HsAuthPermission{
				Id:        tt.fields.Id,
				AppKey:    tt.fields.AppKey,
				ApiKey:    tt.fields.ApiKey,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			hsAuthPermission.StartAddCache()
		})
	}
}

func TestHsAuthPermission_GetSql(t *testing.T) {
	type fields struct {
		Id        int64
		AppKey    string
		ApiKey    string
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		want1  []interface{}
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthPermission := HsAuthPermission{
				Id:        tt.fields.Id,
				AppKey:    tt.fields.AppKey,
				ApiKey:    tt.fields.ApiKey,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			got, got1 := hsAuthPermission.GetSql()
			if got != tt.want {
				t.Errorf("HsAuthPermission.GetSql() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("HsAuthPermission.GetSql() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestHsAuthPermission_SetDBConn(t *testing.T) {
	type fields struct {
		Id        int64
		AppKey    string
		ApiKey    string
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	type args struct {
		db  string
		str string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthPermission := HsAuthPermission{
				Id:        tt.fields.Id,
				AppKey:    tt.fields.AppKey,
				ApiKey:    tt.fields.ApiKey,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			hsAuthPermission.SetDBConn(tt.args.db, tt.args.str)
		})
	}
}

func TestNewHsAuthPermission(t *testing.T) {
	tests := []struct {
		name string
		want HsAuthPermission
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHsAuthPermission(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHsAuthPermission() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getHsAuthPermissionArgsStr(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHsAuthPermissionArgsStr(tt.args.num); got != tt.want {
				t.Errorf("getHsAuthPermissionArgsStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getHsAuthPermissionArgsStrUpdate(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHsAuthPermissionArgsStrUpdate(); got != tt.want {
				t.Errorf("getHsAuthPermissionArgsStrUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHsAuthPermission_Select(t *testing.T) {
	type fields struct {
		Id        int64
		AppKey    string
		ApiKey    string
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	type args struct {
		sql    string
		limit  int
		offset int
		value  []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []interface{}
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthPermission := HsAuthPermission{
				Id:        tt.fields.Id,
				AppKey:    tt.fields.AppKey,
				ApiKey:    tt.fields.ApiKey,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			got, err := hsAuthPermission.Select(tt.args.sql, tt.args.limit, tt.args.offset, tt.args.value...)
			if (err != nil) != tt.wantErr {
				t.Errorf("HsAuthPermission.Select() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HsAuthPermission.Select() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHsAuthPermission_FindByID(t *testing.T) {
	type fields struct {
		Id        int64
		AppKey    string
		ApiKey    string
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthPermission := &HsAuthPermission{
				Id:        tt.fields.Id,
				AppKey:    tt.fields.AppKey,
				ApiKey:    tt.fields.ApiKey,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			got, err := hsAuthPermission.FindByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("HsAuthPermission.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HsAuthPermission.FindByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHsAuthPermission_Add(t *testing.T) {
	type fields struct {
		Id        int64
		AppKey    string
		ApiKey    string
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	tests := []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthPermission := HsAuthPermission{
				Id:        tt.fields.Id,
				AppKey:    tt.fields.AppKey,
				ApiKey:    tt.fields.ApiKey,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			got, err := hsAuthPermission.Add()
			if (err != nil) != tt.wantErr {
				t.Errorf("HsAuthPermission.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HsAuthPermission.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHsAuthPermission_AddBatch(t *testing.T) {
	type fields struct {
		Id        int64
		AppKey    string
		ApiKey    string
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	type args struct {
		obj []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthPermission := HsAuthPermission{
				Id:        tt.fields.Id,
				AppKey:    tt.fields.AppKey,
				ApiKey:    tt.fields.ApiKey,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			if err := hsAuthPermission.AddBatch(tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("HsAuthPermission.AddBatch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHsAuthPermission_Update(t *testing.T) {
	type fields struct {
		Id        int64
		AppKey    string
		ApiKey    string
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	tests := []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthPermission := &HsAuthPermission{
				Id:        tt.fields.Id,
				AppKey:    tt.fields.AppKey,
				ApiKey:    tt.fields.ApiKey,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			got, err := hsAuthPermission.Update()
			if (err != nil) != tt.wantErr {
				t.Errorf("HsAuthPermission.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HsAuthPermission.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHsAuthPermission_UpdateBatch(t *testing.T) {
	type fields struct {
		Id        int64
		AppKey    string
		ApiKey    string
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	type args struct {
		obj []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthPermission := HsAuthPermission{
				Id:        tt.fields.Id,
				AppKey:    tt.fields.AppKey,
				ApiKey:    tt.fields.ApiKey,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			if err := hsAuthPermission.UpdateBatch(tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("HsAuthPermission.UpdateBatch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHsAuthPermission_SDelete(t *testing.T) {
	type fields struct {
		Id        int64
		AppKey    string
		ApiKey    string
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	tests := []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthPermission := HsAuthPermission{
				Id:        tt.fields.Id,
				AppKey:    tt.fields.AppKey,
				ApiKey:    tt.fields.ApiKey,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			got, err := hsAuthPermission.SDelete()
			if (err != nil) != tt.wantErr {
				t.Errorf("HsAuthPermission.SDelete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HsAuthPermission.SDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHsAuthPermission_SDeleteBatch(t *testing.T) {
	type fields struct {
		Id        int64
		AppKey    string
		ApiKey    string
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	type args struct {
		obj []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthPermission := HsAuthPermission{
				Id:        tt.fields.Id,
				AppKey:    tt.fields.AppKey,
				ApiKey:    tt.fields.ApiKey,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			if err := hsAuthPermission.SDeleteBatch(tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("HsAuthPermission.SDeleteBatch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHsAuthPermission_Delete(t *testing.T) {
	type fields struct {
		Id        int64
		AppKey    string
		ApiKey    string
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	tests := []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthPermission := HsAuthPermission{
				Id:        tt.fields.Id,
				AppKey:    tt.fields.AppKey,
				ApiKey:    tt.fields.ApiKey,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			got, err := hsAuthPermission.Delete()
			if (err != nil) != tt.wantErr {
				t.Errorf("HsAuthPermission.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HsAuthPermission.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHsAuthPermission_DeleteBatch(t *testing.T) {
	type fields struct {
		Id        int64
		AppKey    string
		ApiKey    string
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	type args struct {
		obj []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthPermission := HsAuthPermission{
				Id:        tt.fields.Id,
				AppKey:    tt.fields.AppKey,
				ApiKey:    tt.fields.ApiKey,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			if err := hsAuthPermission.DeleteBatch(tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("HsAuthPermission.DeleteBatch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHsAuthPermission_Exec(t *testing.T) {
	type fields struct {
		Id        int64
		AppKey    string
		ApiKey    string
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	type args struct {
		sql   string
		value []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthPermission := HsAuthPermission{
				Id:        tt.fields.Id,
				AppKey:    tt.fields.AppKey,
				ApiKey:    tt.fields.ApiKey,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			got, err := hsAuthPermission.Exec(tt.args.sql, tt.args.value...)
			if (err != nil) != tt.wantErr {
				t.Errorf("HsAuthPermission.Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HsAuthPermission.Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}
package base

import (
	"reflect"
	"testing"
)

func TestSetCacheType(t *testing.T) {
	type args struct {
		ctype string
		clen  int
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetCacheType(tt.args.ctype, tt.args.clen)
		})
	}
}

func TestSetCacheTime(t *testing.T) {
	type args struct {
		t int
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetCacheTime(tt.args.t)
		})
	}
}

func TestGetCacheTime(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCacheTime(); got != tt.want {
				t.Errorf("GetCacheTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetCacheLen(t *testing.T) {
	type args struct {
		l int
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetCacheLen(tt.args.l)
		})
	}
}

func TestGetCacheLen(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCacheLen(); got != tt.want {
				t.Errorf("GetCacheLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCacheUsedLen(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCacheUsedLen(); got != tt.want {
				t.Errorf("GetCacheUsedLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCache(t *testing.T) {
	type args struct {
		uc bool
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UseCache(tt.args.uc)
		})
	}
}

func TestCacheUsed(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CacheUsed(); got != tt.want {
				t.Errorf("CacheUsed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCacheRate(t *testing.T) {
	tests := []struct {
		name string
		want float64
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCacheRate(); got != tt.want {
				t.Errorf("GetCacheRate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetCache(t *testing.T) {
	type args struct {
		k interface{}
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SetCache(tt.args.k, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("SetCache() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetCache(t *testing.T) {
	type args struct {
		k interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCache(tt.args.k)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCache() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCache() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetConn(t *testing.T) {
	type args struct {
		db  string
		str string
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetConn(tt.args.db, tt.args.str)
		})
	}
}

func TestCheckerr(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Checkerr(tt.args.err); (err != nil) != tt.wantErr {
				t.Errorf("Checkerr() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAddBeforeFun(t *testing.T) {
	type args struct {
		f func()
		w string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddBeforeFun(tt.args.f, tt.args.w); got != tt.want {
				t.Errorf("AddBeforeFun() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddAfterFun(t *testing.T) {
	type args struct {
		f func()
		w string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddAfterFun(tt.args.f, tt.args.w); got != tt.want {
				t.Errorf("AddAfterFun() = %v, want %v", got, tt.want)
			}
		})
	}
}
Generated TestHsAuthRecords_checkAddCache
Generated TestHsAuthRecords_StartAddCache
Generated TestHsAuthRecords_GetSql
Generated TestHsAuthRecords_SetDBConn
Generated TestNewHsAuthRecords
Generated Test_getHsAuthRecordsArgsStr
Generated Test_getHsAuthRecordsArgsStrUpdate
Generated TestHsAuthRecords_Select
Generated TestHsAuthRecords_FindByID
Generated TestHsAuthRecords_Add
Generated TestHsAuthRecords_AddBatch
Generated TestHsAuthRecords_Update
Generated TestHsAuthRecords_UpdateBatch
Generated TestHsAuthRecords_SDelete
Generated TestHsAuthRecords_SDeleteBatch
Generated TestHsAuthRecords_Delete
Generated TestHsAuthRecords_DeleteBatch
Generated TestHsAuthRecords_Exec
package base

import (
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func TestHsAuthRecords_checkAddCache(t *testing.T) {
	type fields struct {
		Id        int64
		SecretKey string
		AppKey    string
		Sign      string
		Token     string
		Alg       string
		Ip        string
		Exp       string
		Iat       string
		Type      int8
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthRecords := HsAuthRecords{
				Id:        tt.fields.Id,
				SecretKey: tt.fields.SecretKey,
				AppKey:    tt.fields.AppKey,
				Sign:      tt.fields.Sign,
				Token:     tt.fields.Token,
				Alg:       tt.fields.Alg,
				Ip:        tt.fields.Ip,
				Exp:       tt.fields.Exp,
				Iat:       tt.fields.Iat,
				Type:      tt.fields.Type,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			hsAuthRecords.checkAddCache()
		})
	}
}

func TestHsAuthRecords_StartAddCache(t *testing.T) {
	type fields struct {
		Id        int64
		SecretKey string
		AppKey    string
		Sign      string
		Token     string
		Alg       string
		Ip        string
		Exp       string
		Iat       string
		Type      int8
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthRecords := HsAuthRecords{
				Id:        tt.fields.Id,
				SecretKey: tt.fields.SecretKey,
				AppKey:    tt.fields.AppKey,
				Sign:      tt.fields.Sign,
				Token:     tt.fields.Token,
				Alg:       tt.fields.Alg,
				Ip:        tt.fields.Ip,
				Exp:       tt.fields.Exp,
				Iat:       tt.fields.Iat,
				Type:      tt.fields.Type,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			hsAuthRecords.StartAddCache()
		})
	}
}

func TestHsAuthRecords_GetSql(t *testing.T) {
	type fields struct {
		Id        int64
		SecretKey string
		AppKey    string
		Sign      string
		Token     string
		Alg       string
		Ip        string
		Exp       string
		Iat       string
		Type      int8
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		want1  []interface{}
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthRecords := HsAuthRecords{
				Id:        tt.fields.Id,
				SecretKey: tt.fields.SecretKey,
				AppKey:    tt.fields.AppKey,
				Sign:      tt.fields.Sign,
				Token:     tt.fields.Token,
				Alg:       tt.fields.Alg,
				Ip:        tt.fields.Ip,
				Exp:       tt.fields.Exp,
				Iat:       tt.fields.Iat,
				Type:      tt.fields.Type,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			got, got1 := hsAuthRecords.GetSql()
			if got != tt.want {
				t.Errorf("HsAuthRecords.GetSql() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("HsAuthRecords.GetSql() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestHsAuthRecords_SetDBConn(t *testing.T) {
	type fields struct {
		Id        int64
		SecretKey string
		AppKey    string
		Sign      string
		Token     string
		Alg       string
		Ip        string
		Exp       string
		Iat       string
		Type      int8
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	type args struct {
		db  string
		str string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthRecords := HsAuthRecords{
				Id:        tt.fields.Id,
				SecretKey: tt.fields.SecretKey,
				AppKey:    tt.fields.AppKey,
				Sign:      tt.fields.Sign,
				Token:     tt.fields.Token,
				Alg:       tt.fields.Alg,
				Ip:        tt.fields.Ip,
				Exp:       tt.fields.Exp,
				Iat:       tt.fields.Iat,
				Type:      tt.fields.Type,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			hsAuthRecords.SetDBConn(tt.args.db, tt.args.str)
		})
	}
}

func TestNewHsAuthRecords(t *testing.T) {
	tests := []struct {
		name string
		want HsAuthRecords
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHsAuthRecords(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHsAuthRecords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getHsAuthRecordsArgsStr(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHsAuthRecordsArgsStr(tt.args.num); got != tt.want {
				t.Errorf("getHsAuthRecordsArgsStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getHsAuthRecordsArgsStrUpdate(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHsAuthRecordsArgsStrUpdate(); got != tt.want {
				t.Errorf("getHsAuthRecordsArgsStrUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHsAuthRecords_Select(t *testing.T) {
	type fields struct {
		Id        int64
		SecretKey string
		AppKey    string
		Sign      string
		Token     string
		Alg       string
		Ip        string
		Exp       string
		Iat       string
		Type      int8
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	type args struct {
		sql    string
		limit  int
		offset int
		value  []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []interface{}
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthRecords := HsAuthRecords{
				Id:        tt.fields.Id,
				SecretKey: tt.fields.SecretKey,
				AppKey:    tt.fields.AppKey,
				Sign:      tt.fields.Sign,
				Token:     tt.fields.Token,
				Alg:       tt.fields.Alg,
				Ip:        tt.fields.Ip,
				Exp:       tt.fields.Exp,
				Iat:       tt.fields.Iat,
				Type:      tt.fields.Type,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			got, err := hsAuthRecords.Select(tt.args.sql, tt.args.limit, tt.args.offset, tt.args.value...)
			if (err != nil) != tt.wantErr {
				t.Errorf("HsAuthRecords.Select() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HsAuthRecords.Select() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHsAuthRecords_FindByID(t *testing.T) {
	type fields struct {
		Id        int64
		SecretKey string
		AppKey    string
		Sign      string
		Token     string
		Alg       string
		Ip        string
		Exp       string
		Iat       string
		Type      int8
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthRecords := &HsAuthRecords{
				Id:        tt.fields.Id,
				SecretKey: tt.fields.SecretKey,
				AppKey:    tt.fields.AppKey,
				Sign:      tt.fields.Sign,
				Token:     tt.fields.Token,
				Alg:       tt.fields.Alg,
				Ip:        tt.fields.Ip,
				Exp:       tt.fields.Exp,
				Iat:       tt.fields.Iat,
				Type:      tt.fields.Type,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			got, err := hsAuthRecords.FindByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("HsAuthRecords.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HsAuthRecords.FindByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHsAuthRecords_Add(t *testing.T) {
	type fields struct {
		Id        int64
		SecretKey string
		AppKey    string
		Sign      string
		Token     string
		Alg       string
		Ip        string
		Exp       string
		Iat       string
		Type      int8
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	tests := []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthRecords := HsAuthRecords{
				Id:        tt.fields.Id,
				SecretKey: tt.fields.SecretKey,
				AppKey:    tt.fields.AppKey,
				Sign:      tt.fields.Sign,
				Token:     tt.fields.Token,
				Alg:       tt.fields.Alg,
				Ip:        tt.fields.Ip,
				Exp:       tt.fields.Exp,
				Iat:       tt.fields.Iat,
				Type:      tt.fields.Type,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			got, err := hsAuthRecords.Add()
			if (err != nil) != tt.wantErr {
				t.Errorf("HsAuthRecords.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HsAuthRecords.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHsAuthRecords_AddBatch(t *testing.T) {
	type fields struct {
		Id        int64
		SecretKey string
		AppKey    string
		Sign      string
		Token     string
		Alg       string
		Ip        string
		Exp       string
		Iat       string
		Type      int8
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	type args struct {
		obj []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthRecords := HsAuthRecords{
				Id:        tt.fields.Id,
				SecretKey: tt.fields.SecretKey,
				AppKey:    tt.fields.AppKey,
				Sign:      tt.fields.Sign,
				Token:     tt.fields.Token,
				Alg:       tt.fields.Alg,
				Ip:        tt.fields.Ip,
				Exp:       tt.fields.Exp,
				Iat:       tt.fields.Iat,
				Type:      tt.fields.Type,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			if err := hsAuthRecords.AddBatch(tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("HsAuthRecords.AddBatch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHsAuthRecords_Update(t *testing.T) {
	type fields struct {
		Id        int64
		SecretKey string
		AppKey    string
		Sign      string
		Token     string
		Alg       string
		Ip        string
		Exp       string
		Iat       string
		Type      int8
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	tests := []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthRecords := &HsAuthRecords{
				Id:        tt.fields.Id,
				SecretKey: tt.fields.SecretKey,
				AppKey:    tt.fields.AppKey,
				Sign:      tt.fields.Sign,
				Token:     tt.fields.Token,
				Alg:       tt.fields.Alg,
				Ip:        tt.fields.Ip,
				Exp:       tt.fields.Exp,
				Iat:       tt.fields.Iat,
				Type:      tt.fields.Type,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			got, err := hsAuthRecords.Update()
			if (err != nil) != tt.wantErr {
				t.Errorf("HsAuthRecords.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HsAuthRecords.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHsAuthRecords_UpdateBatch(t *testing.T) {
	type fields struct {
		Id        int64
		SecretKey string
		AppKey    string
		Sign      string
		Token     string
		Alg       string
		Ip        string
		Exp       string
		Iat       string
		Type      int8
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	type args struct {
		obj []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthRecords := HsAuthRecords{
				Id:        tt.fields.Id,
				SecretKey: tt.fields.SecretKey,
				AppKey:    tt.fields.AppKey,
				Sign:      tt.fields.Sign,
				Token:     tt.fields.Token,
				Alg:       tt.fields.Alg,
				Ip:        tt.fields.Ip,
				Exp:       tt.fields.Exp,
				Iat:       tt.fields.Iat,
				Type:      tt.fields.Type,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			if err := hsAuthRecords.UpdateBatch(tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("HsAuthRecords.UpdateBatch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHsAuthRecords_SDelete(t *testing.T) {
	type fields struct {
		Id        int64
		SecretKey string
		AppKey    string
		Sign      string
		Token     string
		Alg       string
		Ip        string
		Exp       string
		Iat       string
		Type      int8
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	tests := []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthRecords := HsAuthRecords{
				Id:        tt.fields.Id,
				SecretKey: tt.fields.SecretKey,
				AppKey:    tt.fields.AppKey,
				Sign:      tt.fields.Sign,
				Token:     tt.fields.Token,
				Alg:       tt.fields.Alg,
				Ip:        tt.fields.Ip,
				Exp:       tt.fields.Exp,
				Iat:       tt.fields.Iat,
				Type:      tt.fields.Type,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			got, err := hsAuthRecords.SDelete()
			if (err != nil) != tt.wantErr {
				t.Errorf("HsAuthRecords.SDelete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HsAuthRecords.SDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHsAuthRecords_SDeleteBatch(t *testing.T) {
	type fields struct {
		Id        int64
		SecretKey string
		AppKey    string
		Sign      string
		Token     string
		Alg       string
		Ip        string
		Exp       string
		Iat       string
		Type      int8
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	type args struct {
		obj []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthRecords := HsAuthRecords{
				Id:        tt.fields.Id,
				SecretKey: tt.fields.SecretKey,
				AppKey:    tt.fields.AppKey,
				Sign:      tt.fields.Sign,
				Token:     tt.fields.Token,
				Alg:       tt.fields.Alg,
				Ip:        tt.fields.Ip,
				Exp:       tt.fields.Exp,
				Iat:       tt.fields.Iat,
				Type:      tt.fields.Type,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			if err := hsAuthRecords.SDeleteBatch(tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("HsAuthRecords.SDeleteBatch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHsAuthRecords_Delete(t *testing.T) {
	type fields struct {
		Id        int64
		SecretKey string
		AppKey    string
		Sign      string
		Token     string
		Alg       string
		Ip        string
		Exp       string
		Iat       string
		Type      int8
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	tests := []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthRecords := HsAuthRecords{
				Id:        tt.fields.Id,
				SecretKey: tt.fields.SecretKey,
				AppKey:    tt.fields.AppKey,
				Sign:      tt.fields.Sign,
				Token:     tt.fields.Token,
				Alg:       tt.fields.Alg,
				Ip:        tt.fields.Ip,
				Exp:       tt.fields.Exp,
				Iat:       tt.fields.Iat,
				Type:      tt.fields.Type,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			got, err := hsAuthRecords.Delete()
			if (err != nil) != tt.wantErr {
				t.Errorf("HsAuthRecords.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HsAuthRecords.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHsAuthRecords_DeleteBatch(t *testing.T) {
	type fields struct {
		Id        int64
		SecretKey string
		AppKey    string
		Sign      string
		Token     string
		Alg       string
		Ip        string
		Exp       string
		Iat       string
		Type      int8
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	type args struct {
		obj []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthRecords := HsAuthRecords{
				Id:        tt.fields.Id,
				SecretKey: tt.fields.SecretKey,
				AppKey:    tt.fields.AppKey,
				Sign:      tt.fields.Sign,
				Token:     tt.fields.Token,
				Alg:       tt.fields.Alg,
				Ip:        tt.fields.Ip,
				Exp:       tt.fields.Exp,
				Iat:       tt.fields.Iat,
				Type:      tt.fields.Type,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			if err := hsAuthRecords.DeleteBatch(tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("HsAuthRecords.DeleteBatch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHsAuthRecords_Exec(t *testing.T) {
	type fields struct {
		Id        int64
		SecretKey string
		AppKey    string
		Sign      string
		Token     string
		Alg       string
		Ip        string
		Exp       string
		Iat       string
		Type      int8
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	type args struct {
		sql   string
		value []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthRecords := HsAuthRecords{
				Id:        tt.fields.Id,
				SecretKey: tt.fields.SecretKey,
				AppKey:    tt.fields.AppKey,
				Sign:      tt.fields.Sign,
				Token:     tt.fields.Token,
				Alg:       tt.fields.Alg,
				Ip:        tt.fields.Ip,
				Exp:       tt.fields.Exp,
				Iat:       tt.fields.Iat,
				Type:      tt.fields.Type,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			got, err := hsAuthRecords.Exec(tt.args.sql, tt.args.value...)
			if (err != nil) != tt.wantErr {
				t.Errorf("HsAuthRecords.Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HsAuthRecords.Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}
Generated TestHsAuthApi_checkAddCache
Generated TestHsAuthApi_StartAddCache
Generated TestHsAuthApi_GetSql
Generated TestHsAuthApi_SetDBConn
Generated TestNewHsAuthApi
Generated Test_getHsAuthApiArgsStr
Generated Test_getHsAuthApiArgsStrUpdate
Generated TestHsAuthApi_Select
Generated TestHsAuthApi_FindByID
Generated TestHsAuthApi_Add
Generated TestHsAuthApi_AddBatch
Generated TestHsAuthApi_Update
Generated TestHsAuthApi_UpdateBatch
Generated TestHsAuthApi_SDelete
Generated TestHsAuthApi_SDeleteBatch
Generated TestHsAuthApi_Delete
Generated TestHsAuthApi_DeleteBatch
Generated TestHsAuthApi_Exec
package base

import (
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func TestHsAuthApi_checkAddCache(t *testing.T) {
	type fields struct {
		Id        int64
		ApiKey    string
		Name      string
		Type      int8
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthApi := HsAuthApi{
				Id:        tt.fields.Id,
				ApiKey:    tt.fields.ApiKey,
				Name:      tt.fields.Name,
				Type:      tt.fields.Type,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			hsAuthApi.checkAddCache()
		})
	}
}

func TestHsAuthApi_StartAddCache(t *testing.T) {
	type fields struct {
		Id        int64
		ApiKey    string
		Name      string
		Type      int8
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthApi := HsAuthApi{
				Id:        tt.fields.Id,
				ApiKey:    tt.fields.ApiKey,
				Name:      tt.fields.Name,
				Type:      tt.fields.Type,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			hsAuthApi.StartAddCache()
		})
	}
}

func TestHsAuthApi_GetSql(t *testing.T) {
	type fields struct {
		Id        int64
		ApiKey    string
		Name      string
		Type      int8
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		want1  []interface{}
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthApi := HsAuthApi{
				Id:        tt.fields.Id,
				ApiKey:    tt.fields.ApiKey,
				Name:      tt.fields.Name,
				Type:      tt.fields.Type,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			got, got1 := hsAuthApi.GetSql()
			if got != tt.want {
				t.Errorf("HsAuthApi.GetSql() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("HsAuthApi.GetSql() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestHsAuthApi_SetDBConn(t *testing.T) {
	type fields struct {
		Id        int64
		ApiKey    string
		Name      string
		Type      int8
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	type args struct {
		db  string
		str string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthApi := HsAuthApi{
				Id:        tt.fields.Id,
				ApiKey:    tt.fields.ApiKey,
				Name:      tt.fields.Name,
				Type:      tt.fields.Type,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			hsAuthApi.SetDBConn(tt.args.db, tt.args.str)
		})
	}
}

func TestNewHsAuthApi(t *testing.T) {
	tests := []struct {
		name string
		want HsAuthApi
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHsAuthApi(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHsAuthApi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getHsAuthApiArgsStr(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHsAuthApiArgsStr(tt.args.num); got != tt.want {
				t.Errorf("getHsAuthApiArgsStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getHsAuthApiArgsStrUpdate(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHsAuthApiArgsStrUpdate(); got != tt.want {
				t.Errorf("getHsAuthApiArgsStrUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHsAuthApi_Select(t *testing.T) {
	type fields struct {
		Id        int64
		ApiKey    string
		Name      string
		Type      int8
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	type args struct {
		sql    string
		limit  int
		offset int
		value  []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []interface{}
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthApi := HsAuthApi{
				Id:        tt.fields.Id,
				ApiKey:    tt.fields.ApiKey,
				Name:      tt.fields.Name,
				Type:      tt.fields.Type,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			got, err := hsAuthApi.Select(tt.args.sql, tt.args.limit, tt.args.offset, tt.args.value...)
			if (err != nil) != tt.wantErr {
				t.Errorf("HsAuthApi.Select() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HsAuthApi.Select() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHsAuthApi_FindByID(t *testing.T) {
	type fields struct {
		Id        int64
		ApiKey    string
		Name      string
		Type      int8
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthApi := &HsAuthApi{
				Id:        tt.fields.Id,
				ApiKey:    tt.fields.ApiKey,
				Name:      tt.fields.Name,
				Type:      tt.fields.Type,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			got, err := hsAuthApi.FindByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("HsAuthApi.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HsAuthApi.FindByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHsAuthApi_Add(t *testing.T) {
	type fields struct {
		Id        int64
		ApiKey    string
		Name      string
		Type      int8
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	tests := []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthApi := HsAuthApi{
				Id:        tt.fields.Id,
				ApiKey:    tt.fields.ApiKey,
				Name:      tt.fields.Name,
				Type:      tt.fields.Type,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			got, err := hsAuthApi.Add()
			if (err != nil) != tt.wantErr {
				t.Errorf("HsAuthApi.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HsAuthApi.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHsAuthApi_AddBatch(t *testing.T) {
	type fields struct {
		Id        int64
		ApiKey    string
		Name      string
		Type      int8
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	type args struct {
		obj []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthApi := HsAuthApi{
				Id:        tt.fields.Id,
				ApiKey:    tt.fields.ApiKey,
				Name:      tt.fields.Name,
				Type:      tt.fields.Type,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			if err := hsAuthApi.AddBatch(tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("HsAuthApi.AddBatch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHsAuthApi_Update(t *testing.T) {
	type fields struct {
		Id        int64
		ApiKey    string
		Name      string
		Type      int8
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	tests := []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthApi := &HsAuthApi{
				Id:        tt.fields.Id,
				ApiKey:    tt.fields.ApiKey,
				Name:      tt.fields.Name,
				Type:      tt.fields.Type,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			got, err := hsAuthApi.Update()
			if (err != nil) != tt.wantErr {
				t.Errorf("HsAuthApi.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HsAuthApi.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHsAuthApi_UpdateBatch(t *testing.T) {
	type fields struct {
		Id        int64
		ApiKey    string
		Name      string
		Type      int8
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	type args struct {
		obj []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthApi := HsAuthApi{
				Id:        tt.fields.Id,
				ApiKey:    tt.fields.ApiKey,
				Name:      tt.fields.Name,
				Type:      tt.fields.Type,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			if err := hsAuthApi.UpdateBatch(tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("HsAuthApi.UpdateBatch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHsAuthApi_SDelete(t *testing.T) {
	type fields struct {
		Id        int64
		ApiKey    string
		Name      string
		Type      int8
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	tests := []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthApi := HsAuthApi{
				Id:        tt.fields.Id,
				ApiKey:    tt.fields.ApiKey,
				Name:      tt.fields.Name,
				Type:      tt.fields.Type,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			got, err := hsAuthApi.SDelete()
			if (err != nil) != tt.wantErr {
				t.Errorf("HsAuthApi.SDelete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HsAuthApi.SDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHsAuthApi_SDeleteBatch(t *testing.T) {
	type fields struct {
		Id        int64
		ApiKey    string
		Name      string
		Type      int8
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	type args struct {
		obj []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthApi := HsAuthApi{
				Id:        tt.fields.Id,
				ApiKey:    tt.fields.ApiKey,
				Name:      tt.fields.Name,
				Type:      tt.fields.Type,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			if err := hsAuthApi.SDeleteBatch(tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("HsAuthApi.SDeleteBatch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHsAuthApi_Delete(t *testing.T) {
	type fields struct {
		Id        int64
		ApiKey    string
		Name      string
		Type      int8
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	tests := []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthApi := HsAuthApi{
				Id:        tt.fields.Id,
				ApiKey:    tt.fields.ApiKey,
				Name:      tt.fields.Name,
				Type:      tt.fields.Type,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			got, err := hsAuthApi.Delete()
			if (err != nil) != tt.wantErr {
				t.Errorf("HsAuthApi.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HsAuthApi.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHsAuthApi_DeleteBatch(t *testing.T) {
	type fields struct {
		Id        int64
		ApiKey    string
		Name      string
		Type      int8
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	type args struct {
		obj []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthApi := HsAuthApi{
				Id:        tt.fields.Id,
				ApiKey:    tt.fields.ApiKey,
				Name:      tt.fields.Name,
				Type:      tt.fields.Type,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			if err := hsAuthApi.DeleteBatch(tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("HsAuthApi.DeleteBatch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHsAuthApi_Exec(t *testing.T) {
	type fields struct {
		Id        int64
		ApiKey    string
		Name      string
		Type      int8
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	type args struct {
		sql   string
		value []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthApi := HsAuthApi{
				Id:        tt.fields.Id,
				ApiKey:    tt.fields.ApiKey,
				Name:      tt.fields.Name,
				Type:      tt.fields.Type,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			got, err := hsAuthApi.Exec(tt.args.sql, tt.args.value...)
			if (err != nil) != tt.wantErr {
				t.Errorf("HsAuthApi.Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HsAuthApi.Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}
Generated TestHsAuthApplication_checkAddCache
Generated TestHsAuthApplication_StartAddCache
Generated TestHsAuthApplication_GetSql
Generated TestHsAuthApplication_SetDBConn
Generated TestNewHsAuthApplication
Generated Test_getHsAuthApplicationArgsStr
Generated Test_getHsAuthApplicationArgsStrUpdate
Generated TestHsAuthApplication_Select
Generated TestHsAuthApplication_FindByID
Generated TestHsAuthApplication_Add
Generated TestHsAuthApplication_AddBatch
Generated TestHsAuthApplication_Update
Generated TestHsAuthApplication_UpdateBatch
Generated TestHsAuthApplication_SDelete
Generated TestHsAuthApplication_SDeleteBatch
Generated TestHsAuthApplication_Delete
Generated TestHsAuthApplication_DeleteBatch
Generated TestHsAuthApplication_Exec
package base

import (
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func TestHsAuthApplication_checkAddCache(t *testing.T) {
	type fields struct {
		Id        int64
		SecretKey string
		AppKey    string
		Name      string
		Ip        string
		Type      int8
		Exp       int64
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthApplication := HsAuthApplication{
				Id:        tt.fields.Id,
				SecretKey: tt.fields.SecretKey,
				AppKey:    tt.fields.AppKey,
				Name:      tt.fields.Name,
				Ip:        tt.fields.Ip,
				Type:      tt.fields.Type,
				Exp:       tt.fields.Exp,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			hsAuthApplication.checkAddCache()
		})
	}
}

func TestHsAuthApplication_StartAddCache(t *testing.T) {
	type fields struct {
		Id        int64
		SecretKey string
		AppKey    string
		Name      string
		Ip        string
		Type      int8
		Exp       int64
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthApplication := HsAuthApplication{
				Id:        tt.fields.Id,
				SecretKey: tt.fields.SecretKey,
				AppKey:    tt.fields.AppKey,
				Name:      tt.fields.Name,
				Ip:        tt.fields.Ip,
				Type:      tt.fields.Type,
				Exp:       tt.fields.Exp,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			hsAuthApplication.StartAddCache()
		})
	}
}

func TestHsAuthApplication_GetSql(t *testing.T) {
	type fields struct {
		Id        int64
		SecretKey string
		AppKey    string
		Name      string
		Ip        string
		Type      int8
		Exp       int64
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		want1  []interface{}
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthApplication := HsAuthApplication{
				Id:        tt.fields.Id,
				SecretKey: tt.fields.SecretKey,
				AppKey:    tt.fields.AppKey,
				Name:      tt.fields.Name,
				Ip:        tt.fields.Ip,
				Type:      tt.fields.Type,
				Exp:       tt.fields.Exp,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			got, got1 := hsAuthApplication.GetSql()
			if got != tt.want {
				t.Errorf("HsAuthApplication.GetSql() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("HsAuthApplication.GetSql() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestHsAuthApplication_SetDBConn(t *testing.T) {
	type fields struct {
		Id        int64
		SecretKey string
		AppKey    string
		Name      string
		Ip        string
		Type      int8
		Exp       int64
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	type args struct {
		db  string
		str string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthApplication := HsAuthApplication{
				Id:        tt.fields.Id,
				SecretKey: tt.fields.SecretKey,
				AppKey:    tt.fields.AppKey,
				Name:      tt.fields.Name,
				Ip:        tt.fields.Ip,
				Type:      tt.fields.Type,
				Exp:       tt.fields.Exp,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			hsAuthApplication.SetDBConn(tt.args.db, tt.args.str)
		})
	}
}

func TestNewHsAuthApplication(t *testing.T) {
	tests := []struct {
		name string
		want HsAuthApplication
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHsAuthApplication(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHsAuthApplication() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getHsAuthApplicationArgsStr(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHsAuthApplicationArgsStr(tt.args.num); got != tt.want {
				t.Errorf("getHsAuthApplicationArgsStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getHsAuthApplicationArgsStrUpdate(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHsAuthApplicationArgsStrUpdate(); got != tt.want {
				t.Errorf("getHsAuthApplicationArgsStrUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHsAuthApplication_Select(t *testing.T) {
	type fields struct {
		Id        int64
		SecretKey string
		AppKey    string
		Name      string
		Ip        string
		Type      int8
		Exp       int64
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	type args struct {
		sql    string
		limit  int
		offset int
		value  []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []interface{}
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthApplication := HsAuthApplication{
				Id:        tt.fields.Id,
				SecretKey: tt.fields.SecretKey,
				AppKey:    tt.fields.AppKey,
				Name:      tt.fields.Name,
				Ip:        tt.fields.Ip,
				Type:      tt.fields.Type,
				Exp:       tt.fields.Exp,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			got, err := hsAuthApplication.Select(tt.args.sql, tt.args.limit, tt.args.offset, tt.args.value...)
			if (err != nil) != tt.wantErr {
				t.Errorf("HsAuthApplication.Select() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HsAuthApplication.Select() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHsAuthApplication_FindByID(t *testing.T) {
	type fields struct {
		Id        int64
		SecretKey string
		AppKey    string
		Name      string
		Ip        string
		Type      int8
		Exp       int64
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthApplication := &HsAuthApplication{
				Id:        tt.fields.Id,
				SecretKey: tt.fields.SecretKey,
				AppKey:    tt.fields.AppKey,
				Name:      tt.fields.Name,
				Ip:        tt.fields.Ip,
				Type:      tt.fields.Type,
				Exp:       tt.fields.Exp,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			got, err := hsAuthApplication.FindByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("HsAuthApplication.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HsAuthApplication.FindByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHsAuthApplication_Add(t *testing.T) {
	type fields struct {
		Id        int64
		SecretKey string
		AppKey    string
		Name      string
		Ip        string
		Type      int8
		Exp       int64
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	tests := []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthApplication := HsAuthApplication{
				Id:        tt.fields.Id,
				SecretKey: tt.fields.SecretKey,
				AppKey:    tt.fields.AppKey,
				Name:      tt.fields.Name,
				Ip:        tt.fields.Ip,
				Type:      tt.fields.Type,
				Exp:       tt.fields.Exp,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			got, err := hsAuthApplication.Add()
			if (err != nil) != tt.wantErr {
				t.Errorf("HsAuthApplication.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HsAuthApplication.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHsAuthApplication_AddBatch(t *testing.T) {
	type fields struct {
		Id        int64
		SecretKey string
		AppKey    string
		Name      string
		Ip        string
		Type      int8
		Exp       int64
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	type args struct {
		obj []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthApplication := HsAuthApplication{
				Id:        tt.fields.Id,
				SecretKey: tt.fields.SecretKey,
				AppKey:    tt.fields.AppKey,
				Name:      tt.fields.Name,
				Ip:        tt.fields.Ip,
				Type:      tt.fields.Type,
				Exp:       tt.fields.Exp,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			if err := hsAuthApplication.AddBatch(tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("HsAuthApplication.AddBatch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHsAuthApplication_Update(t *testing.T) {
	type fields struct {
		Id        int64
		SecretKey string
		AppKey    string
		Name      string
		Ip        string
		Type      int8
		Exp       int64
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	tests := []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthApplication := &HsAuthApplication{
				Id:        tt.fields.Id,
				SecretKey: tt.fields.SecretKey,
				AppKey:    tt.fields.AppKey,
				Name:      tt.fields.Name,
				Ip:        tt.fields.Ip,
				Type:      tt.fields.Type,
				Exp:       tt.fields.Exp,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			got, err := hsAuthApplication.Update()
			if (err != nil) != tt.wantErr {
				t.Errorf("HsAuthApplication.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HsAuthApplication.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHsAuthApplication_UpdateBatch(t *testing.T) {
	type fields struct {
		Id        int64
		SecretKey string
		AppKey    string
		Name      string
		Ip        string
		Type      int8
		Exp       int64
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	type args struct {
		obj []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthApplication := HsAuthApplication{
				Id:        tt.fields.Id,
				SecretKey: tt.fields.SecretKey,
				AppKey:    tt.fields.AppKey,
				Name:      tt.fields.Name,
				Ip:        tt.fields.Ip,
				Type:      tt.fields.Type,
				Exp:       tt.fields.Exp,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			if err := hsAuthApplication.UpdateBatch(tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("HsAuthApplication.UpdateBatch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHsAuthApplication_SDelete(t *testing.T) {
	type fields struct {
		Id        int64
		SecretKey string
		AppKey    string
		Name      string
		Ip        string
		Type      int8
		Exp       int64
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	tests := []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthApplication := HsAuthApplication{
				Id:        tt.fields.Id,
				SecretKey: tt.fields.SecretKey,
				AppKey:    tt.fields.AppKey,
				Name:      tt.fields.Name,
				Ip:        tt.fields.Ip,
				Type:      tt.fields.Type,
				Exp:       tt.fields.Exp,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			got, err := hsAuthApplication.SDelete()
			if (err != nil) != tt.wantErr {
				t.Errorf("HsAuthApplication.SDelete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HsAuthApplication.SDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHsAuthApplication_SDeleteBatch(t *testing.T) {
	type fields struct {
		Id        int64
		SecretKey string
		AppKey    string
		Name      string
		Ip        string
		Type      int8
		Exp       int64
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	type args struct {
		obj []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthApplication := HsAuthApplication{
				Id:        tt.fields.Id,
				SecretKey: tt.fields.SecretKey,
				AppKey:    tt.fields.AppKey,
				Name:      tt.fields.Name,
				Ip:        tt.fields.Ip,
				Type:      tt.fields.Type,
				Exp:       tt.fields.Exp,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			if err := hsAuthApplication.SDeleteBatch(tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("HsAuthApplication.SDeleteBatch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHsAuthApplication_Delete(t *testing.T) {
	type fields struct {
		Id        int64
		SecretKey string
		AppKey    string
		Name      string
		Ip        string
		Type      int8
		Exp       int64
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	tests := []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthApplication := HsAuthApplication{
				Id:        tt.fields.Id,
				SecretKey: tt.fields.SecretKey,
				AppKey:    tt.fields.AppKey,
				Name:      tt.fields.Name,
				Ip:        tt.fields.Ip,
				Type:      tt.fields.Type,
				Exp:       tt.fields.Exp,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			got, err := hsAuthApplication.Delete()
			if (err != nil) != tt.wantErr {
				t.Errorf("HsAuthApplication.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HsAuthApplication.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHsAuthApplication_DeleteBatch(t *testing.T) {
	type fields struct {
		Id        int64
		SecretKey string
		AppKey    string
		Name      string
		Ip        string
		Type      int8
		Exp       int64
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	type args struct {
		obj []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthApplication := HsAuthApplication{
				Id:        tt.fields.Id,
				SecretKey: tt.fields.SecretKey,
				AppKey:    tt.fields.AppKey,
				Name:      tt.fields.Name,
				Ip:        tt.fields.Ip,
				Type:      tt.fields.Type,
				Exp:       tt.fields.Exp,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			if err := hsAuthApplication.DeleteBatch(tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("HsAuthApplication.DeleteBatch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHsAuthApplication_Exec(t *testing.T) {
	type fields struct {
		Id        int64
		SecretKey string
		AppKey    string
		Name      string
		Ip        string
		Type      int8
		Exp       int64
		CreatedAt string
		UpdatedAt string
		DeletedAt string
		StatusAt  int8
	}
	type args struct {
		sql   string
		value []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsAuthApplication := HsAuthApplication{
				Id:        tt.fields.Id,
				SecretKey: tt.fields.SecretKey,
				AppKey:    tt.fields.AppKey,
				Name:      tt.fields.Name,
				Ip:        tt.fields.Ip,
				Type:      tt.fields.Type,
				Exp:       tt.fields.Exp,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
				StatusAt:  tt.fields.StatusAt,
			}
			got, err := hsAuthApplication.Exec(tt.args.sql, tt.args.value...)
			if (err != nil) != tt.wantErr {
				t.Errorf("HsAuthApplication.Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HsAuthApplication.Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}
