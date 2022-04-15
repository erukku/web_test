package main
import(
	"log"
	"github.com/gin-gonic/gin"
	//"net/http"
)

func main(){
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	ua := ""

	router.Use(func(c *gin.Context){
		ua = c.GetHeader("User-Agent")
		c.Next()
	})

	router.GET("/",func(ctx *gin.Context){
		ctx.HTML(200,"index.html",gin.H{
			"message": "hello world",
			"User-Agent": ua,
		})
		
	})

	err :=router.Run("127.0.0.1:8888")
	if err != nil{
		log.Fatal("fail",err)
	}
}
