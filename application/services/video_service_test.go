package services_test

import (
	"log"
	"testing"
	"time"

	"github.com/Lucasvmarangoni/video-encoder/application/repositories"
	"github.com/Lucasvmarangoni/video-encoder/application/services"
	"github.com/Lucasvmarangoni/video-encoder/domain"
	"github.com/Lucasvmarangoni/video-encoder/framework/database"
	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func prepare() (domain.Video, repositories.VideoRepositoryDb) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "currency-converter-demo.mp4"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	// repo.Insert(video)

	return *video, repo
}

func TestVideoServiceDownload(t *testing.T) {
	video, repo := prepare()

	videoService := services.NewVideoService()
	videoService.Video = &video
	videoService.VideoRepository = repo

	err := videoService.Download("encoder-bucket-01")
	require.Nil(t, err)

	err = videoService.Fragment()
	require.Nil(t, err)

	err = videoService.Encode()
	require.Nil(t, err)

	err = videoService.Finish()
	require.Nil(t, err)
}