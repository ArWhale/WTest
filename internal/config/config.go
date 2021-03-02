package config

import (
	"errors"
	"github.com/spf13/viper"
	"go.uber.org/multierr"
	"path/filepath"
	"strings"
)

func checkParamsViper() error {
	var validErrors error
	if viper.GetString(ServiceHostKey) == "" {
		validErrors = multierr.Append(validErrors, errors.New(ServiceHostKey+" param not found"))
	}
	if viper.GetString(ServicePortKey) == "" {
		validErrors = multierr.Append(validErrors, errors.New(ServicePortKey+" param not found"))
	}
	if viper.GetString(ServiceDBUrlKey) == "" {
		validErrors = multierr.Append(validErrors, errors.New(ServiceDBUrlKey+" param not found"))
	}
	if viper.GetString(ServiceLogOutputKey) == "" {
		validErrors = multierr.Append(validErrors, errors.New(ServiceLogOutputKey+" param not found"))
	}
	if viper.GetString(ServiceLogLevelKey) == "" {
		validErrors = multierr.Append(validErrors, errors.New(ServiceLogLevelKey+" param not found"))
	}
	if viper.GetString(ServiceLogFormatKey) == "" {
		validErrors = multierr.Append(validErrors, errors.New(ServiceLogFormatKey+" param not found"))
	}
	return validErrors
}

func NewConfigFromEnv() (*viper.Viper, error) {
	viper.SetEnvPrefix(DefaultPrefixForEnv)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	return viper.GetViper(), checkParamsViper()
}

func NewConfigFromFile(filePath, fileName, fileExtension string) (*viper.Viper, error) {
	if fileName != "" {
		viper.SetConfigName(fileName)
	}

	if filePath != "" {
		fullPath, err := filepath.Abs(filePath)
		if err != nil {
			return nil, err
		}
		viper.AddConfigPath(fullPath)
	}

	viper.SetConfigType(fileExtension)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return viper.GetViper(), checkParamsViper()
}
