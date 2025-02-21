package sort

type SortInterface interface {
	Sort(arr []int) []int
	SortUseThread(arr []int) []int
}

func Sort(arr []int, sort SortInterface) []int {
	return sort.Sort(arr)
}

func SortUseThread(arr []int, sort SortInterface) []int {
	return sort.SortUseThread(arr)
}
