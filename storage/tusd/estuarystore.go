package tusd

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"filecoin-encrypted-data-storage/config"
	"filecoin-encrypted-data-storage/service"
	thirdparty "filecoin-encrypted-data-storage/third_party"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/tus/tusd/pkg/handler"
)

var defaultFilePerm = os.FileMode(0664)

// See the handler.DataStore interface for documentation about the different
// methods.
type EstuaryStore struct {
	// Relative or absolute path to store files in. EstuaryStore does not check
	// whether the path exists, use os.MkdirAll in this case on your own.
	Path string
}

// New creates a new file based storage backend. The directory specified will
// be used as the only storage entry. This method does not check
// whether the path exists, use os.MkdirAll to ensure.
// In addition, a locking mechanism is provided.
func New(path string) EstuaryStore {
	return EstuaryStore{path}
}

// UseIn sets this store as the core data store in the passed composer and adds
// all possible extension to it.
func (store EstuaryStore) UseIn(composer *handler.StoreComposer) {
	composer.UseCore(store)
	composer.UseTerminater(store)
	composer.UseConcater(store)
	composer.UseLengthDeferrer(store)
}

func (store EstuaryStore) NewUpload(ctx context.Context, info handler.FileInfo) (handler.Upload, error) {
	id := thirdparty.Uid()
	binPath := store.binPath(id)
	info.ID = id
	info.Storage = map[string]string{
		"Type": "EstuaryStore",
		"Path": binPath,
	}

	// Create binary file with no content
	file, err := os.OpenFile(binPath, os.O_CREATE|os.O_WRONLY, defaultFilePerm)
	if err != nil {
		if os.IsNotExist(err) {
			err = fmt.Errorf("assets directory does not exist: %s", store.Path)
		}
		return nil, err
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}

	//generate a random 32 byte key for AES-256
	dek := make([]byte, 32)
	if _, err := rand.Read(dek); err != nil {
		log.Fatalf("cipher err 123: %v", err.Error())
	}

	upload := &fileUpload{
		info:     info,
		infoPath: store.infoPath(id),
		binPath:  store.binPath(id),
		dek:      dek,
	}

	// writeInfo creates the file by itself if necessary
	err = upload.writeInfo()
	if err != nil {
		return nil, err
	}

	return upload, nil
}

func (store EstuaryStore) GetUpload(ctx context.Context, id string) (handler.Upload, error) {
	info := handler.FileInfo{}
	data, err := ioutil.ReadFile(store.infoPath(id))
	if err != nil {
		if os.IsNotExist(err) {
			// Interpret os.ErrNotExist as 404 Not Found
			err = handler.ErrNotFound
		}
		return nil, err
	}
	if err := json.Unmarshal(data, &info); err != nil {
		return nil, err
	}

	binPath := store.binPath(id)
	infoPath := store.infoPath(id)
	stat, err := os.Stat(binPath)
	if err != nil {
		if os.IsNotExist(err) {
			// Interpret os.ErrNotExist as 404 Not Found
			err = handler.ErrNotFound
		}
		return nil, err
	}

	info.Offset = stat.Size()

	return &fileUpload{
		info:     info,
		binPath:  binPath,
		infoPath: infoPath,
	}, nil
}

func (store EstuaryStore) AsTerminatableUpload(upload handler.Upload) handler.TerminatableUpload {
	return upload.(*fileUpload)
}

func (store EstuaryStore) AsLengthDeclarableUpload(upload handler.Upload) handler.LengthDeclarableUpload {
	return upload.(*fileUpload)
}

func (store EstuaryStore) AsConcatableUpload(upload handler.Upload) handler.ConcatableUpload {
	return upload.(*fileUpload)
}

// binPath returns the path to the file storing the binary data.
func (store EstuaryStore) binPath(id string) string {
	return filepath.Join(store.Path, id)
}

// infoPath returns the path to the .info file storing the file's info.
func (store EstuaryStore) infoPath(id string) string {
	return filepath.Join(store.Path, id+".info")
}

type fileUpload struct {
	// info stores the current information about the upload
	info handler.FileInfo
	// infoPath is the path to the .info file
	infoPath string
	// binPath is the path to the binary file (which has no extension)
	binPath string
	// dek
	dek []byte
}

func (upload *fileUpload) GetInfo(ctx context.Context) (handler.FileInfo, error) {
	return upload.info, nil
}

func (upload *fileUpload) WriteChunk(ctx context.Context, offset int64, src io.Reader) (int64, error) {
	file, err := os.OpenFile(upload.binPath, os.O_WRONLY|os.O_APPEND, defaultFilePerm)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	n, err := io.Copy(buf, src)
	if err != nil {
		log.Fatalf("cipher err 11: %v", err.Error())
	}

	// Creating block of algorithm
	block, err := aes.NewCipher(upload.dek)
	if err != nil {
		log.Fatalf("cipher err: %v", err.Error())
	}

	// Creating GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("cipher GCM err: %v", err.Error())
	}

	// Generating random nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatalf("nonce  err: %v", err.Error())
	}

	// Decrypt file
	cipherText := gcm.Seal(nonce, nonce, buf.Bytes(), nil)

	// Writing ciphertext file
	err = ioutil.WriteFile(upload.binPath, cipherText, 0777)
	if err != nil {
		log.Fatalf("write file err: %v", err.Error())
	}

	upload.info.Offset += n
	return n, err
}

func (upload *fileUpload) GetReader(ctx context.Context) (io.Reader, error) {
	return os.Open(upload.binPath)
}

func (upload *fileUpload) Terminate(ctx context.Context) error {
	if err := os.Remove(upload.infoPath); err != nil {
		return err
	}
	if err := os.Remove(upload.binPath); err != nil {
		return err
	}
	return nil
}

func (upload *fileUpload) ConcatUploads(ctx context.Context, uploads []handler.Upload) (err error) {
	file, err := os.OpenFile(upload.binPath, os.O_WRONLY|os.O_APPEND, defaultFilePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, partialUpload := range uploads {
		fileUpload := partialUpload.(*fileUpload)

		src, err := os.Open(fileUpload.binPath)
		if err != nil {
			return err
		}

		log.Println(file)
		if _, err := io.Copy(file, src); err != nil {
			return err
		}
	}

	return
}

func (upload *fileUpload) DeclareLength(ctx context.Context, length int64) error {
	upload.info.Size = length
	upload.info.SizeIsDeferred = false
	return upload.writeInfo()
}

// writeInfo updates the entire information. Everything will be overwritten.
func (upload *fileUpload) writeInfo() error {
	data, err := json.Marshal(upload.info)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(upload.infoPath, data, defaultFilePerm)
}

func (upload *fileUpload) FinishUpload(ctx context.Context) error {
	cfg, _ := config.LoadConf("config.yml")
	estuaryService := service.New(cfg)
	content := estuaryService.UploadContent(upload.binPath)
	log.Println(content)
	return nil
}
