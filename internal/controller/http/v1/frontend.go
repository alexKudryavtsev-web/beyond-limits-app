// internal/controller/http/frontend/v1/frontend.go
package v1

import (
	"strconv"

	"github.com/alexKudryavtsev-web/beyond-limits-app/internal/usecase"
	"github.com/alexKudryavtsev-web/beyond-limits-app/pkg/logger"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

type frontendRoutes struct {
	picturesUC usecase.Pictures
	refUC      usecase.References
	l          logger.Interface
}

func NewFrontendRouter(
	handler *gin.Engine,
	logger logger.Interface,
	picturesUC usecase.Pictures,
	referencesUC usecase.References,
) {
	r := &frontendRoutes{
		picturesUC: picturesUC,
		refUC:      referencesUC,
		l:          logger,
	}

	handler.HTMLRender = r.createRenderer()

	handler.Static("/static", "./web/static")
	handler.Static("/uploads", "./uploads")

	handler.GET("/", r.homePage)
	handler.GET("/pictures/:id", r.picturePage)
}

func (r *frontendRoutes) createRenderer() multitemplate.Renderer {
	renderer := multitemplate.NewRenderer()

	renderer.AddFromFiles("home",
		"web/templates/base.html",
		"web/templates/home.html")

	renderer.AddFromFiles("picture",
		"web/templates/base.html",
		"web/templates/picture.html")

	return renderer
}

func (r *frontendRoutes) homePage(c *gin.Context) {
	pictures, err := r.picturesUC.GetPictures(c.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - homePage - get pictures")
	}

	genres, err := r.refUC.GetGenres(c.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - homePage - get genres")
	}

	c.HTML(200, "home", gin.H{
		"Title":    "Каталог картин",
		"Pictures": pictures,
		"Genres":   genres,
	})
}

func (r *frontendRoutes) picturePage(c *gin.Context) {
	id := c.Param("id")
	pictureID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		r.l.Error(err, "http - v1 - picturePage - parse id")
		c.AbortWithStatus(400)
		return
	}

	picture, err := r.picturesUC.GetPictureByID(c.Request.Context(), pictureID)
	if err != nil {
		r.l.Error(err, "http - v1 - picturePage - get picture")
		c.AbortWithStatus(404)
		return
	}

	c.HTML(200, "picture", gin.H{
		"Title":   picture.Title,
		"Picture": picture,
	})
}
