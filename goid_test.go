package faster_goroutine_id

import (
	"fmt"
	"testing"
	"time"

	s "github.com/MegaXChan/faster-goroutine-id/internal"
)

func TestGetGoId(t *testing.T) {

	// 测试 Get() 函数
	n1 := time.Now().UnixNano()
	var id int64
	for i := 0; i < 1; i++ {
		id = GoroutineId()
	}

	n2 := time.Now().UnixNano()
	if id <= 0 {
		t.Errorf("Expected positive goroutine ID, got %d", id)
	}
	fmt.Printf("Goroutine ID: %d time : %v\n", id, n2-n1)
	n1 = time.Now().UnixNano()
	var id2 int64
	for i := 0; i < 1; i++ {
		id2 = s.Gid()
	}
	n2 = time.Now().UnixNano()
	fmt.Printf("RAW Goroutine ID: %d time : %v\n", id2, n2-n1)

	// 在另一个 goroutine 中测试
	ch := make(chan int64)
	go func() {
		fmt.Printf("RAW2 Goroutine ID: %d\n", s.Gid())
		ch <- s.Gid()
	}()
	otherID := <-ch
	if otherID <= 0 {
		t.Errorf("Expected positive goroutine ID in goroutine, got %d", otherID)
	}
	fmt.Printf("Goroutine ID in another goroutine: %d\n", otherID)
}
