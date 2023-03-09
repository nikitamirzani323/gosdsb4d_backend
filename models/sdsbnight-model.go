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

func Fetch_sdsbnightHome() (helpers.Response, error) {
	var obj entities.Model_sdsbnight
	var arraobj []entities.Model_sdsbnight
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
			id_sdsb4dnight , 
			to_char(COALESCE(date_sdsb4dnight,now()), 'YYYY-MM-DD '), 
			prize1_sdsb4dnight , prize2_sdsb4dnight, prize3_sdsb4dnight, 
			create_sdsb4dnight, to_char(COALESCE(createdate_sdsb4dnight,now()), 'YYYY-MM-DD HH24:MI:SS'),  
			update_sdsb4dnight, to_char(COALESCE(updatedate_sdsb4dnight,now()), 'YYYY-MM-DD HH24:MI:SS')   
			FROM ` + configs.DB_tbl_trx_sdsb4d_night + ` 
			ORDER BY date_sdsb4dnight DESC LIMIT 365 
		`

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			id_sdsb4dnight_db                                                                                  int
			date_sdsb4dnight_db, prize1_sdsb4dnight_db, prize2_sdsb4dnight_db, prize3_sdsb4dnight_db           string
			create_sdsb4dnight_db, createdate_sdsb4dnight_db, update_sdsb4dnight_db, updatedate_sdsb4dnight_db string
		)

		err = row.Scan(
			&id_sdsb4dnight_db, &date_sdsb4dnight_db, &prize1_sdsb4dnight_db, &prize2_sdsb4dnight_db, &prize3_sdsb4dnight_db,
			&create_sdsb4dnight_db, &createdate_sdsb4dnight_db, &update_sdsb4dnight_db, &updatedate_sdsb4dnight_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		if create_sdsb4dnight_db != "" {
			create = create_sdsb4dnight_db + ", " + createdate_sdsb4dnight_db
		}
		if update_sdsb4dnight_db != "" {
			update = update_sdsb4dnight_db + ", " + updatedate_sdsb4dnight_db
		}
		obj.Sdsbnight_id = id_sdsb4dnight_db
		obj.Sdsbnight_date = date_sdsb4dnight_db
		obj.Sdsbnight_prize1 = prize1_sdsb4dnight_db
		obj.Sdsbnight_prize2 = prize2_sdsb4dnight_db
		obj.Sdsbnight_prize3 = prize3_sdsb4dnight_db
		obj.Sdsbnight_create = create
		obj.Sdsbnight_update = update
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
func Save_sdsbnightHome(admin, tanggal, sData string, idrecord int) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	if sData == "New" {
		flag = CheckDB(configs.DB_tbl_trx_sdsb4d_night, "date_sdsb4dnight", tanggal)
		if !flag {
			sql_insert := `
				insert into
				` + configs.DB_tbl_trx_sdsb4d_night + ` (
					id_sdsb4dnight  , date_sdsb4dnight, create_sdsb4dnight, createdate_sdsb4dnight
				) values (
					$1 ,$2, $3, $4 
				)
			`
			field_column := configs.DB_tbl_trx_sdsb4d_night + tglnow.Format("YYYY")
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
func Save_sdsbnightGenerator(admin, field, prize, sData string, idrecord int) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()

	if sData == "Edit" {
		sql_update := `
			UPDATE 
			` + configs.DB_tbl_trx_sdsb4d_night + `  
			SET ` + field + ` =$1,  
			update_sdsb4dnight=$2, updatedate_sdsb4dnight=$3  
			WHERE id_sdsb4dnight =$4  
		`
		flag_update, msg_update := Exec_SQL(sql_update, configs.DB_tbl_trx_sdsb4d_night, "UPDATE",
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
func Save_Generatornight(admin string) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	for i := 0; i <= 299; i++ {
		tglnow2, _ := goment.New("2021-01-01")
		tanggal := tglnow2.Add(i, "days").Format("YYYY-MM-DD")
		flag = CheckDB(configs.DB_tbl_trx_sdsb4d_night, "date_sdsb4dnight", tanggal)
		if !flag {
			sql_insert := `
				insert into
				` + configs.DB_tbl_trx_sdsb4d_night + ` (
					id_sdsb4dnight , date_sdsb4dnight, prize1_sdsb4dnight, prize2_sdsb4dnight, prize3_sdsb4dnight, 
					create_sdsb4dnight, createdate_sdsb4dnight
				) values (
					$1 ,$2, $3, $4, $5,
					$6, $7 
				)
			`
			prize_1 := helpers.GenerateNumber(4)
			prize_2 := helpers.GenerateNumber(4)
			prize_3 := helpers.GenerateNumber(4)
			field_column := configs.DB_tbl_trx_sdsb4d_night + tglnow.Format("YYYY")
			idrecord_counter := Get_counter(field_column)
			flag_insert, msg_insert := Exec_SQL(sql_insert, configs.DB_tbl_trx_sdsb4d_night, "UPDATE",
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

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}
