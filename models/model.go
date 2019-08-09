package models

import "time"

type BasicModel struct {
	ID        int       `json:"id,omitempty"`
	CreatedAt time.Time `json:"createAt" example:"创建时间"`
	UpdatedAt time.Time `json:"updateAt" example:"更新时间"`
}

type Music struct {
	BasicModel
	Name  string `json:"name" example:"歌曲名称"`
	Year  string `json:"year" example:"年份"`
	Style string `json:"style" example:"音乐风格"`
}

type Film struct {
	BasicModel
	Name    string `json:"name" example:"电影名称"`
	Year    string `json:"year" example:"年份"`
	Address string `json:"address" example:"出品地区"`
	Actor   string `json:"actor" example:"演员"`
	Desc    string `json:"desc" example:"描述"`
}
