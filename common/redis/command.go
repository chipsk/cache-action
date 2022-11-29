package redis

//redis 命令集
const (
	commandAuth    = "AUTH"
	commandPing    = "PING"
	commandExpire  = "EXPIRE"
	commandPExpire = "PEXPIRE"
	commandExist   = "EXISTS"
	commandDel     = "DEL"
	commandSet     = "SET"
	commandSetEx   = "SETEX"
	commandPSetEx  = "PSETEX"
	commandGet     = "GET"
	commandIncr    = "INCR"
	commandIncrBy  = "INCRBY"

	commandScan  = "SCAN"
	commandHScan = "HSCAN"
)
