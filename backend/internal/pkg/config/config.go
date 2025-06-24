package config

import (
	"os"
	"strconv"
	"sync"
)

// Config holds all the configuration for the application
type Config struct {
	// Environment (development, staging, production)
	ApiEnv string

	// Client configuration
	FrontUrl string

	// Server configuration
	ServerHost string
	ServerPort string

	// Encrypt Key
	EncryptionKey string

	// Database configuration
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	// Logger configuration
	LogLevel     string
	LogDirectory string
	EnableSQLLog bool
	SqlLogLevel  string

	// Authentication configuration
	JWTSecret       string
	JWTDurationHour int

	// Two Factor Authentication configuration
	MFATokenExpiryMinutes  int
	MFATokenResendInterval int

	// Email configuration
	SMTPHost         string
	SMTPPort         int
	SMTPUsername     string
	SMTPPassword     string
	SMTPFromEmail    string
	SMTPFromName     string
	SMTPUseAuth      bool
	SMTPUseTLS       bool
	EmailTemplateDir string

	// AWS Configuration
	AwsAccessKeyID     string
	AwsSecretAccessKey string
	AwsRegion          string

	// S3 configuration
	S3Bucket string
	S3Region string

	// SSH configuration
	SFTPUser                   string
	SFTPHost                   string
	SFTPPort                   int
	SFTPPrivateKey             string
	SFTPHostKeyFingerprint     string // SHA256 fingerprint for host key verification
	RemoteDir                  string
	LocalDir                   string
	PaypayPayinTransactionPath string
	PaypayPayinReportPath      string

	ProviderID int

	// Makeshop service configuration
	MakeshopBaseURL string
	// GMO Aozora Net Bank configuration
	GmoAozoraNetBankAPIEndPoint      string
	GmoAozoraNetBankAuthCallbackURL  string
	GmoAozoraNetBankClientID         string
	GmoAozoraNetBankClientSecret     string
	GmoAozoraNetBankAuthScope        string
	GmoAozoraNetBankPrimaryAccountID string

	// Makeshop SQS configuration
	MakeshopSQSQueueURL string
}

var (
	configInstance *Config
	once           sync.Once
)

// LoadConfig loads the configuration from environment variables
func LoadConfig() *Config {
	once.Do(func() {
		sqlLogLevel := "warn"
		if os.Getenv("API_ENV") == "local" {
			sqlLogLevel = "debug"
		}

		configInstance = &Config{
			ServerHost:             "0.0.0.0",
			ServerPort:             "8080",
			LogLevel:               "warn",
			LogDirectory:           "/app/logs",
			EnableSQLLog:           true,
			SqlLogLevel:            sqlLogLevel,
			JWTDurationHour:        24,
			MFATokenExpiryMinutes:  30,
			MFATokenResendInterval: 1,
			EmailTemplateDir:       "templates/email",
			SMTPFromName:           "Makeshop Payment",
			SMTPUseAuth:            true,
			SMTPUseTLS:             true,
			ProviderID:             1,
		}

		envVars := map[string]*string{
			"API_ENV":                       &configInstance.ApiEnv,
			"SERVER_HOST":                   &configInstance.ServerHost,
			"SERVER_PORT":                   &configInstance.ServerPort,
			"FRONT_URL":                     &configInstance.FrontUrl,
			"DB_HOST":                       &configInstance.DBHost,
			"DB_PORT":                       &configInstance.DBPort,
			"DB_USER":                       &configInstance.DBUser,
			"DB_PASSWORD":                   &configInstance.DBPassword,
			"DB_NAME":                       &configInstance.DBName,
			"LOG_LEVEL":                     &configInstance.LogLevel,
			"SQL_LOG_LEVEL":                 &configInstance.SqlLogLevel,
			"LOG_DIRECTORY":                 &configInstance.LogDirectory,
			"JWT_SECRET":                    &configInstance.JWTSecret,
			"SMTP_HOST":                     &configInstance.SMTPHost,
			"SMTP_USERNAME":                 &configInstance.SMTPUsername,
			"SMTP_PASSWORD":                 &configInstance.SMTPPassword,
			"SMTP_FROM_EMAIL":               &configInstance.SMTPFromEmail,
			"SMTP_FROM_NAME":                &configInstance.SMTPFromName,
			"EMAIL_TEMPLATE_DIR":            &configInstance.EmailTemplateDir,
			"S3_BUCKET":                     &configInstance.S3Bucket,
			"S3_REGION":                     &configInstance.S3Region,
			"AWS_REGION":                    &configInstance.AwsRegion,
			"AWS_ACCESS_KEY_ID":             &configInstance.AwsAccessKeyID,
			"AWS_SECRET_ACCESS_KEY":         &configInstance.AwsSecretAccessKey,
			"SFTP_USER":                     &configInstance.SFTPUser,
			"SFTP_HOST":                     &configInstance.SFTPHost,
			"SFTP_PRIVATE_KEY":              &configInstance.SFTPPrivateKey,
			"SFTP_HOST_KEY_FINGERPRINT":     &configInstance.SFTPHostKeyFingerprint,
			"REMOTE_DIR":                    &configInstance.RemoteDir,
			"PAYPAY_PAYIN_TRANSACTION_PATH": &configInstance.PaypayPayinTransactionPath,
			"PAYPAY_PAYIN_REPORT_PATH":      &configInstance.PaypayPayinReportPath,
			"LOCAL_DIR":                     &configInstance.LocalDir,
			"GMO_AOZORA_API_END_POINT":      &configInstance.GmoAozoraNetBankAPIEndPoint,
			"GMO_AOZORA_AUTH_CALLBACK_URL":  &configInstance.GmoAozoraNetBankAuthCallbackURL,
			"GMO_AOZORA_CLIENT_ID":          &configInstance.GmoAozoraNetBankClientID,
			"GMO_AOZORA_CLIENT_SECRET":      &configInstance.GmoAozoraNetBankClientSecret,
			"GMO_AOZORA_AUTH_SCOPE":         &configInstance.GmoAozoraNetBankAuthScope,
			"GMO_AOZORA_PRIMARY_ACCOUNT_ID": &configInstance.GmoAozoraNetBankPrimaryAccountID,
			"ENCRYPTION_KEY":                &configInstance.EncryptionKey,
			"MAKESHOP_SQS_QUEUE_URL":        &configInstance.MakeshopSQSQueueURL,
			"MAKESHOP_BASE_URL":             &configInstance.MakeshopBaseURL,
		}

		for env, field := range envVars {
			if val := os.Getenv(env); val != "" {
				*field = val
			}
		}

		// Override boolean fields
		boolVars := map[string]*bool{
			"ENABLE_SQL_LOG": &configInstance.EnableSQLLog,
			"SMTP_USE_AUTH":  &configInstance.SMTPUseAuth,
			"SMTP_USE_TLS":   &configInstance.SMTPUseTLS,
		}

		for env, field := range boolVars {
			if val := os.Getenv(env); val != "" {
				parsedVal, err := strconv.ParseBool(val)
				if err == nil {
					*field = parsedVal
				}
			}
		}

		intVars := map[string]*int{
			"JWT_EXPIRATION_HOURS": &configInstance.JWTDurationHour,
			"SMTP_PORT":            &configInstance.SMTPPort,
			"SFTP_PORT":            &configInstance.SFTPPort,
		}

		for env, field := range intVars {
			if val := os.Getenv(env); val != "" {
				intVal, err := strconv.Atoi(val)
				if err == nil {
					*field = intVal
				}
			}
		}
	})

	return configInstance
}

// GetConfig returns the singleton configInstance of the application configuration
func GetConfig() *Config {
	if configInstance == nil {
		return LoadConfig()
	}
	return configInstance
}

func IsLocal() bool {
	return GetConfig().ApiEnv == "local"
}

func IsDevelopment() bool {
	return GetConfig().ApiEnv == "dev"
}
