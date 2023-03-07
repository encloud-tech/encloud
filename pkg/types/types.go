package types

// ConfYaml
type ConfYaml struct {
	Estuary SectionEstuary `yaml:"estuary"`
	Email   EmailStat      `yaml:"email"`
	Stat    SectionStat    `yaml:"stat"`
}

// SectionEstuary is sub section of config.
type SectionEstuary struct {
	BaseApiUrl    string `yaml:"base_api_url"`
	UploadApiUrl  string `yaml:"upload_api_url"`
	GatewayApiUrl string `yaml:"gateway_api_url"`
	CdnApiUrl     string `yaml:"cdn_api_url"`
	Token         string `yaml:"token"`
}

// EmailStat is sub section of config.
type EmailStat struct {
	Server   string `yaml:"server"`
	Port     int64  `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	From     string `yaml:"from"`
}

// SectionStat is sub section of config.
type SectionStat struct {
	BadgerDB    SectionBadgerDB  `yaml:"badgerdb"`
	Couchbase   SectionCouchbase `yaml:"couchbase"`
	StorageType string           `yaml:"storageType"`
	KekType     string           `yaml:"kekType"`
}

// SectionBadgerDB is sub section of config.
type SectionBadgerDB struct {
	Path string `yaml:"path"`
}

// SectionCouchbae is sub section of config.
type SectionCouchbase struct {
	Host     string        `yaml:"host"`
	Username string        `yaml:"username"`
	Password string        `yaml:"password"`
	Bucket   SectionBucket `yaml:"bucket"`
}

type SectionBucket struct {
	Name       string `yaml:"name"`
	Scope      string `yaml:"scope"`
	Collection string `yaml:"collection"`
}

type Contents []Content

type ByCID []struct {
	Content ByCidResponse
}

type Content struct {
	Name string `json:"name"`
	CID  string `json:"cid"`
}

type UploadResponse struct {
	CID       string
	EstuaryId int
}

type ByCidResponse struct {
	Name string
	CID  string
}

type FileData []FileMetadata

type FileMetadata struct {
	Uuid       string   `json:"uuid"`
	Md5Hash    string   `json:"md5Hash"`
	Timestamp  int64    `json:"timestamp"`
	UploadedAt string   `json:"uploadedAt"`
	Name       string   `json:"name"`
	Size       int      `json:"size"`
	FileType   string   `json:"fileType"`
	Cid        []string `json:"cid"`
	Dek        []byte   `json:"dek"`
	DekType    string   `json:"dekType"`
	KekType    string   `json:"kekType"`
}

type GenerateKeyPairResponse struct {
	Status     string
	StatusCode int
	Message    string
	Data       Keys
}

type UploadContentResponse struct {
	Status     string
	StatusCode int
	Message    string
	Data       Uuid
}

type Uuid struct {
	Uuid string
}

type ListContentResponse struct {
	Status     string
	StatusCode int
	Message    string
	Data       FileData
}

type RetrieveByCIDContentResponse struct {
	Status     string
	StatusCode int
	Message    string
	Data       FileMetadata
}

type Keys struct {
	PublicKey  string `json:"PublicKey"`
	PrivateKey string `json:"PrivateKey"`
}

type ConfigResponse struct {
	Status     string
	StatusCode int
	Message    string
	Data       ConfYaml
}

type SharedResponse struct {
	Status     string
	StatusCode int
	Message    string
}

type Error struct {
	Code    int64  `json:"code"`
	Reason  string `json:"reason"`
	Details string `json:"details"`
}

type EstuaryError struct {
	Error Error `json:"error"`
}
