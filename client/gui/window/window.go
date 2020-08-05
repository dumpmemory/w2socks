package window

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"github.com/dreamlu/w2socks/client/data"
	"github.com/dreamlu/w2socks/client/gui/global"
	"github.com/dreamlu/w2socks/client/util/notify"
	"log"
)

// 通用window
// 编辑/添加连接窗体
func OpenWindow(conf *data.Config, add bool) fyne.Window {
	w := fyne.CurrentApp().NewWindow("connect content")
	w.Resize(fyne.NewSize(280, 300))
	comSize := fyne.NewSize(100, 20)

	var id string

	// m名字
	nameEntry := widget.NewEntry()
	// 服务端ip和端口
	serverEntry := widget.NewEntry()
	// 本地端口号
	localPortEntry := widget.NewEntry()
	if add {
		// 添加
		nameEntry.SetPlaceHolder("name")
		serverEntry.SetPlaceHolder("ip:port")
		localPortEntry.SetPlaceHolder("port")
	} else {
		id = conf.ID
		nameEntry.Text = conf.Name
		serverEntry.Text = conf.ServerIpAddr
		localPortEntry.Text = conf.LocalPort
	}
	nameEntry.Resize(comSize)
	serverEntry.Resize(comSize)
	localPortEntry.Resize(comSize)

	form := widget.NewForm(
		widget.NewFormItem("name:", nameEntry),
		widget.NewFormItem("server:", serverEntry),
		widget.NewFormItem("local:", localPortEntry),
	)

	form.CancelText = "cancel"
	form.SubmitText = "save"

	// 取消操作
	form.OnCancel = func() {
		w.Hide()
	}

	// 连接操作
	form.OnSubmit = func() {
		log.Println("提交")
		// 检查验证输入的ip地址是否有效
		b := CheckEntry(serverEntry.Text, localPortEntry.Text)
		if !b {
			return
		}
		conf := data.Config{
			ID:           id,
			Name:         nameEntry.Text,
			ServerIpAddr: serverEntry.Text,
			LocalPort:    localPortEntry.Text,
		}

		var err error
		if add {
			// 添加
			err = data.InsertConfig(conf)
		} else {
			// 编辑
			err = data.UpdateConfig(conf)
		}
		if err != nil {
			notify.SysNotify("warn!!", err.Error())
			return
		}
		notify.SysNotify("info", "连接信息已存入")
		global.G.Refresh <- 1
		w.Close()
	}

	// 窗体
	content := widget.NewVBox(
		widget.NewVBox(
			// 输入服务端的ip地址和端口 以及本地的端口
			widget.NewLabel("Please Enter:"),
			form,
		),
	)
	w.SetContent(content)
	w.SetOnClosed(func() {
		fmt.Println("操作完成,刷新")
		global.G.Refresh <- 1
	})
	w.Show()
	return w
}
