package core

import (
	"github.com/spf13/viper"
	"github.com/jmoiron/sqlx"
	"github.com/gorilla/schema"
	"path"
	"os"
	"fmt"
	"runtime"
	log "github.com/Sirupsen/logrus"
)


// Core preovides ability to interface with basic data models
type Core struct {
	Config     *CoreConfig
	db         *sqlx.DB
	DebugMode  bool
	reqDecoder *schema.Decoder
}

type CoreConfig struct {
	DBConnStr string
}

// New instantiates a new database connection
func New(envName string, logLevel log.Level) *Core {
	c := &Core{}
	c.initConfig(envName)
	c.initLogger(logLevel)
	c.initDatabase()
	c.reqDecoder = schema.NewDecoder()
	return c
}

// Close closes any connections to the database
func (c *Core) Close() {
	c.db.Close()
	log.Debugln("disconnected from database")
}

func (c *Core) initConfig(envName string) {

	_, thisFilename, _, _ := runtime.Caller(1)
	configPath := path.Join(path.Dir(thisFilename), "config")

	viper.AddConfigPath(configPath)
	viper.SetConfigName(envName)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// load in-memory config settings
	conf := &CoreConfig{}
	c.Config = conf
	return
}

// initLogger sets log levels based on core config settings
func (c *Core) initLogger(logLevel log.Level) {
	log.Debugf("initializing logger")
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(logLevel)
	log.Infof("initialized logger to", log.GetLevel())
}

func (c *Core) initDatabase() {
	c.Config.DBConnStr = viper.GetString("db_connection_dsn")
	if len(c.Config.DBConnStr) == 0 {
		panic(fmt.Errorf("db_connection_dsn missing from config"))
	}

	// connect
	log.Infof("connecting to database: %v", c.Config.DBConnStr)
	sql, err := sqlx.Connect("mysql", c.Config.DBConnStr)
	if err != nil {
		panic(err.Error())
	}

	// force a connection and test ping
	err = sql.Ping()
	if err != nil {
		log.Errorf("couldn't connect to database: %v", c.Config.DBConnStr)
		panic(err.Error())
	}
	c.db = sql.Unsafe()
	log.Infoln("connected to database")
	return
}


