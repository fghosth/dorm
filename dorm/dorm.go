package dorm

import (
	"strconv"

	"github.com/aymerick/raymond"
	"jvole.com/createProject/lexer"
	"jvole.com/createProject/util"
)

type dorm struct{}

var (
	ut       = new(util.Dstring)
	sl       = new(lexer.StructLexer)
	baseName = "base"
)

func CreateDorm(pkname, structStr string) string {
	var filestring string
	header := CreateHeader(pkname) + "\n"
	field := CreateField(structStr) + "\n"
	function := CreateFunction(structStr) + "\n"
	getArgsStr := CreateGetArgsStrFun(structStr) + "\n"
	selec := CreateSelect(structStr) + "\n"
	findByID := CreateFindByID(structStr) + "\n"
	add := CreateAdd(structStr) + "\n"
	addbatch := CreateAddBatch(structStr) + "\n"
	update := CreateUpdate(structStr) + "\n"
	updatebatch := CreateUpdateBatch(structStr) + "\n"
	delete := CreateDelete(structStr) + "\n"
	deletebatch := CreateDeleteBatch(structStr) + "\n"
	exec := CreateExec(structStr) + "\n"
	softDel := CreateSoftDeleteFun(structStr) + "\n"

	filestring = header + field + structStr + "\n" + function + getArgsStr + selec + findByID + add + addbatch + update + updatebatch + softDel + delete + deletebatch + exec
	return filestring
}

/*
location当前包的名字如"jvole.com/createProject/ormstruct/base" 就是『jvole.com/createProject/』
pkname 包名 如 ormstruct
*/
func CreateDAO(location, pkname, structStr string) string {
	obj := sl.StructName(structStr)
	field := sl.FieldName(structStr)
	objvar, err := ut.FUPer(obj)
	fields := make([]string, len(field))
	for i, v := range field {
		fields[i] = objvar + "." + v["field"] + " = dao." + v["field"]
	}
	ctx := map[string]interface{}{
		"obj":         obj,
		"objvar":      objvar,
		"pkname":      pkname,
		"field":       fields,
		"modelImport": ut.CheckAndAdd(location, "/") + ut.CheckAndAdd(pkname, "/") + baseName,
	}
	str, err := raymond.Render(DAO_TPL, ctx)
	ut.Checkerr(err)
	return str
}
func CreateSoftDeleteFun(structStr string) string {
	obj := sl.StructName(structStr)
	objvar, err := ut.FUPer(obj)
	ctx := map[string]interface{}{
		"obj":    obj,
		"objvar": objvar,
	}
	str, err := raymond.Render(SDEL_TPL, ctx)
	ut.Checkerr(err)
	return str
}
func CreateGetArgsStrFun(structStr string) string {
	obj := sl.StructName(structStr)
	field := sl.FieldName(structStr)
	var fields, pqfields, sqlField string
	if len(field) > 0 {
		sqlField = ut.CalToUnder(field[0]["field"])
	}
	for i, v := range field {
		if i > 0 { //去除id

			if fields == "" {
				fields = ut.CalToUnder(v["field"]) + "=?"
				pqfields = ut.CalToUnder(v["field"]) + "=$" + strconv.Itoa(i)
			} else {
				fields = fields + "," + ut.CalToUnder(v["field"]) + "=?"
				pqfields = pqfields + "," + ut.CalToUnder(v["field"]) + "=$" + strconv.Itoa(i)
			}
		}

	}
	fields = fields + " WHERE \"  + SDELFLAG + \"=0 and " + sqlField + "=?"
	pqfields = pqfields + " WHERE \"  + SDELFLAG + \"=0 and " + sqlField + "=$" + strconv.Itoa(len(field))
	ctx := map[string]interface{}{
		"obj":              obj,
		"mysqlField":       fields,
		"mariadbField":     fields,
		"cockroachDBField": pqfields,
		"postgresqlField":  pqfields,
	}
	str, err := raymond.Render(GetArgsStrFun_TPL, ctx)
	ut.Checkerr(err)
	return str
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

func CreateHeader(pkname string) string {
	var imp = [...]string{`"database/sql"`, `"log"`, `"strconv"`}

	ctx := map[string]interface{}{
		"field":  imp,
		"pkname": pkname,
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

			values[i-1] = "args[" + strconv.Itoa(i-1) + "] = v." + v["field"]

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

			values[i-1] = "args[" + strconv.Itoa(i-1) + "] = " + objvar + "." + v["field"]

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
	var fields, parms, cockroachParms string
	length := len(field) - 1
	values := make([]string, len(field))
	for i, v := range field {
		if i > 0 { //去除id
			values[i-1] = "args[" + strconv.Itoa(i-1) + "]=v." + v["field"]
			if fields == "" {
				parms = "?"
				cockroachParms = "$" + strconv.Itoa(i)
				fields = ut.CalToUnder(v["field"])
			} else {
				parms = parms + ",?"
				cockroachParms = cockroachParms + ",$" + strconv.Itoa(i)
				fields = fields + "," + ut.CalToUnder(v["field"])
			}
		}

	}
	ctx := map[string]interface{}{
		"objvar":         objvar,
		"obj":            obj,
		"fields":         fields, //sql字段
		"tableName":      tableName,
		"field":          values,
		"mysqlparms":     parms,
		"cockroachparms": cockroachParms,
		"len":            length,
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

		values[i] = "values[" + strconv.Itoa(i) + "] = &" + objvar + "." + v["field"]

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

		values[i] = "values[" + strconv.Itoa(i) + "] = &" + objvar + "." + v["field"]

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
