package fontends

import (
	"fmt"
	"github.com/midoks/novelsearch/app/models"
	"strings"
)

type IndexController struct {
	CommonController
}

func (this *IndexController) Index() {
	this.display()
}

func (this *IndexController) Top() {
	this.display()
}

func (this *IndexController) Soso() {
	kw := this.GetString("wd")
	if strings.EqualFold(kw, "") {
		return
	}
	fmt.Println(kw)
	list := models.SosoNovelByKw(kw)
	fmt.Println(list)
	this.display()
}

func (this *IndexController) Page() {
	this.display()
}

func (this *IndexController) Details() {
	this.display()
}
