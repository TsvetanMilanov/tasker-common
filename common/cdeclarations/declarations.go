package cdeclarations

import "github.com/TsvetanMilanov/tasker-common/common/ctypes"

// IHTTPClient is client which helps with HTTP requests
type IHTTPClient interface {
	PostJSON(url string, body interface{}, headers map[string]string, out interface{}) error
	GetJSON(url string, headers map[string]string, out interface{}) error
}

// IConfig handles common server configurations.
type IConfig interface {
	GetAuth0Config() ctypes.Auth0Config
	GetAuth0ManagementConfig() ctypes.Auth0MgmtConfig
}
