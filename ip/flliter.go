package ip

import (
	"IP_pkg_analyze/ui"
	"fyne.io/fyne/v2/data/binding"
)

var portChan chan int
var ipChan chan string

func flliterBySourceIp(list binding.StringList, ip string, PkgInfos []PkgRow) {
	ipChan = make(chan string, 2)
	var res []PkgRow
	for _, v := range PkgInfos {
		if v.Source == ip {
			res = append(res, v)
		}
	}
	ui.LoadPkgList(list, PkgInfos)
	ipChan <- ip
}
func flliterByDstIp(list binding.StringList, ip string, PkgInfos []PkgRow) {
	var res []PkgRow
	for _, v := range PkgInfos {
		if v.Dest == ip {
			res = append(res, v)
		}
	}
	ui.LoadPkgList(list, PkgInfos)

}
func flliterByPort(list binding.StringList, port int, PkgInfos []PkgRow) {
	ui.LoadPkgList(list, PkgInfos)
}
