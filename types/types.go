package types

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
	Uuid      string   `json:"uuid"`
	Md5Hash   string   `json:"md5Hash"`
	Timestamp int64    `json:"timestamp"`
	Name      string   `json:"name"`
	Size      int      `json:"size"`
	FileType  string   `json:"fileType"`
	Cid       []string `json:"cid"`
	Dek       []byte   `json:"dek"`
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

type ErrorResponse struct {
	Status     string
	StatusCode int
	Message    string
}
