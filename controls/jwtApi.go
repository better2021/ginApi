package controls

import (
	"fmt"
	"ginApi/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var jwtSecret = []byte("deng")

func generateToken(username string) (string, error) {
	nowTime := time.Now()
	// 设置有效时间
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "AustinDeng",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

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
	data := make(map[string]interface{})
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
		token, err := generateToken(username)
		if err != nil {
			fmt.Println("生成token时出错")
		} else {
			data["token"] = token
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"code":    1,
			"message": "登录成功！",
			"data":    data,
		})
	}
}

// 用户列表
func UserList(c *gin.Context) {
	var registerInfos []models.RegistInfo
	db.Find(&registerInfos)
	fmt.Println(registerInfos, "--")
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"status":  http.StatusOK,
		"data":    registerInfos,
	})
}

// 删除用户
func UserDelete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		panic(err)
	}

	fmt.Println(id, "--")
	db.Where("id=?", id).Delete(models.RegistInfo{})
	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
		"status":  http.StatusOK,
	})
}
