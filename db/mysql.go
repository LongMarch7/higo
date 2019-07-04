package db


import (
    "github.com/go-xorm/core"
    "github.com/go-xorm/xorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var dataBase  map[string]*Db
type Db struct {
    opts   dbOpt
    engine *xorm.Engine
}

func init(){
    dataBase = make(map[string]*Db)
}

const  DefaultNAME = "default"
func defaultConfig() dbOpt{
    return dbOpt{
        dialect: "mysql",
        args:    "root:123456@tcp(localhost:13306)/higo?charset=utf8",
        show: false,
        level: "ERROR",
        maxOpenConns: 100,
        maxIdleConns: 100,
    }
}

func NewDb(name string, opts ...DOption) (db *Db){
    opt := defaultConfig()
    if value,ok := dataBase[name]; ok{
        if value.Engine() != nil {
            db = value
            return
        }
        opt = value.opts
        delete(dataBase, name)
    }
    for _, o := range opts {
        o(&opt)
    }
    engine, err := xorm.NewEngine(opt.dialect, opt.args)
    if err != nil {
        panic(err)
    }
    engine.SetMaxIdleConns(opt.maxIdleConns)
    engine.SetMaxOpenConns(opt.maxOpenConns)
    engine.ShowSQL(opt.show)
    engine.Logger().SetLevel(convertLogLevel(opt.level))
    db = &Db{
        engine: engine,
        opts: opt,
    }
    dataBase[name] = db
    return db
}

func DeleteDB(name string){
    if value,ok := dataBase[name]; ok{
        value.engine.Close()
        value.engine = nil
    }
}

func (db* Db)Engine() *xorm.Engine{
    return db.engine
}

func convertLogLevel(levelStr string) (level core.LogLevel) {
    switch levelStr {
    case "debug":
        fallthrough
    case "DEBUG":
        level = core.LOG_DEBUG
    case "info":
        fallthrough
    case "INFO":
        level = core.LOG_INFO
    case "warn":
        fallthrough
    case "WARN":
        level = core.LOG_WARNING
    case "error":
        fallthrough
    case "ERROR":
        level = core.LOG_ERR
    case "off":
        fallthrough
    case "OFF":
        level = core.LOG_OFF
    default:
        level = core.LOG_UNKNOWN
    }
    return
}