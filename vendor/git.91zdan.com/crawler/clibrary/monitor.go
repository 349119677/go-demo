package clibrary

import (
	"database/sql"
	"fmt"
	"net/http"
)

type DoHttpRequest func(req *http.Request, jar http.CookieJar) (*http.Response, error)
type PingMongo func() error

const NETWORK_MONITOR_URL = "https://www.baidu.com"

// 监控网络情况
// 访问 https://www.baidu.com 测试
func MonitorNetwork(do DoHttpRequest) error {
	req, err := http.NewRequest(http.MethodGet, NETWORK_MONITOR_URL, nil)
	if err != nil {
		return err
	}
	res, err := do(req, nil)
	if err != nil {
		log.Warningf("[Monitor] Network offline! Error: %s", err.Error())
		return err
	}
	defer res.Body.Close()
	return nil
}

// 监控 MySQL 服务状态
func MonitorMySQL(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		log.Warningf("[Monitor] MySQL is down! Error: %s", err.Error())
	}
	return err
}

// 监控 MongoDB 服务状态
func MonitorMongo(ping PingMongo) error {
	err := ping()
	if err != nil {
		log.Warningf("[Monitor] Mongo is down! Error: %s", err.Error())
	}
	return err
}

// 监控方案 A
// 监控网络情况和 MySQL 连接
// 服务正常返回 http code 200, 服务宕机返回 http code 503
// 并在 http body 体内返回消息
func MonitorSchemeA(w http.ResponseWriter, do DoHttpRequest, db *sql.DB) {
	isDown := false
	isNetworkDown := false
	isMySQLDown := false
	err := MonitorNetwork(do)
	if isNetworkDown = err != nil; isNetworkDown {
		isDown = true
	}
	err = MonitorMySQL(db)
	if isMySQLDown = err != nil; isMySQLDown {
		isDown = true
	}

	if isDown {
		w.WriteHeader(http.StatusServiceUnavailable)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	fmt.Fprint(w, fmt.Sprintf(`
		Network is down: %t
		MySQL is down: 	%t
	`, isNetworkDown, isMySQLDown))
}
