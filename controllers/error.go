package controllers

type ErrorController struct {
	baseController
}

func (this *ErrorController) Error404() {
	this.Data["content"] = "page not found"
	this.TplName = "404.html"
}
