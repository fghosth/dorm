package dorm

import (
	"strconv"

	"github.com/aymerick/raymond"
	"jvole.com/createProject/lexer"
	"jvole.com/createProject/util"
)

type dorm struct{}

var (
	ut = new(util.Dstring)
	sl = new(lexer.StructLexer)
)

func CreateDorm(pkname, sqlDriver, structStr string) string {
	var filestring string
	header := CreateHeader(pkname, sqlDriver) + "\n"
	field := CreateField(structStr) + "\n"
	function := CreateFunction(structStr) + "\n"
	selec := CreateSelect(structStr) + "\n"
	findByID := CreateFindByID(structStr) + "\n"
	add := CreateAdd(structStr) + "\n"
	addbatch := CreateAddBatch(structStr) + "\n"
	update := CreateUpdate(structStr) + "\n"
	updatebatch := CreateUpdateBatch(structStr) + "\n"
	delete := CreateDelete(structStr) + "\n"
	deletebatch := CreateDeleteBatch(structStr) + "\n"
	exec := CreateExec(structStr) + "\n"

	filestring = header + field + structStr + "\n" + function + selec + findByID + add + addbatch + update + updatebatch + delete + deletebatch + exec
	return filestring
}
func CreateModel(pkname string) string {
	ctx := map[string]interface{}{
		"pkname": pkname,
	}
	str, err := raymond.Render(MODEL_TPL, ctx)
	ut.Checkerr(err)
	return str
}
func CreateField(structStr string) string {
	obj := sl.StructName(structStr)

	ctx := map[string]interface{}{
		"obj": obj,
	}
	str, err := raymond.Render(Field_TPL, ctx)
	ut.Checkerr(err)
	return str
}

func CreateHeader(pkname, sqlDriver string) string {
	var imp = [...]string{`"database/sql"`, `"log"`, `"strconv"`}
	var dbDriver string
	switch sqlDriver {
	case "mysql":
		dbDriver = "github.com/go-sql-driver/mysql"
	case "mariadb":
		dbDriver = "github.com/go-sql-driver/mysql"
	case "cockroachDB":
		dbDriver = "github.com/lib/pq"
	}

	ctx := map[string]interface{}{
		"field":     imp,
		"pkname":    pkname,
		"sqlDriver": dbDriver,
	}
	str, err := raymond.Render(Header_TPL, ctx)
	ut.Checkerr(err)
	return str
}

func CreateFunction(structStr string) string {
	obj := sl.StructName(structStr)
	objvar, err := ut.FUPer(obj)
	ut.Checkerr(err)

	ctx := map[string]interface{}{
		"objvar": objvar,
		"obj":    obj,
	}
	str, err := raymond.Render(Function_TPL, ctx)
	ut.Checkerr(err)
	return str
}

func CreateExec(structStr string) string {
	obj := sl.StructName(structStr)
	objvar, err := ut.FUPer(obj)
	ut.Checkerr(err)

	ctx := map[string]interface{}{
		"objvar": objvar,
		"obj":    obj,
	}
	str, err := raymond.Render(Exec_TPL, ctx)
	ut.Checkerr(err)
	return str
}

func CreateDeleteBatch(structStr string) string {
	obj := sl.StructName(structStr)
	objvar, err := ut.FUPer(obj)
	ut.Checkerr(err)
	tableName := ut.CalToUnder(obj)
	field := sl.FieldName(structStr)
	var fields, structField, sqlField string
	if len(field) > 0 {
		structField = field[0]["field"]
		sqlField = ut.CalToUnder(field[0]["field"])
	}
	ctx := map[string]interface{}{
		"objvar":      objvar,
		"obj":         obj,
		"fields":      fields,
		"tableName":   tableName,
		"sqlField":    sqlField,
		"structField": structField,
	}
	str, err := raymond.Render(DeleteBatch_TPL, ctx)
	ut.Checkerr(err)
	return str
}

func CreateDelete(structStr string) string {
	obj := sl.StructName(structStr)
	objvar, err := ut.FUPer(obj)
	ut.Checkerr(err)
	tableName := ut.CalToUnder(obj)
	field := sl.FieldName(structStr)
	var fields, structField, sqlField string
	if len(field) > 0 {
		structField = field[0]["field"]
		sqlField = ut.CalToUnder(field[0]["field"])
	}
	ctx := map[string]interface{}{
		"objvar":      objvar,
		"obj":         obj,
		"fields":      fields,
		"tableName":   tableName,
		"sqlField":    sqlField,
		"structField": structField,
	}
	str, err := raymond.Render(Delete_TPL, ctx)
	ut.Checkerr(err)
	return str
}

func CreateUpdateBatch(structStr string) string {
	obj := sl.StructName(structStr)
	objvar, err := ut.FUPer(obj)
	ut.Checkerr(err)
	tableName := ut.CalToUnder(obj)
	field := sl.FieldName(structStr)
	var fields, sqlField string
	if len(field) > 0 {
		sqlField = ut.CalToUnder(field[0]["field"])
	}
	length := len(field)
	values := make([]string, len(field))
	for i, v := range field {
		if i > 0 { //去除id
			values[i-1] = "args[" + strconv.Itoa(i-1) + "]=v." + v["field"]
			if fields == "" {
				fields = ut.CalToUnder(v["field"]) + "=?"
			} else {
				fields = fields + "," + ut.CalToUnder(v["field"]) + "=?"
			}
		} else {
			values[length-1] = "args[" + strconv.Itoa(length-1) + "]=v." + v["field"]
		}

	}
	ctx := map[string]interface{}{
		"objvar":    objvar,
		"obj":       obj,
		"fields":    fields,
		"tableName": tableName,
		"field":     values,
		"sqlField":  sqlField,
		"len":       length,
	}
	str, err := raymond.Render(UpdateBatch_TPL, ctx)
	ut.Checkerr(err)
	return str
}

func CreateUpdate(structStr string) string {
	obj := sl.StructName(structStr)
	objvar, err := ut.FUPer(obj)
	ut.Checkerr(err)
	tableName := ut.CalToUnder(obj)
	field := sl.FieldName(structStr)
	var fields, sqlField string
	if len(field) > 0 {
		sqlField = ut.CalToUnder(field[0]["field"])
	}
	length := len(field)
	values := make([]string, len(field))
	for i, v := range field {
		if i > 0 { //去除id
			values[i-1] = "args[" + strconv.Itoa(i-1) + "]=&" + objvar + "." + v["field"]
			if fields == "" {
				fields = ut.CalToUnder(v["field"]) + "=?"
			} else {
				fields = fields + "," + ut.CalToUnder(v["field"]) + "=?"
			}
		} else {
			values[length-1] = "args[" + strconv.Itoa(length-1) + "]=&" + objvar + "." + v["field"]
		}

	}

	ctx := map[string]interface{}{
		"objvar":    objvar,
		"obj":       obj,
		"fields":    fields,
		"tableName": tableName,
		"field":     values,
		"sqlField":  sqlField,
		"len":       length,
	}
	str, err := raymond.Render(Update_TPL, ctx)
	ut.Checkerr(err)
	return str
}

func CreateAddBatch(structStr string) string {
	obj := sl.StructName(structStr)
	objvar, err := ut.FUPer(obj)
	ut.Checkerr(err)
	tableName := ut.CalToUnder(obj)
	field := sl.FieldName(structStr)
	var fields, parms string
	length := len(field) - 1
	values := make([]string, len(field))
	for i, v := range field {
		if i > 0 { //去除id
			values[i-1] = "args[" + strconv.Itoa(i-1) + "]=v." + v["field"]
			if fields == "" {
				parms = "?"
				fields = ut.CalToUnder(v["field"])
			} else {
				parms = parms + ",?"
				fields = fields + "," + ut.CalToUnder(v["field"])
			}
		}

	}
	ctx := map[string]interface{}{
		"objvar":    objvar,
		"obj":       obj,
		"fields":    fields,
		"tableName": tableName,
		"field":     values,
		"parms":     parms,
		"len":       length,
	}
	str, err := raymond.Render(AddBatch_TPL, ctx)
	ut.Checkerr(err)
	return str
}

func CreateAdd(structStr string) string {
	obj := sl.StructName(structStr)
	objvar, err := ut.FUPer(obj)
	ut.Checkerr(err)
	tableName := ut.CalToUnder(obj)
	field := sl.FieldName(structStr)
	var fields, parms string
	length := len(field) - 1
	values := make([]string, len(field))
	for i, v := range field {
		if i > 0 { //去除id
			values[i-1] = "args[" + strconv.Itoa(i-1) + "]=&" + objvar + "." + v["field"]
			if fields == "" {
				parms = "?"
				fields = ut.CalToUnder(v["field"])
			} else {
				parms = parms + ",?"
				fields = fields + "," + ut.CalToUnder(v["field"])
			}
		}

	}
	ctx := map[string]interface{}{
		"objvar":    objvar,
		"obj":       obj,
		"fields":    fields,
		"tableName": tableName,
		"field":     values,
		"parms":     parms,
		"len":       length,
	}
	str, err := raymond.Render(Add_TPL, ctx)
	ut.Checkerr(err)
	return str
}

func CreateFindByID(structStr string) string {
	obj := sl.StructName(structStr)
	objvar, err := ut.FUPer(obj)
	ut.Checkerr(err)
	tableName := ut.CalToUnder(obj)
	field := sl.FieldName(structStr)
	var fields, sqlField string
	if len(field) > 0 {
		sqlField = ut.CalToUnder(field[0]["field"])
	}

	values := make([]string, len(field))
	for i, v := range field {
		values[i] = "values[" + strconv.Itoa(i) + "]=&" + objvar + "." + v["field"]
		if fields == "" {
			fields = ut.CalToUnder(v["field"])
		} else {
			fields = fields + "," + ut.CalToUnder(v["field"])
		}
	}
	ctx := map[string]interface{}{
		"objvar":    objvar,
		"obj":       obj,
		"fields":    fields,
		"tableName": tableName,
		"field":     values,
		"sqlField":  sqlField,
	}
	str, err := raymond.Render(FindByID_TPL, ctx)
	ut.Checkerr(err)
	return str
}

func CreateSelect(structStr string) string {
	obj := sl.StructName(structStr)
	objvar, err := ut.FUPer(obj)
	ut.Checkerr(err)
	tableName := ut.CalToUnder(obj)
	// pp.Println(objvar + "===" + tableName)

	field := sl.FieldName(structStr)
	fields := ""
	values := make([]string, len(field))
	for i, v := range field {
		values[i] = "values[" + strconv.Itoa(i) + "]=&" + objvar + "." + v["field"]
		if fields == "" {
			fields = ut.CalToUnder(v["field"])
		} else {
			fields = fields + "," + ut.CalToUnder(v["field"])
		}
	}
	// pp.Println(field)
	// pp.Println(fields)
	// pp.Println(values)
	ctx := map[string]interface{}{
		"objvar":    objvar,
		"obj":       obj,
		"fields":    fields,
		"tableName": tableName,
		"field":     values,
	}
	selectString, err := raymond.Render(Select_TPL, ctx)
	ut.Checkerr(err)
	return selectString
	// fmt.Println(selectString)
}
