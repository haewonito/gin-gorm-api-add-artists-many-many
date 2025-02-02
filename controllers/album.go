package controller
//https://gorm.io/docs/query.html
//https://gorm.io/docs/create.html
import (
	"github.com/gin-gonic/gin"
	"github.com/haewonito/gin-gorm-api-add-artists-many-many/models"
	"github.com/haewonito/gin-gorm-api-add-artists-many-many/config"
)

func GetAlbums(c *gin.Context) {
	albums := []models.Album{}
	config.DB.Find(&albums)
	c.JSON(200, &albums)
}

func CreateAlbum(c *gin.Context) {
	var album models.Album
	c.BindJSON(&album)
	config.DB.Create(&album)
	c.JSON(200, &album)
}

func GetAlbumById(c *gin.Context) {
	var album models.Album
	config.DB.First(&album, c.Param("id"))
	c.JSON(200, &album)
}

//todo haewon - need to delete related songs
func DeleteAlbum(c *gin.Context) {
	var album models.Album

	if err := config.DB.Preload("Songs").First(&album, c.Param("id")).Error; err != nil {
		panic(err)
	}

	var songs []models.Song
	for _, song := range album.Songs {
		songs = append(songs, song)
	}
	config.DB.Where("id = ?", c.Param("id")).Delete(&album)
	config.DB.Delete(&songs)

	c.JSON(200, &album)
}

func UpdateAlbum(c *gin.Context) {
	var album models.Album
	config.DB.Where("id = ?", c.Param("id")).First(&album)
	c.BindJSON(&album)
	config.DB.Save(&album)
	c.JSON(200, &album)
}