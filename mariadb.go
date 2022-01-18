package mariadb

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	serverTags map[string]*sqlx.DB = make(map[string]*sqlx.DB)
)

func New(serverTag, ip, port, username, password, dbname string) error {
	connstr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		username, password, ip, port, dbname)
	sqldb, err := sqlx.Open("mysql", connstr)
	//sqldb.SetConnMaxLifetime(time.Second * 5)
	if err != nil {
		return err
	}
	err = sqldb.Ping()
	if err != nil {
		return err
	}
	serverTags[serverTag] = sqldb
	return nil
}

// 该方法端口号为int，主要配合新的conf文件格式
func New2(serverTag, ip string, port int, username, password, dbname string) error {
	connstr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		username, password, ip, port, dbname)
	sqldb, err := sqlx.Open("mysql", connstr)
	//sqldb.SetConnMaxLifetime(time.Second * 5)
	if err != nil {
		return err
	}
	err = sqldb.Ping()
	if err != nil {
		return err
	}
	serverTags[serverTag] = sqldb
	return nil
}

func Connect(serverTag string) (*sqlx.DB, error) {
	sqldb, ok := serverTags[serverTag]
	if !ok {
		return nil, fmt.Errorf("mariadb[%s] not existing", serverTag)
	}
	return sqldb, nil
}

func Destroy() {
	for k := range serverTags {
		delete(serverTags, k)
	}
}

func SetConnMaxLifetime(serverTag string, d time.Duration) error {
	sqldb, ok := serverTags[serverTag]
	if !ok {
		return fmt.Errorf("mariadb[%s] not existing", serverTag)
	}
	sqldb.SetConnMaxLifetime(d)
	return nil
}
