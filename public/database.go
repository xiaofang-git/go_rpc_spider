package public


var DbConfig = map[string]interface{} {
	"Default":         "mysql_dev",// 默认数据库配置
	"SetMaxOpenConns": 300,          // (连接池)最大打开的连接数，默认值为0表示不限制
	"SetMaxIdleConns": 10,          // (连接池)闲置的连接数, 默认1

	"Connections":map[string]map[string]string{
		"mysql_dev": {// 定义名为 mysql_dev 的数据库配置
			"host": "127.0.0.1", // 数据库地址
			"username": "root",       // 数据库用户名
			"password": "0805",       // 数据库密码
			"port": "3306",            // 端口
			"database": "spider",        // 链接的数据库名字
			"charset": "utf8",         // 字符集
			"protocol": "tcp",         // 链接协议
			"prefix": "",              // 表前缀
			"driver": "mysql",         // 数据库驱动(mysql,sqlite,postgres,oracle,mssql)
		},
		"sqlite_dev": {
		 "database": "./foo.db",
		 "prefix": "",
		 "driver": "sqlite3",
		},
	},
}