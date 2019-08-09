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

// 获取音乐列表
func MusicList(c *gin.Context) {
	var musics []models.Music

	name := c.Query("name")
	pageNum := com.StrTo(c.Query("pageNum")).MustInt()
	pageSize := com.StrTo(c.Query("pageSize")).MustInt()

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
		"page":    pageNum,
		"total":   count,
	})
}

// 创建音乐列表
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

// 更新列表
func MusicUpdate(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	fmt.Println(id, "--")

	music := &models.Music{ID: id} // 修改条件，根据ID修改
	// 需要更新的元素
	data := &models.Music{}
	err := c.Bind(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	db.Model(music).Update(data)
	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
		"status":  http.StatusOK,
		"data":    data,
	})
}

// 删除列表
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

