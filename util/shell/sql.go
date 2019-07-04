package shell
import (
    "github.com/LongMarch7/higo/config"
    "strings"
)
func Database_Create(config config.Sql) (string, error) {
    sql := strings.Replace(`CREATE DATABASE {database} DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;`, "{database}", config.Db, -1)
    command := "mysql -P {port} -h {address} -u{username} -p{password} -e\"" + sql + "\""
    command = strings.Replace(command, "{username}", config.User, 1) //username
    command = strings.Replace(command, "{password}", config.Pwd, 1) //password
    command = strings.Replace(command, "{address}", config.Addr, 1)  //address
    command = strings.Replace(command, "{port}", config.Port, 1)         //port
    return Exec_Shell(command)
}

func Database_Import(config config.Sql) (string, error) {
    command := "mysql -P {port} -h {address} -u{username} -p{password} {database} < {source}"
    command = strings.Replace(command, "{username}", config.User, 1) //username
    command = strings.Replace(command, "{password}", config.Pwd, 1) //password
    command = strings.Replace(command, "{database}", config.Db, 1)          //database
    command = strings.Replace(command, "{address}", config.Addr, 1)  //address
    command = strings.Replace(command, "{source}", config.File, 1)             //sql
    command = strings.Replace(command, "{port}", config.Port, 1)         //port
    return Exec_Shell(command)
}