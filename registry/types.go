package registry

type SearchResult struct {
	StarCount   int    `json:"star_count"`
	IsOfficial  bool   `json:"is_official"`
	Name        string `json:"name"`
	IsTrusted   bool   `json:"is_trusted"`
	Description string `json:"description"`
}

type SearchResults struct {
	Query      string         `json:"query"`
	NumResults int            `json:"num_results"`
	Results    []SearchResult `json:"results"`
}

type RepositoryData struct {
	ImgList   map[string]*ImgData
	Endpoints []string
	Tokens    []string
}

type ImgData struct {
	ID              string `json:"id"`
	Checksum        string `json:"checksum,omitempty"`
	ChecksumPayload string `json:"-"`
	Tag             string `json:",omitempty"`
}

type RegistryInfo struct {
	Version    string `json:"version"`
	Standalone bool   `json:"standalone"`
}

type APIVersion int

func (av APIVersion) String() string {
	return apiVersions[av]
}

var DefaultAPIVersion APIVersion = APIVersion1
var apiVersions = map[APIVersion]string{
	1: "v1",
	2: "v2",
}

const (
	_           = iota
	APIVersion1 = iota
	APIVersion2
)
