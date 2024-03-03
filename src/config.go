package vitejs

import (
	"strings"
)

const defaultDevServerClient = "@vite/client"

type Config struct {
	httpPort             int
	httpHost             string
	httpScheme           string
	buildDirectoryFsPath string
	assetsHttpPath       string
}

func NewConfig(httpPort int, httpHost string, httpScheme string, buildDirectoryFsPath string, assetsHttpPath string) *Config {
	config := Default()

	config.SetHttpPort(httpPort)
	config.SetHttpHost(httpHost)
	config.SetHttpScheme(httpScheme)
	config.SetBuildDirectoryFsPath(buildDirectoryFsPath)
	config.SetAssetsHttpPath(assetsHttpPath)

	return config
}

func Default() *Config {
	config := &Config{}

	config.SetHttpPort(5173)
	config.SetHttpHost(GetLocalIp())
	config.SetHttpScheme("http")
	config.SetBuildDirectoryFsPath("../public/build")
	config.SetAssetsHttpPath("assets")

	return config
}

func (c *Config) SetHttpPort(httpPort int) {
	isValidHttpPortOrFail(httpPort)

	c.httpPort = httpPort
}

func (c *Config) SetHttpHost(httpHost string) {
	c.httpHost = strings.Trim(httpHost, "/")
}

func (c *Config) SetHttpScheme(httpScheme string) {
	isValidHttpSchemeOrFail(httpScheme)

	c.httpScheme = httpScheme
}

func (c *Config) SetBuildDirectoryFsPath(buildDirectoryFsPath string) {
	c.buildDirectoryFsPath = strings.TrimRight(buildDirectoryFsPath, "/")
}

func (c *Config) SetAssetsHttpPath(assetsHttpPath string) {
	c.assetsHttpPath = strings.Trim(assetsHttpPath, "/")
}
