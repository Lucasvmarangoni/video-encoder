package repositories_test

import (
	"testing"

	"github.com/Lucasvmarangoni/video-encoder/application/repositories"
	"github.com/Lucasvmarangoni/video-encoder/domain"
	"github.com/Lucasvmarangoni/video-encoder/framework/database"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestVideoRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"

	repo := repositories.VideoRepositoryDb{Db:db}
	repo.Insert(video)

	v, err := repo.Find(video.ID)

	require.NotEmpty(t, v.ID)
	require.Nil(t, err)
	require.Equal(t, v.ID, video.ID)
}
