package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-oci8"
	"os"
	// "strings"
)

func getDSN() string {
	var dsn string
	if len(os.Args) > 1 {
		dsn = os.Args[1]
		if dsn != "" {
			return dsn
		}
	}
	dsn = os.Getenv("GO_OCI8_CONNECT_STRING")
	if dsn != "" {
		return dsn
	}
	fmt.Fprintln(os.Stderr, `Please specifiy connection parameter in GO_OCI8_CONNECT_STRING environment variable,
or as the first argument! (The format is user/name@host:port/sid)`)
	return "vcowner/dbora999@127.0.0.1:1521/OPAT"
}

func main() {
	os.Setenv("NLS_LANG", "SIMPLIFIED CHINESE_CHINA.AL32UTF8")

	db, err := sql.Open("oci8", getDSN())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT OID, AD_TYPE_NAME FROM (SELECT * FROM \"ADTYPE_JA\") WHERE rownum <= 50")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var oid string
		var ad_type_name string
		rows.Scan(&oid, &ad_type_name)
		fmt.Println(oid, ad_type_name)
	}
}
