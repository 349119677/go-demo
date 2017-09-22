package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// 获取数据
func GetLogsOriginalBillData() []LogsOriginalBill {
	db, err := sql.Open("mysql", "crawler_read:=!!$z+q+)l_9$^^@tcp(rm-uf66up8u0x5447kvq.mysql.rds.aliyuncs.com:3306)/crawler?charset=utf8&loc=Local")
	if err != nil {
		panic(err)
	}
	//查询数据
	//rows, err := db.Query("SELECT id,crawl_id,platform,url,action,request_header,request_body,response_header,response_body,create_time FROM logs_original_bill limit 2")
	rows, err := db.Query("SELECT * FROM logs_original_bill limit 100")
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
	response := []LogsOriginalBill{}
	// 读取结果集
	for rows.Next() {
		exam := new(LogsOriginalBill)
		rows.Columns()
		err = rows.Scan(&exam.Id, &exam.Crawl_id, &exam.Platform, &exam.Url, &exam.Action, &exam.Request_header, &exam.Request_body, &exam.Response_header, &exam.Response_body, &exam.Create_time)
		if err != nil {
			log.Fatal(err)
		}
		response = append(response, *exam)
	}
	return response
}
