package couchbase

import (
	"fmt"
	"log"
	"time"

	"github.com/encloud-tech/encloud/pkg/types"

	"github.com/couchbase/gocb/v2"
)

// New func implements the storage interface
func New(config *types.ConfYaml) *Storage {
	return &Storage{
		config: config,
	}
}

// Storage is interface structure
type Storage struct {
	config  *types.ConfYaml
	cluster *gocb.Cluster
	bucket  *gocb.Bucket
}

// Init client storage.
func (s *Storage) Init() error {
	var err error
	s.cluster, err = gocb.Connect(s.config.Stat.Couchbase.Host, gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: s.config.Stat.Couchbase.Username,
			Password: s.config.Stat.Couchbase.Password,
		},
	})

	if err != nil {
		return err
	}

	err = s.cluster.WaitUntilReady(10*time.Second, nil)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	s.bucket = s.cluster.Bucket(s.config.Stat.Couchbase.Bucket.Name)

	err = s.bucket.WaitUntilReady(10*time.Second, nil)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return err
}

func (s *Storage) Create(key string, metadata types.FileMetadata) {
	col := s.bucket.Scope(s.config.Stat.Couchbase.Bucket.Scope).Collection(s.config.Stat.Couchbase.Bucket.Collection)

	// Create and store a Document
	_, err := col.Insert(key,
		types.FileMetadata{
			Name:       metadata.Name,
			Uuid:       metadata.Uuid,
			Md5Hash:    metadata.Md5Hash,
			PublicKey:  metadata.PublicKey,
			Size:       metadata.Size,
			Timestamp:  metadata.Timestamp,
			FileType:   metadata.FileType,
			Cid:        metadata.Cid,
			Dek:        metadata.Dek,
			KekType:    metadata.KekType,
			DekType:    metadata.DekType,
			UploadedAt: metadata.UploadedAt,
		}, nil)

	if err != nil {
		fmt.Printf("ERROR saving to couchbase db : %s\n", err)
	}
}

func (s *Storage) FetchKeys() types.ListKeys {
	col := s.bucket.Scope(s.config.Stat.Couchbase.Bucket.Scope)
	var ival types.ListKeys
	query := "SELECT publicKey, COUNT(*) As files FROM `" + s.config.Stat.Couchbase.Bucket.Name + "`.`" + s.config.Stat.Couchbase.Bucket.Scope + "`." + s.config.Stat.Couchbase.Bucket.Collection + " GROUP BY publicKey"
	queryResult, err := col.Query(
		query,
		&gocb.QueryOptions{},
	)
	if err != nil {
		log.Fatal(err)
	}

	// Print each found Row
	for queryResult.Next() {
		var result types.FetchKeysResponse
		err := queryResult.Row(&result)
		if err != nil {
			log.Fatal(err)
		}
		ival = append(ival, result)
	}

	if err := queryResult.Err(); err != nil {
		log.Fatal(err)
	}

	return ival
}

func (s *Storage) Read(key string) types.FileData {
	col := s.bucket.Scope(s.config.Stat.Couchbase.Bucket.Scope)
	var ival types.FileData
	query := "SELECT x.* FROM `" + s.config.Stat.Couchbase.Bucket.Collection + "` x WHERE x.md5Hash='" + key + "';"
	queryResult, err := col.Query(
		query,
		&gocb.QueryOptions{},
	)
	if err != nil {
		log.Fatal(err)
	}

	// Print each found Row
	for queryResult.Next() {
		var result types.FileMetadata
		err := queryResult.Row(&result)
		if err != nil {
			log.Fatal(err)
		}
		ival = append(ival, result)
	}

	if err := queryResult.Err(); err != nil {
		log.Fatal(err)
	}

	return ival
}

func (s *Storage) ReadByCid(key string) types.FileMetadata {
	col := s.bucket.Scope(s.config.Stat.Couchbase.Bucket.Scope).Collection(s.config.Stat.Couchbase.Bucket.Collection)
	queryResult, err := col.Get(key, nil)
	if err != nil {
		log.Fatal(err)
	}

	var ival types.FileMetadata
	err = queryResult.Content(&ival)
	if err != nil {
		panic(err)
	}

	return ival
}
