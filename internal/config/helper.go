package config

import (
	"fmt"
)

func GetSiteConfig() *SiteConfig {
	return &Cfg.Site
}

func GetServerConfig() *ServerConfig {
	return &Cfg.Server
}

func GetSwaggerConfig() *SwaggerConfig {
	return &Cfg.Swagger
}

func GetStripeConfig() *StripeConfig {
	return &Cfg.Stripe
}

func GetMaxMind() *MaxMindConfig {
	return &Cfg.MaxMind
}

func GetSentryConfig() *SentryConfig {
	return &Cfg.Sentry
}

func GetSMTPConfig(nickname string) (*SMTPConfig, error) {
	for _, email := range Cfg.Emails {
		if email.Nickname == nickname {
			return &email.SMTP, nil
		}
	}
	return nil, fmt.Errorf("SMTP configuration not found for nickname: %s", nickname)
}

func GetCORSConfig() *CORSConfig {
	return &Cfg.CORS
}

func GetDatabaseConfig() *DatabaseConfig {
	return &Cfg.Database
}

func GetOauthConfig() *OAuthProviders {
	return &Cfg.OAuth
}
