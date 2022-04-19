package main
import(
	"log"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	//"net/http"
)

type Todo struct {
	gorm.Model
	Memo string
}

func main(){
	router := gin.Default()
	router.Static("styles", "./styles")
	router.Static("css", "./css")
	router.Static("js", "./js")
	router.Static("image", "./image")
	router.LoadHTMLGlob("templates/*")
	dbInit()
	

	router.GET("/",getHandler)
	router.POST("/new",postHandler)
	router.Run()

	err :=router.Run("127.0.0.1:8888")
	if err != nil{
		log.Fatal("fail",err)
	}
}


func getHandler(c *gin.Context){
	todo := getAll()
	c.HTML(200,"index.html",gin.H{"todo":todo})
}

func postHandler(c *gin.Context){
	memo := c.PostForm("memo")
	create(memo)
	c.Redirect(302,"/")
}

func dbInit(){
	db,err := gorm.Open("sqlite3","todo.splite3")
	if err != nil {
		panic("fail to connect db\n")
	}
	db.AutoMigrate(&Todo{})
}


func create(memo string){
	db,err := gorm.Open("sqlite3","todo.splite3")
	if err != nil {

		panic("fail to connect db\n")
	}
	db.Create(&Todo{Memo:memo})
}

func getAll() []Todo {
	db,err := gorm.Open("sqlite3","todo.splite3")
	if err != nil {
		panic("fail to connect db\n")
	}
	var todo []Todo
	db.Find(&todo)
	return todo
}
