package controllers

type IndexHandle struct {
	baseController
}

func (this *IndexHandle) Start() {
	this.TplName = "index.html"
}

func (this *IndexHandle) Index() {
	var (
		info models.PostInfo
	)
	this.TplName = "index.html"
}
