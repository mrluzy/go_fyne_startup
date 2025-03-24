package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"time"
)

func main() {
	a := app.NewWithID("01")
	w := a.NewWindow("高端检测器")

	button := widget.NewButton("检验是否开机", func() {
		progress := widget.NewProgressBarInfinite()
		progressContainer := container.NewVBox(progress)

		loadingDialog := dialog.NewCustom("加载中...", "取消", progressContainer, w)
		loadingDialog.Show()

		go func() {
			time.Sleep(10 * time.Second)
			loadingDialog.Hide()
			dialog.ShowInformation("检测结果", "电脑是开机", w)

		}()
	})

	w.SetContent(container.NewVBox(button))
	w.Resize(fyne.NewSize(400, 400))
	w.CenterOnScreen()
	w.ShowAndRun()
}
