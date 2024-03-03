package vitejs

const manifestPathToFile = ".vite/manifest.json"

type ManifestItem struct {
	File    string `json:"file"`
	IsEntry bool   `json:"isEntry"`
	Src     string `json:"src"`
}

type ManifestItems = map[string]*ManifestItem
