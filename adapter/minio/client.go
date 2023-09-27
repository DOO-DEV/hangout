package minioadapter

import (
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Config struct {
	Endpoint string `koanf:"endpoint"`
	Username string `koanf:"username"`
	Password string `koanf:"password"`
}

type Adapter struct {
	cfg    Config
	client *minio.Client
}

func New(cfg Config) *Adapter {
	minioClient, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.Username, cfg.Password, ""),
		Secure: false,
	})

	if err != nil {
		fmt.Println("error", err)
		fmt.Printf("can't create a new minio client: %s\n", err)
		return nil
	}

	return &Adapter{client: minioClient, cfg: cfg}
}

func (a Adapter) Conn() *minio.Client {
	return a.client
}
