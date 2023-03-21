package models

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"bitbucket.org/isbtotogroup/sdsb4d-backend/configs"
	"bitbucket.org/isbtotogroup/sdsb4d-backend/db"
	"bitbucket.org/isbtotogroup/sdsb4d-backend/entities"
	"bitbucket.org/isbtotogroup/sdsb4d-backend/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/nleeper/goment"
)

func Fetch_japanHome(tipe string) (helpers.Response, error) {
	var obj entities.Model_japan
	var arraobj []entities.Model_japan
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()
	sql_select := ""
	if tipe == "day" {
		sql_select = `SELECT 
			idjapan_japanday , 
			to_char(COALESCE(datejapan_japanday,now()), 'YYYY-MM-DD '), 
			prize1_japanday , prize2_japanday, prize3_japanday,   
			create_japanday, to_char(COALESCE(createdate_japanday,now()), 'YYYY-MM-DD HH24:MI:SS'),  
			update_japanday, to_char(COALESCE(updatedate_japanday,now()), 'YYYY-MM-DD HH24:MI:SS')   
			FROM ` + configs.DB_tbl_trx_japan_day + ` 
			ORDER BY datejapan_japanday DESC LIMIT 365 
		`
	} else {
		sql_select = `SELECT 
			idjapan_japannight , 
			to_char(COALESCE(datejapan_japannight,now()), 'YYYY-MM-DD '), 
			prize1_japannight , prize2_japannight, prize3_japannight,   
			create_japannight, to_char(COALESCE(createdate_japannight,now()), 'YYYY-MM-DD HH24:MI:SS'),  
			update_japannight, to_char(COALESCE(updatedate_japannight,now()), 'YYYY-MM-DD HH24:MI:SS')   
			FROM ` + configs.DB_tbl_trx_japan_night + ` 
			ORDER BY datejapan_japannight DESC LIMIT 365 
		`
	}

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idjapan_japan_db                                                           int
			datejapan_japan_db, prize1_japan_db, prize2_japan_db, prize3_japan_db      string
			create_japan_db, createdate_japan_db, update_japan_db, updatedate_japan_db string
		)

		err = row.Scan(
			&idjapan_japan_db, &datejapan_japan_db, &prize1_japan_db, &prize2_japan_db, &prize3_japan_db,
			&create_japan_db, &createdate_japan_db, &update_japan_db, &updatedate_japan_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		if create_japan_db != "" {
			create = create_japan_db + ", " + createdate_japan_db
		}
		if update_japan_db != "" {
			update = update_japan_db + ", " + updatedate_japan_db
		}
		obj.Japan_id = idjapan_japan_db
		obj.Japan_date = datejapan_japan_db
		obj.Japan_prize1 = prize1_japan_db
		obj.Japan_prize2 = prize2_japan_db
		obj.Japan_prize3 = prize3_japan_db
		obj.Japan_create = create
		obj.Japan_update = update
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
func Save_japanHome(admin, tanggal, tipe, sData string, idrecord int) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	if sData == "New" {
		if tipe == "day" {
			flag = CheckDB(configs.DB_tbl_trx_japan_day, "datejapan_japanday", tanggal)
			if !flag {
				sql_insert := `
				insert into
				` + configs.DB_tbl_trx_japan_day + ` (
					idjapan_japanday  , datejapan_japanday, create_japanday, createdate_japanday
				) values (
					$1 ,$2, $3, $4 
				)
			`
				field_column := configs.DB_tbl_trx_japan_day + tglnow.Format("YYYY")
				idrecord_counter := Get_counter(field_column)
				flag_insert, msg_insert := Exec_SQL(sql_insert, configs.DB_tbl_trx_japan_day, "INSERT",
					tglnow.Format("YY")+strconv.Itoa(idrecord_counter), tanggal,
					admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

				if !flag_insert {
					fmt.Println(msg_insert)
				} else {
					msg = "Success"
				}
			} else {
				msg = "Duplicate Entry"
			}
		} else {
			flag = CheckDB(configs.DB_tbl_trx_japan_night, "datejapan_japannight", tanggal)
			if !flag {
				sql_insert := `
				insert into
				` + configs.DB_tbl_trx_japan_night + ` (
					idjapan_japannight  , datejapan_japannight, create_japannight, createdate_japannight
				) values (
					$1 ,$2, $3, $4 
				)
			`
				field_column := configs.DB_tbl_trx_japan_night + tglnow.Format("YYYY")
				idrecord_counter := Get_counter(field_column)
				flag_insert, msg_insert := Exec_SQL(sql_insert, configs.DB_tbl_trx_japan_night, "INSERT",
					tglnow.Format("YY")+strconv.Itoa(idrecord_counter), tanggal,
					admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

				if !flag_insert {
					fmt.Println(msg_insert)
				} else {
					msg = "Success"
				}
			} else {
				msg = "Duplicate Entry"
			}
		}

	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}
func Save_japanGenerator(admin, field, prize, tipe, sData string, idrecord int) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	statuspasaran := "OFFLINE"
	tglnow, _ := goment.New()
	render_page := time.Now()

	if sData == "Edit" {
		prize_time := ""
		if tipe == "day" {
			prize_time = " 09:55:00"
			tglskrg := _getJapan(idrecord, "day")
			tglnow2, _ := goment.New(tglskrg)
			tglhariini := tglnow.Format("YYYY-MM-DD HH:mm:ss")
			tglpasaranend := tglnow2.Format("YYYY-MM-DD") + prize_time
			fmt.Println(tglhariini)
			fmt.Println(tglpasaranend)
			if tglhariini >= tglpasaranend {
				statuspasaran = "ONLINE"
			} else {
				if tglhariini >= tglpasaranend {
					statuspasaran = "ONLINE"
				} else {
					msg = "Offline"
				}
			}
			// statuspasaran = "OFFLINE"
			if statuspasaran == "ONLINE" {
				sql_update := `
				UPDATE
				` + configs.DB_tbl_trx_japan_day + `
				SET ` + field + ` =$1,
				update_japanday=$2, updatedate_japanday=$3
				WHERE idjapan_japanday=$4
			`
				flag_update, msg_update := Exec_SQL(sql_update, configs.DB_tbl_trx_japan_day, "UPDATE",
					prize, admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord)

				if !flag_update {
					fmt.Println(msg_update)
				} else {
					msg = "Success"
				}
			}
		} else {
			prize_time = " 17:55:00"
			tglskrg := _getJapan(idrecord, "night")
			tglnow2, _ := goment.New(tglskrg)
			tglhariini := tglnow.Format("YYYY-MM-DD HH:mm:ss")
			tglpasaranend := tglnow2.Format("YYYY-MM-DD") + prize_time
			fmt.Println(tglhariini)
			fmt.Println(tglpasaranend)
			if tglhariini >= tglpasaranend {
				statuspasaran = "ONLINE"
			} else {
				if tglhariini >= tglpasaranend {
					statuspasaran = "ONLINE"
				} else {
					msg = "Offline"
				}
			}
			// statuspasaran = "OFFLINE"
			if statuspasaran == "ONLINE" {
				sql_update := `
					UPDATE
					` + configs.DB_tbl_trx_japan_night + `
					SET ` + field + ` =$1,
					update_japannight=$2, updatedate_japannight=$3
					WHERE idjapan_japannight=$4
			`
				flag_update, msg_update := Exec_SQL(sql_update, configs.DB_tbl_trx_japan_night, "UPDATE",
					prize, admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord)

				if !flag_update {
					fmt.Println(msg_update)
				} else {
					msg = "Success"
				}
			}
		}

	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}
func Save_Generatorjapan(admin, tipe string) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	go func() {
		for i := 0; i <= 808; i++ {
			tglnow2, _ := goment.New("2021-01-01")
			tanggal := tglnow2.Add(i, "days").Format("YYYY-MM-DD")
			if tipe == "day" { //day
				flag = CheckDB(configs.DB_tbl_trx_japan_day, "datejapan_japanday", tanggal)
				if !flag {
					sql_insert := `
					insert into
					` + configs.DB_tbl_trx_japan_day + ` (
						idjapan_japanday, datejapan_japanday, 
						prize1_japanday, prize2_japanday, prize3_japanday, 
						create_japanday, createdate_japanday
					) values (
						$1, $2, 
						$3, $4, $5, 
						$6, $7  
					)
				`
					prize_1 := helpers.GenerateNumber(4)
					prize_2 := helpers.GenerateNumber(4)
					prize_3 := helpers.GenerateNumber(4)
					field_column := configs.DB_tbl_trx_japan_day + tglnow.Format("YYYY")
					idrecord_counter := Get_counter(field_column)
					flag_insert, msg_insert := Exec_SQL(sql_insert, configs.DB_tbl_trx_japan_day, "UPDATE",
						tglnow.Format("YY")+strconv.Itoa(idrecord_counter),
						tanggal, prize_1, prize_2, prize_3,
						admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

					if !flag_insert {
						fmt.Println(msg_insert)
					} else {
						msg = "Success"
					}
				} else {
					break
				}
			} else { //night
				flag = CheckDB(configs.DB_tbl_trx_japan_night, "datejapan_japannight", tanggal)
				if !flag {
					sql_insert := `
					insert into
					` + configs.DB_tbl_trx_japan_night + ` (
						idjapan_japannight, datejapan_japannight, 
						prize1_japannight, prize2_japannight, prize3_japannight, 
						create_japannight, createdate_japannight
					) values (
						$1, $2, 
						$3, $4, $5, 
						$6, $7  
					)
				`
					prize_1 := helpers.GenerateNumber(4)
					prize_2 := helpers.GenerateNumber(4)
					prize_3 := helpers.GenerateNumber(4)
					field_column := configs.DB_tbl_trx_japan_night + tglnow.Format("YYYY")
					idrecord_counter := Get_counter(field_column)
					flag_insert, msg_insert := Exec_SQL(sql_insert, configs.DB_tbl_trx_japan_night, "UPDATE",
						tglnow.Format("YY")+strconv.Itoa(idrecord_counter),
						tanggal, prize_1, prize_2, prize_3,
						admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

					if !flag_insert {
						fmt.Println(msg_insert)
					} else {
						msg = "Success"
					}
				} else {
					break
				}
			}

		}
	}()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}
func _getJapan(idrecord int, tipe string) string {
	con := db.CreateCon()
	ctx := context.Background()
	var date string
	date = ""
	sql_select := ""
	if tipe == "day" {
		sql_select = `SELECT 
					datejapan_japanday
					FROM ` + configs.DB_tbl_trx_japan_day + ` 
					WHERE idjapan_japanday = $1 
				`
	} else {
		sql_select = `SELECT 
			datejapan_japannight
				FROM ` + configs.DB_tbl_trx_japan_night + ` 
				WHERE idjapan_japannight = $1 
			`
	}

	row := con.QueryRowContext(ctx, sql_select, idrecord)
	switch e := row.Scan(&date); e {
	case sql.ErrNoRows:
		fmt.Println("CHECKDBTHREEFIELD - No rows were returned!")
	case nil:
	default:

	}
	return date
}
