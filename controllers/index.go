package controllers

import "phpsong/models"

type IndexHandle struct {
	baseController
}

func (this *IndexHandle) Start() {
	this.TplName = "index.html"
}

func (this *IndexHandle) Index() {
	var (
		info models.PostsInfo
	)
	info.GetList()
	this.TplName = "index.html"
}
