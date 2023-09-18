package postgres

type Config struct {
	Username string `koanf:"username"`
	Password string `koanf:"password"`
	Host     string `koanf:"host"`
	Name     string `koanf:"name"`
}
