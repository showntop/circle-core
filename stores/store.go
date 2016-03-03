package stores

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq"
	"github.com/showntop/circle-core/logger"
)

type Store struct {
	master   *gorp.DbMap
	replicas []*gorp.DbMap
}

func NewStore(config map[string]interface{}) *Store {
	store := &Store{}
	store.master = setupConnection("master", config["DriverName"].(string),
		config["DataSource"].(string), int(config["MaxIdleNum"].(float64)), int(config["MaxOpenNum"].(float64)), config["Trace"].(bool))
	return store
}

func setupConnection(clusterType string, driver string, dataSource string, maxIdle int, maxOpen int, trace bool) *gorp.DbMap {

	db, err := sql.Open(driver, dataSource)
	if err != nil {
		//log
		time.Sleep(time.Second)
		panic(fmt.Sprintf("store.sql.open_conn.critical", err.Error()))
	}

	err = db.Ping()
	if err != nil {

		time.Sleep(time.Second)
		panic(fmt.Sprintf("store.sql.open_conn.panic", err.Error()))
	}

	db.SetMaxIdleConns(maxIdle)
	db.SetMaxOpenConns(maxOpen)

	logger.Info("connect db success..")
	////根据数据库配置判断
	dbmap := &gorp.DbMap{Db: db, TypeConverter: circleConverter{}, Dialect: gorp.PostgresDialect{}}

	if trace {
		dbmap.TraceOn("", log.New(os.Stdout, "sql-trace:", log.Lmicroseconds))
	}

	return dbmap
}

func (ss Store) GetMaster() *gorp.DbMap {
	return ss.master
}

func (ss Store) GetReplica() *gorp.DbMap {
	return nil
}

func (ss Store) GetAllConns() []*gorp.DbMap {
	all := make([]*gorp.DbMap, len(ss.replicas)+1)
	copy(all, ss.replicas)
	all[len(ss.replicas)] = ss.master
	return all
}

func (ss Store) Close() {

	ss.master.Db.Close()
	for _, replica := range ss.replicas {
		replica.Db.Close()
	}
}

type circleConverter struct{}

func (me circleConverter) ToDb(val interface{}) (interface{}, error) {

	return val, nil
}

func (me circleConverter) FromDb(target interface{}) (gorp.CustomScanner, bool) {

	return gorp.CustomScanner{}, false
}
