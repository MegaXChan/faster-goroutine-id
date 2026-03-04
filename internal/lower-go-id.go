package internal

import (
	"runtime"
	"strconv"
	"strings"
)

func Gid() (id int64) {
	id = getGoidByStack()
	return
}

func getGoidByStack() int64 {
	var (
		buf [40]byte
		n   = runtime.Stack(buf[:], false)
		stk = strings.TrimPrefix(string(buf[:n]), "goroutine ")
	)

	idField := strings.Fields(stk)[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(err)
	}

	return int64(id)
}
