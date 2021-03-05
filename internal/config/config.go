package config

import (
	"errors"
	"github.com/SArtemJ/WTest/internal/consts"
	"github.com/spf13/viper"
	"go.uber.org/multierr"
	"strings"
)

func checkParamsViper() error {
	var validErrors error
	if viper.GetString(consts.ServiceHostKey) == "" {
		validErrors = multierr.Append(validErrors, errors.New(consts.ServiceHostKey+" param not found"))
	}
	if viper.GetString(consts.ServicePortKey) == "" {
		validErrors = multierr.Append(validErrors, errors.New(consts.ServicePortKey+" param not found"))
	}
	if viper.GetString(consts.ServiceDBUrlKey) == "" {
		validErrors = multierr.Append(validErrors, errors.New(consts.ServiceDBUrlKey+" param not found"))
	}
	if viper.GetString(consts.ServiceLogOutputKey) == "" {
		validErrors = multierr.Append(validErrors, errors.New(consts.ServiceLogOutputKey+" param not found"))
	}
	if viper.GetString(consts.ServiceLogLevelKey) == "" {
		validErrors = multierr.Append(validErrors, errors.New(consts.ServiceLogLevelKey+" param not found"))
	}
	if viper.GetString(consts.ServiceLogFormatKey) == "" {
		validErrors = multierr.Append(validErrors, errors.New(consts.ServiceLogFormatKey+" param not found"))
	}
	return validErrors
}

func NewConfigFromEnv() (*viper.Viper, error) {
	viper.SetEnvPrefix(consts.DefaultPrefixForEnv)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	return viper.GetViper(), checkParamsViper()
}
