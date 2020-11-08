package controls

import (
	"fmt"
	"ginApi/config"
	"ginApi/models"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

var db = config.Config()

// @Summary 获取音乐列表
// @Description 音乐列表
// @Tags 音乐
// @Accept json
// @Produce json
// @Param name query string false "name"
// @Param pageNum query string true "pageNum"
// @Param pageSize query string true "pageSize"
// @Success 200 {object} models.Music
// @Failure 400 {string} json "{ "code": 400, "message": "请求失败" }"
// @Failure 500 {string} json "{ "code": 500, "message": "服务器错误" }"
// @Router /api/v2/music/ [get]
func MusicList(c *gin.Context) {
	var musics []models.Music
	fmt.Println("ip", c.ClientIP()) // 客户ip

	name := c.Query("name")
	pageNum := com.StrTo(c.DefaultQuery("pageNum", "1")).MustInt()    // 设置pageNum的默认参数 1
	pageSize := com.StrTo(c.DefaultQuery("pageSize", "10")).MustInt() // 设置pageSize默认参数 2

	fmt.Println(name, pageNum, pageSize, "-*-")
	/*
		迷糊搜索，name为搜索的条件，根据电影的名称name来搜索
		Offset 其实条数
		Limit	 每页的条数
		Order("id desc") 根据id倒序排序
		总条数 Count(&count)
	*/

	var count int
	db.Model(&musics).Where("name LiKE?", "%"+name+"%").Count(&count)
	db.Offset((pageNum-1)*pageSize).Limit(pageSize).Order("id desc").Where("name LiKE?", "%"+name+"%").Find(&musics)

	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"status":  http.StatusOK,
		"data":    musics,
		"ip":      c.ClientIP(),
		"attributes": gin.H{
			"page":  pageNum,
			"total": count,
		},
	})
}

// @Summary 创建音乐列表
// @Description 创建音乐
// @Tags 音乐
// @Accept json
// @Produce json
// @Param name query string false "name"
// @Param year query string false "year"
// @Param style query string false "style"
// @Success 200 {object} models.Music
// @Failure 400 {string} json "{ "code": 400, "message": "请求失败" }"
// @Router /api/v2/music/ [post]
func MusicCreate(c *gin.Context) {
	/*
	 gin还提供了更加高级方法，c.Bind，
	 它会更加content-type自动推断是bind表单还是json的参数
	 json格式application/json或者表单格式x-www-form-urlencoded
	*/
	data := &models.Music{}
	err := c.Bind(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data, "--")
	db.Create(data)
	db.Save(data)
	c.JSON(http.StatusOK, gin.H{
		"message": "创建成功",
		"status":  http.StatusOK,
		"data":    data,
	})
}

// @Summary 更新音乐列表
// @Description 音乐列表
// @Tags 音乐
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Music
// @Failure 400 {string} json "{ "code": 400, "message": "id必传" }"
// @Router /api/v2/music/{id} [put]
func MusicUpdate(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	fmt.Println(id, "--")

	// 需要更新的元素
	data := &models.Music{}
	err := c.Bind(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 根据id更新对应的数据
	db.Model(data).Where("id=?", id).Update(data)
	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
		"status":  http.StatusOK,
		"data":    data,
	})
}

// @Summary 删除音乐列表
// @Description 音乐列表
// @Tags 音乐
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} models.Music
// @Failure 400 {string} json "{ "code": 400, "message": "id必传" }"
// @Router /api/v2/music/{id} [delete]
func MusicDelete(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	fmt.Println("id")

	db.Where("id=?", id).Delete(models.Music{})
	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
		"status":  http.StatusOK,
		"data":    id,
	})
}
