package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MainController struct{}

func (con MainController) Index(c *gin.Context) {
	// //获取userinfo 对应的session
	// session := sessions.Default(c)
	// userinfo := session.Get("userinfo")
	// //类型断言 来判断 userinfo是不是一个string
	// userinfoStr, ok := userinfo.(string)

	// if ok {
	// 	var userinfoStruct []models.Manager
	// 	json.Unmarshal([]byte(userinfoStr), &userinfoStruct)
	// 	fmt.Println(userinfoStruct)
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"username": userinfoStruct[0].Username,
	// 	})
	// } else {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"username": "session不存在",
	// 	})
	// }

	c.HTML(http.StatusOK, "admin/main/index.html", gin.H{})
}

func (con MainController) Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/main/welcome.html", gin.H{})
}
