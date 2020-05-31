package lib

type Config struct {
	Personal Personal `mapstructure:"personal"`
}

type Personal struct {
	Name    string `mapstructure:"name"`
	Address string `mapstructure:"address"`
	Phone   string `mapstructure:"phone"`
	email   string `mapstructure:"email"`
}
