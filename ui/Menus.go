package ui

import (
	"IP_pkg_analyze/ip"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func loadMenus(w fyne.Window) {
	tools_key := []string{"文件(F)"} // "过滤(E)", "排序(V)", "跳转(G)",
	//"捕获(C)", "分析(A)", "统计(S)", "电话(Y)",
	//"无线(W)", "工具(T)", "帮助(H)",

	var tools = [][]string{{"保存"}}
	var toolsFunc = [][]func(){{
		func() { //保存
			path, err := ip.SaveAsPcap(ip.PkgInfos)
			if err == nil {
				dialog.NewInformation("提示", "成功保存至 "+path, w).Show()
			}
		},
	},
	}
	Menus := []*fyne.Menu{}
	for i, key := range tools_key {
		tMenu := []*fyne.MenuItem{}
		for index, item := range tools[i] {
			newMenuItem := fyne.NewMenuItem(item, toolsFunc[i][index])
			tMenu = append(tMenu, newMenuItem)
		}
		Menus = append(Menus, fyne.NewMenu(key, tMenu...))
	}
	w.SetMainMenu(fyne.NewMainMenu(Menus...))

}
