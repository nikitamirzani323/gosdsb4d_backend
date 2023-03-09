package models

import (
	"context"
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

func Fetch_sdsbdayHome() (helpers.Response, error) {
	var obj entities.Model_sdsbday
	var arraobj []entities.Model_sdsbday
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
			idsdsb4dday , 
			to_char(COALESCE(datesdsb4dday,now()), 'YYYY-MM-DD '), 
			prize1_sdsb4dday , prize2_sdsb4dday, prize3_sdsb4dday, 
			create_sdsb4dday, to_char(COALESCE(createdate_sdsb4dday,now()), 'YYYY-MM-DD HH24:MI:SS'),  
			update_sdsb4dday, to_char(COALESCE(updatedate_sdsb4dday,now()), 'YYYY-MM-DD HH24:MI:SS')   
			FROM ` + configs.DB_tbl_trx_sdsb4d_day + ` 
			ORDER BY datesdsb4dday DESC LIMIT 365 
		`

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idsdsb4dday_db                                                                             int
			datesdsb4dday_db, prize1_sdsb4dday_db, prize2_sdsb4dday_db, prize3_sdsb4dday_db            string
			create_sdsb4dday_db, createdate_sdsb4dday_db, update_sdsb4dday_db, updatedate_sdsb4dday_db string
		)

		err = row.Scan(
			&idsdsb4dday_db, &datesdsb4dday_db, &prize1_sdsb4dday_db, &prize2_sdsb4dday_db, &prize3_sdsb4dday_db,
			&create_sdsb4dday_db, &createdate_sdsb4dday_db, &update_sdsb4dday_db, &updatedate_sdsb4dday_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		if create_sdsb4dday_db != "" {
			create = create_sdsb4dday_db + ", " + createdate_sdsb4dday_db
		}
		if update_sdsb4dday_db != "" {
			update = update_sdsb4dday_db + ", " + updatedate_sdsb4dday_db
		}
		obj.Sdsbday_id = idsdsb4dday_db
		obj.Sdsbday_date = datesdsb4dday_db
		obj.Sdsbday_prize1 = prize1_sdsb4dday_db
		obj.Sdsbday_prize2 = prize2_sdsb4dday_db
		obj.Sdsbday_prize3 = prize3_sdsb4dday_db
		obj.Sdsbday_create = create
		obj.Sdsbday_update = update
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
func Save_sdsbdayHome(admin, tanggal, sData string, idrecord int) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	if sData == "New" {
		flag = CheckDB(configs.DB_tbl_trx_sdsb4d_day, "datesdsb4dday", tanggal)
		if !flag {
			sql_insert := `
				insert into
				` + configs.DB_tbl_trx_sdsb4d_day + ` (
					idsdsb4dday , datesdsb4dday, create_sdsb4dday, createdate_sdsb4dday
				) values (
					$1 ,$2, $3, $4 
				)
			`
			field_column := configs.DB_tbl_trx_sdsb4d_day + tglnow.Format("YYYY")
			idrecord_counter := Get_counter(field_column)
			flag_insert, msg_insert := Exec_SQL(sql_insert, configs.DB_tbl_trx_sdsb4d_day, "UPDATE",
				tglnow.Format("YY")+strconv.Itoa(idrecord_counter),
				tanggal,
				admin,
				tglnow.Format("YYYY-MM-DD HH:mm:ss"))

			if !flag_insert {
				log.Println(msg_insert)
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
func Save_sdsbdayGenerator(admin, field, prize, sData string, idrecord int) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()

	if sData == "Edit" {
		sql_update := `
			UPDATE 
			` + configs.DB_tbl_trx_sdsb4d_day + `  
			SET ` + field + ` =$1,  
			update_sdsb4dday=$2, updatedate_sdsb4dday=$3 
			WHERE idsdsb4dday=$4  
		`
		flag_update, msg_update := Exec_SQL(sql_update, configs.DB_tbl_trx_sdsb4d_day, "UPDATE",
			prize, admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord)

		if !flag_update {
			log.Println(msg_update)
		}
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}
func Save_Generator(admin string) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	for i := 0; i <= 299; i++ {
		tglnow2, _ := goment.New("2021-01-01")
		tanggal := tglnow2.Add(i, "days").Format("YYYY-MM-DD")
		flag = CheckDB(configs.DB_tbl_trx_sdsb4d_day, "datesdsb4dday", tanggal)
		if !flag {
			sql_insert := `
				insert into
				` + configs.DB_tbl_trx_sdsb4d_day + ` (
					idsdsb4dday , datesdsb4dday, prize1_sdsb4dday, prize2_sdsb4dday, prize3_sdsb4dday, 
					create_sdsb4dday, createdate_sdsb4dday
				) values (
					$1 ,$2, $3, $4, $5,
					$6, $7 
				)
			`
			prize_1 := helpers.GenerateNumber(4)
			prize_2 := helpers.GenerateNumber(4)
			prize_3 := helpers.GenerateNumber(4)
			field_column := configs.DB_tbl_trx_sdsb4d_day + tglnow.Format("YYYY")
			idrecord_counter := Get_counter(field_column)
			flag_insert, msg_insert := Exec_SQL(sql_insert, configs.DB_tbl_trx_sdsb4d_day, "UPDATE",
				tglnow.Format("YY")+strconv.Itoa(idrecord_counter),
				tanggal,
				prize_1,
				prize_2,
				prize_3,
				admin,
				tglnow.Format("YYYY-MM-DD HH:mm:ss"))

			if !flag_insert {
				log.Println(msg_insert)
			}
		} else {
			break
		}
	}

	if flag {
		res.Status = fiber.StatusOK
		res.Message = msg
		res.Record = nil
		res.Time = time.Since(render_page).String()
	} else {
		res.Status = fiber.StatusBadRequest
		res.Message = msg
		res.Record = nil
		res.Time = time.Since(render_page).String()
	}

	return res, nil
}
