package ormstruct

type Auth_App struct {
	model Model
	HsAuthApplication
}

func (aa Auth_App) Select(sql string, limit, offset int, value ...interface{}) ([]interface{}, error) {
	return aa.model.Select(sql, limit, offset, value...)
}

func (aa Auth_App) FindByID(id int64) (interface{}, error) {
	return aa.model.FindByID(id)
}

func (aa Auth_App) Add() (int64, error) {
	return aa.model.Add()
}

func (aa Auth_App) AddBatch(obj []interface{}) error {
	return aa.model.AddBatch(obj)
}

func (aa Auth_App) Update() (int64, error) {
	return aa.model.Update()
}

func (aa Auth_App) UpdateBatch(obj []interface{}) error {
	return aa.model.UpdateBatch(obj)
}

func (aa Auth_App) Delete() (int64, error) {
	return aa.model.Delete()
}

func (aa Auth_App) DeleteBatch(obj []interface{}) error {
	return aa.model.DeleteBatch(obj)
}

func (aa Auth_App) Exec(sql string, value ...interface{}) (int64, error) {
	return aa.model.Exec(sql, value...)
}
func (aa Auth_App) GetSql() (string, []interface{}) {
	return aa.model.GetSql()
}
func (aa Auth_App) SetDBConn(db, str string) {
	aa.model.SetDBConn(db, str)
}

func New() Auth_App {
	ap := NewHsAuthApplication()
	aa := HsAuthApplication{}
	return Auth_App{&ap, aa}
}
