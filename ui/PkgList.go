package ui

import (
	"IP_pkg_analyze/ip"
	fyne2 "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	//"time"
)

func loadPkgList() (*fyne2.Container, binding.StringList) {
	pkgText := []string{
		"No.     ", "Time                          ", "Source              ", "Dst                      ",
		"Protocol   ", "Length     ", "Info                                 ",
	}
	pkgTextContainer := container.NewHBox()
	for _, v := range pkgText {
		pkgTextContainer.Add(widget.NewButton(v, func() {

		}))
	}
	PkgStringList := binding.NewStringList()
	PkgList := widget.NewListWithData(PkgStringList,
		func() fyne2.CanvasObject {
			return widget.NewLabel("")
		},
		func(item binding.DataItem, object fyne2.CanvasObject) {
			i := item.(binding.String)
			l := object.(*widget.Label)
			s, _ := i.Get()
			l.SetText(s)
		})
	PkgList.OnSelected = func(id widget.ListItemID) {
		NewLayersData(id+1, ip.PkgInfos[id])
		LayersWidget.Refresh()
		NewPkgInfoData(ip.PkgInfos[id])
		//pkg.Set(PkgBytes2String(ip.PkgInfos[id].Data()))

	}
	s := widget.NewSeparator()
	PkgListContainer := container.NewBorder(pkgTextContainer, nil, nil, s, PkgList)
	PkgListContainer.Resize(fyne2.NewSize(1600, 280))
	return PkgListContainer, PkgStringList

}
func LoadPkgList(list binding.StringList, Pkgs []ip.PkgRow) {
	err := list.Set([]string{})
	if err != nil {
		fyne2.LogError("loadPkgList重置list失败", err)
	}
	for _, v := range Pkgs {
		list.Append(v.FormatePkgListInfo())
	}
}
