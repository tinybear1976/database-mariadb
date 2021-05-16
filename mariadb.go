package mariadb

import (
	"errors"
	"fmt"

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

func Connect(serverTag string) (*sqlx.DB, error) {
	sqldb, ok := serverTags[serverTag]
	if !ok {
		return nil, errors.New(fmt.Sprintf("mariadb[%s] not existing", serverTag))
	}
	return sqldb, nil
}

func Destroy() {
	for k := range serverTags {
		delete(serverTags, k)
	}
}
