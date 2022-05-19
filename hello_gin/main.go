package main
import(
	"os"
	"io"
	"fmt"
	"bufio"
	"log"
	"net"
	"math"
	"strconv"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	//"net/http"
)

type Todo struct {
	gorm.Model
	Memo string
}

type Colors struct {
	Color string
}

func udpGet(){
	txtName := "id.txt"

	udp := &net.UDPAddr{
		IP:   net.ParseIP("localhost"),
		Port: 8081,
	}
	updLn, err2 := net.ListenUDP("udp", udp)
	if err2 != nil{
		log.Fatal("fail",err2)
	}

	buf := make([]byte, 1024)
	log.Println("Starting UDP Server...")

	for {
		n, addr, err := updLn.ReadFromUDP(buf)
		if err != nil {
			log.Fatalln(err)
		}
		data := string(buf[:n])

		dataS := strings.Split(data,",")
		if len(dataS) != 0 && dataS[0] != ""{

		}
		if err:= writerow(txtName,dataS); err != nil {
			log.Fatalln(err)
		}

		go func() {
			log.Printf("From: %v Reciving data: %s", addr.String(), data)
		}()
	}
}

func readIds(txt string,colors *[]Colors) error {
	file,err := os.Open(txt)
	if err != nil{
		return err
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, _ ,err := reader.ReadLine()
		if err == io.EOF{
			break
		}
		if err != nil {
			return err
		}

		str := string(line)
		colr, err := strconv.Atoi(str)
		if err != nil{
			return err
		}
		mod := int(math.Pow(16,6))
		colr = colr % mod
		color := fmt.Sprintf("%06x", colr)
		color = "#"+color
		*colors = append(*colors,Colors{
			Color: color,
		})

	}

	return nil
}

func writerow(txt string,idlist []string) error{
	file,err := os.Create(txt)
	if err != nil{
		return err
	}

	defer file.Close()

	for _,id := range idlist{
		_,err := file.WriteString(id+"\n")
		if err != nil{
			return err
		}
	}
	return nil
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
	router.GET("/next.html",getHandler2)
	log.Printf("aa")
	router.POST("/new",postHandler)
	go udpGet()
	router.Run()
	
	err :=router.Run(":8888")
	if err != nil{
		log.Fatal("fail",err)
	}
}


func getHandler(c *gin.Context){
	todo := getAll()
	c.HTML(200,"index.html",gin.H{"todo":todo})
}

func getHandler2(c *gin.Context){
	//todo := getAll()
	//c.HTML(200,"next.html",gin.H{"todo":todo})

	var colorlist []Colors
	txt := "id.txt"
	if err := readIds(txt,&colorlist);err != nil{
		log.Fatal(err)
	}
	size := len(colorlist)
	c.HTML(200,"next.html",gin.H{"Size":size,"Colors":colorlist})
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
	return todo[max(0,len(todo)-4):]
}

func max (a int,b int) int {
	if a > b{
		return a
	}else{
		return b
	}
}
