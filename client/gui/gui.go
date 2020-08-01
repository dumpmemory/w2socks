package gui

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/dreamlu/w2socks/client/data"
	text2 "github.com/dreamlu/w2socks/client/gui/cuscom/text"
)

var (
	G fyne.Window
)

// 主窗体
func Gui() fyne.Window {
	// 主程序
	majorApp := app.New()
	// light theme
	majorApp.Settings().SetTheme(theme.LightTheme())
	// logo
	majorApp.SetIcon(data.Logo())
	addWindow := majorApp.NewWindow("w2socks")
	addWindow.Resize(fyne.NewSize(280, 300))

	// 主菜单
	addWindow.SetMainMenu(MainMenu())
	// 布局
	//top := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), Toolbar())
	//bom := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), list.Content)
	vert := widget.NewVScrollContainer(widget.NewVBox(mainList()...))
	addWindow.SetContent(fyne.NewContainerWithLayout(layout.NewAdaptiveGridLayout(1), vert))
	return addWindow
}

// 列表核心
func mainList() []fyne.CanvasObject {
	var items []fyne.CanvasObject

	conf := data.GetConfig()
	for _, v := range conf {
		item :=
			NewLine(
				text2.NewSelectClickText(fmt.Sprintf("%s", v.Name), *v),
				text2.NewSelectClickText(fmt.Sprintf("%s", v.ServerIpAddr), *v),
				//NewSelectClickText(fmt.Sprintf("%s", v.LocalPort), *v),
			)
		items = append(items, item)
	}
	return items
}

// new line
//func NewLine(w ...fyne.Widget) fyne.Widget {
//	w := widget.NewHBox()
//	return w
//}

// new line box
func NewLine(obj ...fyne.CanvasObject) fyne.CanvasObject {

	c := widget.NewVBox(obj...)
	//c := fyne.NewContainerWithLayout(layout.NewAdaptiveGridLayout(3), obj...)

	return c
}
