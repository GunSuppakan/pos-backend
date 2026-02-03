package storage

import (
	"log"
	"mime/multipart"

	"github.com/nedpals/supabase-go"
	"github.com/spf13/viper"
)

type StorageRepository interface {
	SaveFile(file *multipart.FileHeader, path string) error
	DeleteFile(path string) error
}

type storageRepository struct {
	client     *supabase.Client
	bucketName string
}

func NewStorageRepository() StorageRepository {
	client := supabase.CreateClient(
		viper.GetString("supabase.url"),
		viper.GetString("supabase.service_role_key"),
	)

	return &storageRepository{
		client:     client,
		bucketName: viper.GetString("supabase.bucket"),
	}
}

func (s *storageRepository) SaveFile(file *multipart.FileHeader, path string) error {
	log.Println("[Storage] SaveFile called")
	log.Println("[Storage] bucket:", s.bucketName)
	log.Println("[Storage] path:", path)

	src, err := file.Open()
	if err != nil {
		log.Println("[Storage] open file error:", err)
		return err
	}
	defer src.Close()

	log.Println("[Storage] filename:", file.Filename)
	log.Println("[Storage] size:", file.Size)
	log.Println("[Storage] content-type:", file.Header.Get("Content-Type"))

	defer func() {
		if r := recover(); r != nil {
			log.Printf("[Storage] panic during upload: %v\n", r)
		}
	}()

	resp := s.client.Storage.
		From(s.bucketName).
		Upload(
			path,
			src,
			&supabase.FileUploadOptions{
				ContentType: file.Header.Get("Content-Type"),
				Upsert:      false,
			},
		)

	log.Printf("[Storage] upload response: %+v\n", resp)

	return nil
}

// func (s *storageRepository) SaveFile(file *multipart.FileHeader, path string) error {
// 	src, err := file.Open()
// 	if err != nil {
// 		return err
// 	}
// 	defer src.Close()

// 	s.client.Storage.From(s.bucketName).Upload(path, src, &supabase.FileUploadOptions{
// 		ContentType: file.Header.Get("Content-Type"),
// 		Upsert:      false,
// 	},
// 	)

// 	return nil
// }

func (s *storageRepository) DeleteFile(path string) error {
	s.client.Storage.
		From(s.bucketName).
		Remove([]string{path})

	return nil
}
