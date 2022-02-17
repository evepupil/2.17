package tools

import (
	"IP_pkg_analyze/ip"
	"sort"
)

type rowsBySourceIp []ip.PkgRow
type rowsByDstIp []ip.PkgRow
type rowsByLength []ip.PkgRow
type rowsByPort []ip.PkgRow

func PkgSortBySourceIp(PkgInfos []ip.PkgRow) []ip.PkgRow {
	sort.Sort(rowsBySourceIp(PkgInfos))
	return PkgInfos
}
func PkgSortBySourceIpReverse(PkgInfos []ip.PkgRow) []ip.PkgRow {
	sort.Sort(rowsBySourceIp(PkgInfos))
	sort.Reverse(rowsBySourceIp(PkgInfos))
	return PkgInfos
}
func PkgSortByDstIp(PkgInfos []ip.PkgRow) []ip.PkgRow {
	sort.Sort(rowsByDstIp(PkgInfos))
	return PkgInfos
}
func PkgSortByDstIpReverse(PkgInfos []ip.PkgRow) []ip.PkgRow {
	sort.Sort(rowsByDstIp(PkgInfos))
	sort.Reverse(rowsByDstIp(PkgInfos))
	return PkgInfos
}

func PkgSortByLength(PkgInfos []ip.PkgRow) []ip.PkgRow {
	sort.Sort(rowsByLength(PkgInfos))
	return PkgInfos
}
func PkgSortByLengthReverse(PkgInfos []ip.PkgRow) []ip.PkgRow {
	sort.Sort(rowsByLength(PkgInfos))
	sort.Reverse(rowsByLength(PkgInfos))
	return PkgInfos
}

func PkgSortByPort(PkgInfos []ip.PkgRow) []ip.PkgRow {
	sort.Sort(rowsByPort(PkgInfos))
	return PkgInfos
}
func PkgSortByPortReverse(PkgInfos []ip.PkgRow) []ip.PkgRow {
	sort.Sort(rowsByPort(PkgInfos))
	return PkgInfos
}
func (receiver rowsBySourceIp) Len() int {
	return len(receiver)
}
func (receiver rowsBySourceIp) Swap(i, j int) {
	receiver[i], receiver[j] = receiver[j], receiver[i]
}
func (receiver rowsBySourceIp) Less(i, j int) bool {
	return receiver[i].Source < receiver[j].Source
}

func (receiver rowsByDstIp) Len() int {
	return len(receiver)
}
func (receiver rowsByDstIp) Swap(i, j int) {
	receiver[i], receiver[j] = receiver[j], receiver[i]
}
func (receiver rowsByDstIp) Less(i, j int) bool {
	return receiver[i].Source < receiver[j].Source
}

func (receiver rowsByLength) Len() int {
	return len(receiver)
}
func (receiver rowsByLength) Swap(i, j int) {
	receiver[i], receiver[j] = receiver[j], receiver[i]
}
func (receiver rowsByLength) Less(i, j int) bool {
	return receiver[i].Source < receiver[j].Source
}

func (receiver rowsByPort) Len() int {
	return len(receiver)
}
func (receiver rowsByPort) Swap(i, j int) {
	receiver[i], receiver[j] = receiver[j], receiver[i]
}
func (receiver rowsByPort) Less(i, j int) bool {
	return receiver[i].Source < receiver[j].Source
}
