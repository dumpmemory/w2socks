package main

import (
	"github.com/dreamlu/w2socks/client/data"
	"github.com/dreamlu/w2socks/client/gui"
	"github.com/getlantern/systray"
)

// 开发环境: ubuntu
// 安装依赖: sudo apt-get install libgl1-mesa-dev xorg-dev libgtk-3-dev libappindicator3-dev -y

// 运行方式:
// 1.命令行
// 2.GUI
func main() {
	gui.G = gui.Gui()
	gui.G.Show()
	go systray.Run(onReady, nil)
	//w.SetOnClosed(func() {
	//	w = window()
	//	//w.ShowAndRun()
	//	//systray.Run(onReady, nil)
	//})
	gui.G.ShowAndRun()
}

// 驻后台
func onReady() {
	systray.SetTemplateIcon(data.LogoData, data.LogoData)
	//systray.SetTitle("w2socks")
	systray.SetTooltip("w2socks")
	// 托盘菜单
	mUrl := systray.AddMenuItem("恢复", "my home")
	mQuit := systray.AddMenuItem("退出", "Quit the whole app")
	//systray.AddSeparator() // 分隔线
	for {
		select {
		case <-mUrl.ClickedCh:
			gui.G.Show()
			//o <- 0
		case <-mQuit.ClickedCh:
			systray.Quit()
			return
		}
	}
}
