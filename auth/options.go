package auth

import (
)
type authOpt struct {
	driverName         string
	dataSourceName     string
	baseName           string
	dbSpecified        bool
	ruleText           string
}
type AOption func(o *authOpt)

func DriverName(driverName string) AOption {
	return func(o *authOpt) {
		o.driverName = driverName
	}
}

func DataSourceName(dataSourceName string) AOption {
	return func(o *authOpt) {
		o.dataSourceName = dataSourceName
	}
}

func BaseName(baseName string) AOption {
	return func(o *authOpt) {
		o.baseName = baseName
	}
}

func DbSpecified(dbSpecified bool) AOption {
	return func(o *authOpt) {
		o.dbSpecified = dbSpecified
	}
}

func RuleText(ruleText string) AOption {
	return func(o *authOpt) {
		o.ruleText = ruleText
	}
}

