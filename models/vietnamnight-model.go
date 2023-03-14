package models

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	"bitbucket.org/isbtotogroup/sdsb4d-backend/configs"
	"bitbucket.org/isbtotogroup/sdsb4d-backend/db"
	"bitbucket.org/isbtotogroup/sdsb4d-backend/entities"
	"bitbucket.org/isbtotogroup/sdsb4d-backend/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/nleeper/goment"
)

func Fetch_vietnamenightHome() (helpers.Response, error) {
	var obj entities.Model_vietnamenight
	var arraobj []entities.Model_vietnamenight
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
			idvietnamnight , 
			to_char(COALESCE(datevietnamnight,now()), 'YYYY-MM-DD '), 
			prize_1300 , prize_1700, prize_2000, prize_2200,  
			create_vietnamnight, to_char(COALESCE(createdate_vietnamnight,now()), 'YYYY-MM-DD HH24:MI:SS'),  
			update_vietnamnight, to_char(COALESCE(updatedate_vietnamnight,now()), 'YYYY-MM-DD HH24:MI:SS')   
			FROM ` + configs.DB_tbl_trx_vietnam_night + ` 
			ORDER BY datevietnamnight DESC LIMIT 365 
		`

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idvietnamnight_db                                                                                      int
			datevietnamnight_db, prize_1300_db, prize_1700_db, prize_2000_db, prize_2200_db                        string
			create_vietnamnight_db, createdate_vietnamnight_db, update_vietnamnight_db, updatedate_vietnamnight_db string
		)

		err = row.Scan(
			&idvietnamnight_db, &datevietnamnight_db, &prize_1300_db, &prize_1700_db, &prize_2000_db, &prize_2200_db,
			&create_vietnamnight_db, &createdate_vietnamnight_db, &update_vietnamnight_db, &updatedate_vietnamnight_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		if create_vietnamnight_db != "" {
			create = create_vietnamnight_db + ", " + createdate_vietnamnight_db
		}
		if update_vietnamnight_db != "" {
			update = update_vietnamnight_db + ", " + updatedate_vietnamnight_db
		}
		obj.Vietnamnight_id = idvietnamnight_db
		obj.Vietnamnight_date = datevietnamnight_db
		obj.Vietnamnight_prize1_1300 = prize_1300_db
		obj.Vietnamnight_prize1_1700 = prize_1700_db
		obj.Vietnamnight_prize1_2000 = prize_2000_db
		obj.Vietnamnight_prize1_2200 = prize_2200_db
		obj.Vietnamnight_create = create
		obj.Vietnamnight_update = update
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
func Save_vietnamenightHome(admin, tanggal, sData string, idrecord int) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	if sData == "New" {
		flag = CheckDB(configs.DB_tbl_trx_vietnam_night, "datevietnamnight", tanggal)
		if !flag {
			sql_insert := `
				insert into
				` + configs.DB_tbl_trx_vietnam_night + ` (
					idvietnamnight  , datevietnamnight, create_vietnamnight, createdate_vietnamnight
				) values (
					$1 ,$2, $3, $4 
				)
			`
			field_column := configs.DB_tbl_trx_vietnam_night + tglnow.Format("YYYY")
			idrecord_counter := Get_counter(field_column)
			flag_insert, msg_insert := Exec_SQL(sql_insert, configs.DB_tbl_trx_sdsb4d_day, "INSERT",
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

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}
func Save_vietnamenightGenerator(admin, field, prize, sData string, idrecord int) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	statuspasaran := "OFFLINE"
	tglnow, _ := goment.New()
	render_page := time.Now()

	if sData == "Edit" {
		prize_time := ""
		switch field {
		case "prize_1300":
			prize_time = " 12:55:00"
		case "prize_1700":
			prize_time = " 16:55:00"
		case "prize_2000":
			prize_time = " 19:55:00"
		case "prize_2200":
			prize_time = " 20:55:00"
		}
		tglskrg := _getVietnamNight(idrecord)
		tglnow2, _ := goment.New(tglskrg)
		tglhariini := tglnow.Format("YYYY-MM-DD HH:mm:ss")
		tglstart := tglnow2.Format("YYYY-MM-DD HH:mm:ss")
		tglskrgend := tglnow.Format("YYYY-MM-DD") + prize_time
		log.Println(tglhariini)
		log.Println(tglstart)
		log.Println(tglskrgend)
		if tglhariini >= tglskrgend {
			log.Println("level1")
			statuspasaran = "ONLINE"
		} else {
			log.Println("level2")
			if tglhariini >= tglskrgend {
				statuspasaran = "ONLINE"
			} else {
				msg = "Offline"
			}
		}
		// statuspasaran = "OFFLINE"
		if statuspasaran == "ONLINE" {
			sql_update := `
				UPDATE
				` + configs.DB_tbl_trx_vietnam_night + `
				SET ` + field + ` =$1,
				update_vietnamnight=$2, updatedate_vietnamnight=$3
				WHERE idvietnamnight =$4
			`
			flag_update, msg_update := Exec_SQL(sql_update, configs.DB_tbl_trx_vietnam_night, "UPDATE",
				prize, admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord)

			if !flag_update {
				fmt.Println(msg_update)
			} else {
				msg = "Success"
			}
		}
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}
func Save_Generatorvietnamenight(admin string) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	go func() {
		for i := 0; i <= 799; i++ {
			tglnow2, _ := goment.New("2021-01-01")
			tanggal := tglnow2.Add(i, "days").Format("YYYY-MM-DD")
			flag = CheckDB(configs.DB_tbl_trx_vietnam_night, "datevietnamnight", tanggal)
			if !flag {
				sql_insert := `
					insert into
					` + configs.DB_tbl_trx_vietnam_night + ` (
						idvietnamnight, datevietnamnight, 
						prize_1300, prize_1700, prize_2000, prize_2200, 
						create_vietnamnight, createdate_vietnamnight
					) values (
						$1, $2, 
						$3, $4, $5, $6,
						$7, $8 
					)
				`
				prize_1_1300 := helpers.GenerateNumber(4)
				prize_1_1700 := helpers.GenerateNumber(4)
				prize_1_2000 := helpers.GenerateNumber(4)
				prize_1_2200 := helpers.GenerateNumber(4)
				field_column := configs.DB_tbl_trx_vietnam_night + tglnow.Format("YYYY")
				idrecord_counter := Get_counter(field_column)
				flag_insert, msg_insert := Exec_SQL(sql_insert, configs.DB_tbl_trx_vietnam_night, "UPDATE",
					tglnow.Format("YY")+strconv.Itoa(idrecord_counter),
					tanggal,
					prize_1_1300,
					prize_1_1700,
					prize_1_2000,
					prize_1_2200,
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
	}()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}
func _getVietnamNight(idrecord int) string {
	con := db.CreateCon()
	ctx := context.Background()
	var date string
	date = ""
	sql_select := `SELECT 
					datevietnamnight
					FROM ` + configs.DB_tbl_trx_vietnam_night + ` 
					WHERE idvietnamnight = $1 
				`

	row := con.QueryRowContext(ctx, sql_select, idrecord)
	switch e := row.Scan(&date); e {
	case sql.ErrNoRows:
		fmt.Println("CHECKDBTHREEFIELD - No rows were returned!")
	case nil:
	default:

	}
	return date
}
