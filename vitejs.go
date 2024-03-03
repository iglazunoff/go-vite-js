package go_vite_js

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

type ViteJs struct {
	config               *Config
	manifestItems        *ManifestItems
	baseUri              string
	productionUri        string
	developmentUri       string
	developmentClientUri string
	isProduction         bool
}

func NewViteJs(config *Config) *ViteJs {
	info("WORK")
	baseUri := config.httpScheme + "://" + config.httpHost
	productionUri := baseUri + "/" + config.assetsHttpPath
	developmentUri := baseUri + ":" + strconv.Itoa(config.httpPort)
	developmentClientUri := makeUriForDevClient(developmentUri)
	isProduction := checkIsProduction(developmentClientUri)
	manifestPath := config.buildDirectoryFsPath + "/" + manifestPathToFile

	info(fmt.Sprintf("Base URL:   %v", baseUri))

	var manifestItems *ManifestItems

	if isProduction {
		info(fmt.Sprintf("Prod URL:   %v", productionUri))

		manifestItems = importManifestItems(manifestPath)
	} else {
		info(fmt.Sprintf("Dev URL:    %v", developmentUri))
		info(fmt.Sprintf("Dev client: %v", developmentClientUri))

		emptyManifestItems := make(ManifestItems)
		manifestItems = &emptyManifestItems
	}

	viteJs := &ViteJs{
		config:               config,
		manifestItems:        manifestItems,
		baseUri:              baseUri,
		productionUri:        productionUri,
		developmentUri:       developmentUri,
		developmentClientUri: developmentClientUri,
		isProduction:         isProduction,
	}

	info("DONE")

	return viteJs
}

func (a *ViteJs) IsProduction() bool {
	return a.isProduction
}

func (a *ViteJs) Asset(path string) string {
	path = strings.Trim(path, "/")

	if !a.isProduction {
		return a.developmentUri + "/" + path
	}

	manifest := *a.manifestItems
	item := manifest[path]

	if item == nil {
		return ""
	}

	return a.productionUri + "/" + item.File
}

func (a *ViteJs) Client() string {
	return a.developmentClientUri
}

func GetLocalIp() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		panic(err)
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			panic(fmt.Sprintf("Inspected close connection error: %v", err.Error()))
		}
	}(conn)

	return conn.LocalAddr().(*net.UDPAddr).IP.String()
}
