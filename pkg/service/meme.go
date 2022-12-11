package service

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	memeConfig "github.com/zhews/memed-simple/pkg/config/meme"
	"github.com/zhews/memed-simple/pkg/domain"
	"io"
	"log"
	"os"
	"time"
)

type MemeService struct {
	Config     memeConfig.Config
	Repository domain.MemeRepository
}

func (ms *MemeService) GetMemes() ([]domain.Meme, error) {
	memes, err := ms.Repository.GetAll()
	return memes, err
}

func (ms *MemeService) GetMemeById(id uuid.UUID) (domain.Meme, error) {
	meme, err := ms.Repository.GetById(id)
	return meme, err
}

var contentTypeExtensions = map[string]string{
	"image/jpeg": "jpeg",
	"image/png":  "png",
}

var ErrorInvalidContentType = errors.New("invalid content type")

func (ms *MemeService) UploadMeme(title, contentType string, file io.Reader, userId uuid.UUID) error {
	extension, ok := contentTypeExtensions[contentType]
	if !ok {
		return ErrorInvalidContentType
	}
	imageFileName := fmt.Sprintf("%s.%s", uuid.New().String(), extension)
	imageFile := fmt.Sprintf("%s/%s", ms.Config.MemeDirectory, imageFileName)
	err := writeMemeToDisk(imageFile, file)
	if err != nil {
		return err
	}
	now := time.Now().Unix()
	meme := domain.Meme{
		Id:        uuid.New(),
		Title:     title,
		Image:     imageFileName,
		CreatedBy: userId,
		CreatedAt: now,
		UpdatedAt: now,
	}
	err = ms.Repository.Insert(meme)
	return err
}

func writeMemeToDisk(fileName string, meme io.Reader) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println("Could not close file: ", err)
		}
	}(file)
	buffer := make([]byte, 1024)
	for {
		n, err := meme.Read(buffer)
		if err != nil && !errors.Is(err, io.EOF) {
			return err
		}
		if n == 0 {
			break
		}
		if _, err = file.Write(buffer); err != nil {
			return err
		}
	}
	return nil
}

func (ms *MemeService) UpdateMeme(id uuid.UUID, title string, userId uuid.UUID) error {
	now := time.Now().Unix()
	meme := domain.Meme{
		Id:        id,
		Title:     title,
		UpdatedAt: now,
	}
	err := ms.Repository.Update(meme, userId)
	return err
}

func (ms *MemeService) DeleteMeme(id uuid.UUID) error {
	err := ms.Repository.Delete(id)
	return err
}
