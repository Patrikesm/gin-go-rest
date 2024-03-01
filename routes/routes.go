package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/patrike-miranda/gin-go-rest/controllers"
)

func HandleRequests() {
	r := gin.Default()              //default config for gin
	r.LoadHTMLGlob("templates/*")   // realiza a renderização de todos os templates existentes
	r.Static("/assets", "./assets") // realiza o load dos arquivos estaticos

	r.GET("/:name", controllers.Welcome)
	r.GET("/alunos", controllers.GetAll)
	r.GET("/alunos/:id", controllers.GetOne)

	r.PATCH("/alunos/:id", controllers.Update)

	r.POST("/alunos", controllers.New)

	r.DELETE("/alunos/:id", controllers.Delete)

	r.GET("/alunos/document/:document", controllers.GetByDocument)

	r.GET("/index", controllers.RenderizeIndexPage)
	r.NoRoute(controllers.NotFoundRoute)

	//diferente do http padrão aqui temos apenas um run e não um listen and serve
	r.Run(":3000")
}
