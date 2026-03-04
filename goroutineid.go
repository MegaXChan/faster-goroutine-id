package faster_goroutine_id

import (
	"fmt"
	"sync"
	"unsafe"

	"github.com/MegaXChan/faster-goroutine-id/internal"
)

func getg() unsafe.Pointer

var offset uintptr = 0 // 默认值，会在 init 中校准

func init() {
	// 动态计算 offset
	calibrateOffset()
	if offset == 0 {
		panic("can not get offset")
	}
	fmt.Printf("offset is %d\n", offset)
}

func mostFrequent(nums []uintptr) (uintptr, int) {
	if len(nums) == 0 {
		return 0, 0
	}

	counter := make(map[uintptr]int)

	var (
		maxVal   uintptr
		maxCount int
	)

	for _, v := range nums {
		counter[v]++
		if counter[v] > maxCount {
			maxCount = counter[v]
			maxVal = v
		}
	}

	return maxVal, maxCount
}

func calibrateOffset() {

	successOffs := []uintptr{}
	lock := sync.Mutex{}

	// 在多个 goroutine 中验证 offset 的准确性
	const numGoroutines = 10
	results := make(chan bool, numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			off := uintptr(0)
			for {
				off = findOffsetInitial(off)
				if off > 0 {
					lock.Lock()
					successOffs = append(successOffs, off)
					lock.Unlock()
					break
				} else if off == 0 {
					break
				}
				off++
			}

			// 比较结果
			results <- (off > 0)
		}()
	}

	// 收集结果
	correctCount := 0
	for i := 0; i < numGoroutines; i++ {
		if <-results {
			correctCount++
		}
	}

	frequent, i := mostFrequent(successOffs)
	if i > numGoroutines/2 {
		offset = frequent
	}

}

func findOffsetInitial(offset uintptr) uintptr {
	// 获取当前 goroutine ID 作为参考值
	expectedID := internal.Gid()
	gp := getg()
	if gp == nil {
		panic("can not get g")
		return 0
	}

	base := uintptr(gp)
	fmt.Printf("base is %d\n", base)

	// 在 g 结构体附近搜索 goroutine ID
	// 使用已知的 offset 范围进行搜索
	const searchRange = 512

	// 定义一个安全读取内存的函数
	readSafe := func(addr uintptr) (int64, bool) {
		defer func() {
			r := recover() // 忽略 panic
			if r != nil {
				fmt.Printf("recover is %v", r)
			}
		}()
		return *(*int64)(unsafe.Pointer(addr)), true
	}

	for off := offset; off < searchRange; off += 1 {
		addr := base + off
		if id, ok := readSafe(addr); ok && id == expectedID {
			fmt.Printf("id is %d off is %v\n", id, off)
			return off
		}
	}

	return 0
}

func GoroutineId() int64 {
	gp := getg()
	return *(*int64)(unsafe.Pointer(uintptr(gp) + offset))
}
