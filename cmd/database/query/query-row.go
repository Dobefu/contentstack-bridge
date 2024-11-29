package query

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/Dobefu/csb/cmd/database"
	"github.com/Dobefu/csb/cmd/database/structs"
	"github.com/Dobefu/csb/cmd/database/utils"
)

func QueryRow(table string, fields []string, where []structs.QueryWhere) *sql.Row {
	fieldsString := strings.Join(fields, ", ")

	switch os.Getenv("DB_TYPE") {
	case "mysql":
		return queryRowMysql(table, fieldsString, where)
	case "sqlite3":
		return queryRowSqlite3(table, fieldsString, where)
	default:
		return nil
	}
}

func queryRowMysql(table string, fields string, where []structs.QueryWhere) *sql.Row {
	sql := []string{fmt.Sprintf(
		"SELECT %s FROM %s",
		fields,
		table,
	)}

	var args []any

	if where != nil {
		whereString, newArgs := utils.ConstructWhere(where)

		sql = append(sql, whereString)
		args = append(args, newArgs...)
	}

	sql = append(sql, "LIMIT 1")
	return database.DB.QueryRow(strings.Join(sql, " "), args...)
}

func queryRowSqlite3(table string, fields string, where []structs.QueryWhere) *sql.Row {
	return queryRowMysql(table, fields, where)
}
