package lexer

const (
	STRUCT_TMP = `package {{{packageName}}}

   import (
   	"database/sql"
   	"fmt"
   	_ "github.com/go-sql-driver/mysql"
   )
   type {{{name}}} struct{
     {{#each field}}
 		{{{this}}}
 	   {{/each}}
   }
   `
	MYSQL_SCRIPT_TMP = `
	 CREATE TABLE ` + "`{{{tableName}}}`" + ` (
		 {{#each field}}
 		{{{this}}}
 	   {{/each}}
	 ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='';
	`
)
