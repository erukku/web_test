package main
import(
	"log"
	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()

	router.GET("/",func(ctx *gin.Context){
		ctx.String(200,"Hello, World!")
	})

	err :=router.Run("127.0.0.1:8888")
	if err != nil{
		log.Fatal("fail",err)
	}
}
