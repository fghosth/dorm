package lexer

const (
	STRUCT_TMP = `package {{{packageName}}}

   
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
		 {{{primaryKey}}}
	 ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='';
	`
	COCKROACH_SCRIPT_TMP = `
	 CREATE TABLE IF NOT EXISTS {{{tableName}}} (
		 {{#each field}}
 		{{{this}}}
 	   {{/each}}
		{{{pk}}}
		{{{uq}}}
		  {{#each ik}}
		{{{this}}}
		  {{/each}}
	 );
	`
	COCKROACH_INSERT_TMP = `
		{{#each field}}
	 {{{this}}}
		{{/each}}
	`
)
