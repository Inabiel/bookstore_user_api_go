package app

import "github.com/gin-gonic/gin"

var (
	r = gin.Default()
)

func StartApplication() {
	MapUrls()
	r.Run(":9005")
}
