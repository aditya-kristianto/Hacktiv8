package gin

import (
	"sesi6/gin/router"
)

func Start() {

	router.CreateRouter().Run(":4000")
}
