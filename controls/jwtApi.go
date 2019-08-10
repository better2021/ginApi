package controls

import (
	"fmt"
	"ginApi/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

// Register 注册用户
func RegisterUser(c *gin.Context) {
	var registerInfo = &models.RegistInfo{}
	_ = c.BindJSON(registerInfo)

	fmt.Println(registerInfo.Username, "+++")
	err := db.Where(models.RegistInfo{Username: registerInfo.Username}).First(registerInfo).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		db.Create(registerInfo)
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"msg":    "注册成功！",
			"data":   registerInfo,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "用户名已注册！",
		})
	}
}

// Login 登录
func Login(c *gin.Context) {
	var registerInfo = &models.RegistInfo{}
	username := c.PostForm("username") // 表单上传
	pwd := c.PostForm("pwd")

	fmt.Println(username, pwd, "---")
	err := db.Where(models.RegistInfo{Username: username}).First(registerInfo).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"code":    0,
			"message": "用户名不存在！",
		})
		return
	}

	if registerInfo.Pwd != pwd {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"code":    -1,
			"message": "密码错误！",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"code":    1,
			"message": "登录成功！",
		})
	}
}
