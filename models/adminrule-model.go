package models

import (
	"context"
	"fmt"
	"time"

	"bitbucket.org/isbtotogroup/sdsb4d-backend/configs"
	"bitbucket.org/isbtotogroup/sdsb4d-backend/db"
	"bitbucket.org/isbtotogroup/sdsb4d-backend/entities"
	"bitbucket.org/isbtotogroup/sdsb4d-backend/helpers"
	"github.com/gofiber/fiber/v2"
)

func Fetch_adminruleHome() (helpers.Response, error) {
	var obj entities.Model_adminruleall
	var arraobj []entities.Model_adminruleall
	var res helpers.Response
	msg := "Error"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
			idadmin , ruleadmingroup 
			FROM ` + configs.DB_tbl_admingroup + ` 
			ORDER BY idadmin ASC  
		`

	row, err := con.QueryContext(ctx, sql_select)

	var no int = 0
	helpers.ErrorCheck(err)
	for row.Next() {
		no += 1
		var (
			idadmin_db, ruleadmingroup_db string
		)

		err = row.Scan(&idadmin_db, &ruleadmingroup_db)

		helpers.ErrorCheck(err)

		obj.Idadmin = idadmin_db
		obj.Ruleadmingroup = ruleadmingroup_db
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(start).String()

	return res, nil
}
func Save_adminrule(admin, idadmin, rule, sData string) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	render_page := time.Now()
	flag := false

	if sData == "New" {
		flag = CheckDB(configs.DB_tbl_admingroup, "idadmin ", idadmin)
		if !flag {
			sql_insert := `
				insert into
				` + configs.DB_tbl_admingroup + ` (
					idadmin 
				) values (
					$1 
				)
			`

			flag_insert, msg_insert := Exec_SQL(sql_insert, configs.DB_tbl_admingroup, "INSERT", idadmin)

			if !flag_insert {
				fmt.Println(msg_insert)
			} else {
				msg = "Success"
			}

		} else {
			msg = "Duplicate Entry"
		}
	} else {
		sql_update := `
			UPDATE 
			` + configs.DB_tbl_admingroup + `   
			SET ruleadmingroup =$1
			WHERE idadmin  =$2  
		`
		flag_update, msg_update := Exec_SQL(sql_update, configs.DB_tbl_admingroup, "UPDATE", rule, idadmin)

		if !flag_update {
			fmt.Println(msg_update)
		} else {
			msg = "Success"
		}

	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}
