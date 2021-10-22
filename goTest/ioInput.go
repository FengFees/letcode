package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func main() {

	h := IntHeap{2, 1, 5}
	heap.Init(&h)
	heap.Push(&h, 3)
	fmt.Printf("堆顶元素: %d\n", h[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(&h))
	}
	var d1 []byte
	d1 = append(d1, '1')
	d1 = append(d1, '2')
	fmt.Println(string(d1[:1]))
	//r := bufio.NewScanner(os.Stdin)
	////math.MinInt64
	////reflect.d
	//for r.Scan() {
	//	//input := r.Text()
	//	//nums := strings.Fields(input)
	//	//n1,_ := strconv.ParseFloat(nums[0],64)
	//	//n2,_ := strconv.ParseFloat(nums[1],64)
	//	//if n1 == n2 {
	//	//	fmt.Println("YES")
	//	//}else {
	//	//	fmt.Println("NO")
	//	//}
	//	//str1 := r.Text()
	//	//r.Scan()
	//	//
	//	//dp := make([][501]int, 501)
	//	//for i := 1; i < ; i++ {
	//	//
	//	//}
	//	//m := make(map[int]int)
	//}

	//dp := make([][501]int,501)
	//dp[500][500] =1
	//fmt.Println(dp[500][500])
	//math.

}
