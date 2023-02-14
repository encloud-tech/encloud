package service

import (
	"encloud/config"
	"encloud/pkg/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDBService(t *testing.T) {
	cfg, err := config.LoadConf()
	if err != nil {
		panic("failed to load config.yaml from file")
	}
	dbService := NewDB(cfg)

	// Store data
	var cids []string
	cids = append(cids, "bafkreie3e2vuaydtlyrvoddsghzipjculmsbu2qz2izytkecy3qq3yooja")
	uuid := "efcb05f1-8cf7-4d4a-8c0e-2a95090d29eb"
	var hash = "3a4aa5ad1efff0897afc9e122725fad7"
	fileData := types.FileMetadata{
		Timestamp:  1674753941,
		UploadedAt: "2006-01-02 15:04:05",
		Name:       "Provider2InputData.csv",
		Size:       1070,
		FileType:   ".csv",
		Dek:        []byte("Q3H7Zni3EhuBh3D1nzgK89LaevTL351tUquA7SIT1qWxNSeNOxK44SF9Ylczm+YOfV8/Wj2DabU48YEWDXhM3LWHcf6lTYb9EbBWNHn48jmqh7u68UUCGTb4EQmr9N/3LIrPCf43nLlZf9jnZHs8yodbmAEpBjrZ3x2Z+k5ksKj0CgrLXeyWSYyBjK8518VZSy5rSfjz5AHgXTmTsYhCyYwwvdEiX1M7Krb8k5vb8cxZC0IAvov1XhGEYn8721fX4MiJmhv/g6/P4kidXSJOy/1uPAP5JKAAUCnjV10b/kmJd0KjGtzPA7JIUry23U0IwbkkwRoAWQNa7WpXedyfaxHj8gN7mVkvcLgPJ6iwkafSZCr+LVXlgNlNqwdkkBDfH849nEQjrQ5KE9u3fYVPenc2DhEheudsX1UY6wPPtq+rlscZJDq/AxbYeQbee9hONKcVBegVW8+3k8cPfVF0kq8E3UziXDgJtb8iSCSRsjCP8mg/q3uN+S6T95jTQsF2"),
		Cid:        cids,
		Uuid:       uuid,
		Md5Hash:    hash,
	}

	dbService.Store(hash+":"+uuid, fileData)

	// Fetch data
	records := dbService.Fetch(hash)
	assert.Equal(t, isExists(hash, records), true)

	// Fetch data by id
	record := dbService.FetchByCid(hash + ":" + uuid)
	assert.Equal(t, record.Md5Hash, hash)
	assert.Equal(t, record.Uuid, uuid)
}

func isExists(value string, data types.FileData) (exists bool) {
	for _, search := range data {
		if search.Md5Hash == value {
			return true
		}
	}
	return false
}
