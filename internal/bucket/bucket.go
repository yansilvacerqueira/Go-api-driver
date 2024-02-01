package bucket

import (
	"io"
	"os"
	"reflect"
  "fmt"
)

const (
  S3 BucketType = iota
)

type BucketType int

type BucketInterface interface {
  UploadFile(io.Reader, string) error
  DownloadFile(string, string) (*os.File,error)
  DeleteFile(string) error
}

type Bucket struct {
  provider BucketInterface
}

func NewBucket(bucketType BucketType, config any) (*Bucket, error) {
  rt := reflect.TypeOf(config)

  switch bucketType {
  case S3:
    if rt.Name() == "S3" {
      // TODO: Implement S3 provider
    }
  default:
    return nil, fmt.Errorf("Invalid bucket type")
  }

  //TODO: Fix this
  return nil, nil 
}

func (b *Bucket) UploadFile(file io.Reader, key string) error {
  return b.provider.UploadFile(file, key)
}

func (b *Bucket) DownloadFile(source string, destination string) (*os.File, error) {
  return b.provider.DownloadFile(source, destination)
}

func (b *Bucket) DeleteFile(key string) error {
  return b.provider.DeleteFile(key)
}

