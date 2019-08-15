package controls

import (
	"fmt"
	"ginApi/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary 获取书籍列表
// @Description 书籍列表
// @Tags 书籍
// @Accept json
// @Produce json
// @Param name query string false "name"
// @Param pageNum query string true "pageNum"
// @Param pageSize query string true "pageSize"
// @Success 200 {object} models.Book
// @Failure 400 {string} json "{ "code": 400, "message": "请求失败" }"
// @Router /api/v2/book/ [get]
func BookList(c *gin.Context) {
	var books []models.Book

	name := c.Query("name")

	pageNum, err := strconv.Atoi(c.Query("pageNum"))
	if err != nil {
		fmt.Println("请输入分页pageNum")
		c.JSON(http.StatusOK, gin.H{
			"message": "请输入分页pageNum",
		})
		return
	}
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	if err != nil {
		fmt.Println("请输入分页pageSize")
		c.JSON(http.StatusOK, gin.H{
			"message": "请输入分页pageSize",
		})
		return
	}

	fmt.Println(name, pageNum, "---")
	var count int
	db.Model(&books).Where("name LIKE?", "%"+name+"%").Count(&count)
	db.Offset((pageNum-1)*pageSize).Limit(pageSize).Order("created_at desc").Where("name LIKE?", "%"+name+"%").Find(&books)

	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"status":  http.StatusOK,
		"data":    books,
		"page":    pageNum,
		"total":   count,
	})
}

// @Summary 创建书籍列表
// @Description 创建书籍
// @Tags 书籍
// @Accept json
// @Produce json
// @Param name query string false "name"
// @Param year query string false "year"
// @Param author query string false "author"
// @Param desc query string false "desc"
// @Success 200 {object} models.Book
// @Failure 400 {string} json "{ "code": 400, "message": "请求失败" }"
// @Router /api/v2/book/ [post]
func BookCreate(c *gin.Context) {
	var data = &models.Book{}
	err := c.BindJSON(data) // bindjson 后可以获取传的json数据
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

// @Summary 更新书籍列表
// @Tags 书籍
// @Accept json
// @Produce json
// @Param name query string false "name"
// @Param year query string false "year"
// @Param author query string false "author"
// @Param desc query string false "desc"
// @Success 200 {object} models.Book
// @Failure 400 {string} json "{ "code": 400, "message": "请求失败" }"
// @Router /api/v2/book/{id} [put]
func BookUpdate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id, "--")

	condition := &models.BasicModel{ID: id}
	data := &models.Book{}
	err := c.BindJSON(data)
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

// @Summary 删除书籍列表
// @Tags 书籍
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Book
// @Failure 400 {string} json "{ "code": 400, "message": "请求失败" }"
// @Router /api/v2/book/{id} [delete]
func BookDelete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id, "--")

	db.Where("id=?", id).Delete(models.Book{})
	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
		"status":  http.StatusOK,
		"data":    id,
	})
}
