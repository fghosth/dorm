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
