package setting

type Config struct {
	Mysql MysqlSetting `mapstructure:"mysql"`
}

type MysqlSetting struct {
	Host            string `mapstructure:"host"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Dbname          string `mapstructure:"dbname"`
	Port            int    `mapstructure:"port"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
}
