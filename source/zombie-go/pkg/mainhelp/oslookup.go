package mainhelp

import (
	"fmt"
	"strings"
)

type OsLookupEnv func(string) (string, bool)

func (m *mainHelp) GetAllowedCorsOrigins(osLookupEnv OsLookupEnv, key string) (map[string]bool, error) {
	val, found := osLookupEnv(key)
	if !found {
		return nil, fmt.Errorf("could not find environment variable %s", val)
	}

	allowed := make(map[string]bool)
	for _, cors := range strings.Split(val, ",") {
		allowed[cors] = true
	}

	return allowed, nil
}

func (m *mainHelp) PrintAllowedCorsOrigins(allowedCorsOrigins map[string]bool) {
	m.log.Info("ALLOWED_CORS_ORIGINS")
	for k := range allowedCorsOrigins {
		m.log.Infof("- %s\n", k)
	}
}
