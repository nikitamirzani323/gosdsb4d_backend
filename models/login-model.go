package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"bitbucket.org/isbtotogroup/sdsb4d-backend/configs"
	"bitbucket.org/isbtotogroup/sdsb4d-backend/db"
	"bitbucket.org/isbtotogroup/sdsb4d-backend/helpers"
	"github.com/nleeper/goment"
)

func Login_Model(username, password, ipaddress, timezone string) (bool, string, error) {
	con := db.CreateCon()
	ctx := context.Background()
	flag := false
	tglnow, _ := goment.New()
	var passwordDB, idadminDB string
	sql_select := `
			SELECT
			password, idadmin    
			FROM ` + configs.DB_tbl_admin + ` 
			WHERE username  = $1 
			AND statuslogin = 'Y' 
		`
	row := con.QueryRowContext(ctx, sql_select, username)
	switch e := row.Scan(&passwordDB, &idadminDB); e {
	case sql.ErrNoRows:
		return false, "", errors.New("Username and Password Not Found")
	case nil:
		flag = true
	default:
		return false, "", errors.New("Username and Password Not Found")
	}

	hashpass := helpers.HashPasswordMD5(password)
	fmt.Println("Password : " + hashpass)
	fmt.Println("Hash : " + passwordDB)
	if hashpass != passwordDB {
		return false, "", nil
	}

	if flag {
		sql_update := `
			UPDATE ` + configs.DB_tbl_admin + ` 
			SET lastlogin=$1, ipaddress=$2 , timezone=$3, 
			updateadmin=$4,  updatedateadmin=$5   
			WHERE username= $6  
			AND statuslogin = 'Y' 
		`
		flag_update, msg_update := Exec_SQL(sql_update, configs.DB_tbl_admin, "UPDATE",
			tglnow.Format("YYYY-MM-DD HH:mm:ss"),
			ipaddress, timezone, username,
			tglnow.Format("YYYY-MM-DD HH:mm:ss"), username)

		if flag_update {
			flag = true
			fmt.Println(msg_update)
		} else {
			fmt.Println(msg_update)
		}

	}

	return true, idadminDB, nil
}
