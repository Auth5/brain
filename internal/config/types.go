package config

type SiteConfig struct {
	Name   string `koanf:"name" validate:"required"`
	URL    string `koanf:"url" validate:"required,url"`
	APIURL string `koanf:"api_url" validate:"required,url"`
}

type ServerConfig struct {
	Host string `koanf:"host" validate:"required,ip"`
	Port int    `koanf:"port" validate:"required,min=1,max=65535"`
}

type SwaggerConfig struct {
	Web  bool   `koanf:"web"`
	Path string `koanf:"path" validate:"required"`
}

type WebhookConfig struct {
	Secret string `koanf:"secret" validate:"required"`
}

type StripeConfig struct {
	SecretKey string        `koanf:"secret_key" validate:"required"`
	Webhook   WebhookConfig `koanf:"webhook" validate:"required"`
}

type MaxMindConfig struct {
	GeoLite2 GeoLite2Config `koanf:"geolite2" validate:"required"`
}

type GeoLite2Config struct {
	Country string `koanf:"country" validate:"required,url"`
}

type SentryConfig struct {
	DSN string `koanf:"dsn" validate:"required"`
}

type SMTPConfig struct {
	Name     string `koanf:"name" validate:"required"`
	From     string `koanf:"from" validate:"required,email"`
	Username string `koanf:"username" validate:"required,email"`
	Password string `koanf:"password" validate:"required"`
	Host     string `koanf:"host" validate:"required"`
	Port     int    `koanf:"port" validate:"required,min=1,max=65535"`
	TLS      bool   `koanf:"tls"`
}

type EmailConfig struct {
	Nickname string     `koanf:"nickname" validate:"required"`
	SMTP     SMTPConfig `koanf:"smtp" validate:"required"`
}

type CORSConfig struct {
	Origins []string `koanf:"origins" validate:"required,min=1,dive,url"`
}

type DatabaseConfig struct {
	MongoDB MongoDBConfig `koanf:"mongodb" validate:"required"`
	Badger  BadgerConfig  `koanf:"badger" validate:"required"`
}

type MongoDBConfig struct {
	URI    string `koanf:"uri" validate:"required"`
	DBName string `koanf:"db_name" validate:"required"`
}

type BadgerConfig struct {
	Dir string `koanf:"dir" validate:"required"`
}

type OAuthConfig struct {
	ClientID     string `koanf:"client_id" validate:"required"`
	ClientSecret string `koanf:"client_secret" validate:"required"`
	RedirectURL  string `koanf:"redirect_url" validate:"required,url"`
}

type OAuthProviders struct {
	Google OAuthConfig `koanf:"google" validate:"required"`
	GitHub OAuthConfig `koanf:"github" validate:"required"`
}

type Config struct {
	Server   ServerConfig   `koanf:"server" validate:"required"`
	Swagger  SwaggerConfig  `koanf:"swagger" validate:"required"`
	Stripe   StripeConfig   `koanf:"stripe" validate:"required"`
	MaxMind  MaxMindConfig  `koanf:"maxmind" validate:"required"`
	Sentry   SentryConfig   `koanf:"sentry" validate:"required"`
	Emails   []EmailConfig  `koanf:"emails" validate:"required,min=1,dive"`
	CORS     CORSConfig     `koanf:"cors" validate:"required"`
	Database DatabaseConfig `koanf:"database" validate:"required"`
	Site     SiteConfig     `koanf:"site" validate:"required"`
	OAuth    OAuthProviders `koanf:"oauth" validate:"required"`
}
