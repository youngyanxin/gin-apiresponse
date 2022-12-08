package gin_apiresponse

import (
	"github.com/gin-gonic/gin"
	"testing"
)

//func main() {
//	g := gin.Default()
//	Test(g)
//	g.Run(":9999")
//}

func main(ginCtx *gin.Engine) {

	//New().Success(ginCtx, 1)

}

func Test(t *testing.T) {
	gCtx := &gin.Context{}
	Success(gCtx, 1)
}
