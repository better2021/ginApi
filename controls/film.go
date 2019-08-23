package controls

import (
	"fmt"
	"ginApi/models"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

// var db = config.Config()

// @Summary 获取电影列表
// @Description 电影列表
// @Tags 电影
// @Accept json
// @Produce json
// @Param name query string false "name"
// @Param pageNum query string true "pageNum"
// @Param pageSize query string true "pageSize"
// @Success 200 {object} models.Film
// @Failure 400 {string} json "{ "code": 400, "message": "请求失败" }"
// @Router /api/v2/film/ [get]
func FilmList(c *gin.Context) {
	var films []models.Film

	name := c.Query("name")
	pageNum := com.StrTo(c.DefaultQuery("pageNum", "1")).MustInt()
	pageSize := com.StrTo(c.DefaultQuery("pageSize", "10")).MustInt()

	fmt.Print(name, pageNum, "---")
	var count int // 总条数
	db.Model(&films).Where("name LiKE?", "%"+name+"%").Count(&count)
	db.Offset((pageNum-1)*pageSize).Limit(pageSize).Order("created_at desc").Where("name LIKE?", "%"+name+"%").Find(&films)

	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"status":  http.StatusOK,
		"data":    films,
		"page":    pageNum,
		"total":   count,
	})
}

// @Summary 创建电影列表
// @Description 创建电影
// @Tags 电影
// @Accept json
// @Produce json
// @Param name query string false "name"
// @Param year query string false "year"
// @Param address query string false "address"
// @Param actor query string false "actor"
// @Param desc query string false "desc"
// @Success 200 {object} models.Film
// @Failure 400 {string} json "{ "code": 400, "message": "请求失败" }"
// @Router /api/v2/film/ [post]
func FilmCreate(c *gin.Context) {
	var data = &models.Film{}
	err := c.Bind(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(data, "--")
	db.Create(data)
	c.JSON(http.StatusOK, gin.H{
		"message": "创建成功",
		"status":  http.StatusOK,
		"data":    data,
	})
}

// @Summary 更新电影列表
// @Description 电影列表
// @Tags 电影
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Film
// @Failure 400 {string} json "{ "code": 400, "message": "id必传" }"
// @Router /api/v2/film/{id} [put]
func FilmUpdate(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	fmt.Println(id, "--")

	condition := &models.BasicModel{ID: id}
	data := &models.Film{}
	err := c.Bind(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	db.Model(condition).Update(data)
	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
		"status":  http.StatusOK,
		"data":    data,
	})
}

// @Summary 删除电影列表
// @Description 电影列表
// @Tags 电影
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Film
// @Failure 400 {string} json "{ "code": 400, "message": "id必传" }"
// @Router /api/v2/film/{id} [delete]
func FilmDelete(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	fmt.Println("id", "--")

	db.Where("id=?", id).Delete(models.Film{})
	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
		"status":  http.StatusOK,
		"data":    id,
	})
}
