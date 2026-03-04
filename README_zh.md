# faster-goroutine-id

[English](README.md) | 中文

一个通过汇编直接获取goroutine结构体指针(getg)来高效获取Goroutine ID的Go库。

## 特性

- 🚀 **高性能**: 通过汇编直接访问goroutine结构体，避免调用`runtime.Stack`
- 🔧 **多架构支持**: 支持x86、ARM、MIPS、PowerPC、RISC-V等多种CPU架构
- 🛡️ **安全可靠**: 动态校准offset，确保在不同Go版本中正常工作
- 📦 **简单易用**: 提供简洁的API接口
- ⚡ **零依赖**: 纯Go实现，无外部依赖

## 背景

在Go语言中，获取当前goroutine ID通常需要通过`runtime.Stack`解析堆栈信息，这种方式性能较差。本项目通过汇编直接获取goroutine结构体指针，然后从结构体中读取goroutine ID，性能显著提升。

## 安装

```bash
go get github.com/MegaXChan/faster-goroutine-id
```

## 使用方法

### 基本使用

```go
package main

import (
    "fmt"
    "sync"
    
    goid "github.com/MegaXChan/faster-goroutine-id"
)

func main() {
    wait := sync.WaitGroup{}
    for i := 0; i < 10; i++ {
        wait.Add(1)
        go func(v int) {
            id := goid.GoroutineId()
            fmt.Printf("index: %v, goroutine id: %v\n", v, id)
            defer wait.Done()
        }(i)
    }
    wait.Wait()
}
```

### 性能对比

```go
package faster_goroutine_id

import (
    "fmt"
    "testing"
    "time"
    
    s "github.com/MegaXChan/faster-goroutine-id/internal"
)

func TestPerformance(t *testing.T) {
    // 使用本库
    start1 := time.Now().UnixNano()
    id1 := GoroutineId()
    elapsed1 := time.Now().UnixNano() - start1
    
    // 使用传统方法
    start2 := time.Now().UnixNano()
    id2 := s.Gid()
    elapsed2 := time.Now().UnixNano() - start2
    
    fmt.Printf("faster-goroutine-id: ID=%d, time=%v ns\n", id1, elapsed1)
    fmt.Printf("traditional method: ID=%d, time=%v ns\n", id2, elapsed2)
}
```

## 实现原理

### 核心思想

1. **获取goroutine结构体指针**: 通过汇编指令直接获取当前goroutine的`g`结构体指针
2. **动态校准offset**: 在初始化时动态计算goroutine ID在`g`结构体中的偏移量
3. **读取goroutine ID**: 通过指针和偏移量直接读取goroutine ID

### 汇编实现

针对不同CPU架构提供了相应的汇编实现：

- `getg_amd64.s`: x86-64架构
- `getg_arm64.s`: ARM64架构
- `getg_386.s`: x86架构
- `getg_arm.s`: ARM架构
- `getg_ppc64.s`: PowerPC 64位架构
- `getg_riscv64.s`: RISC-V 64位架构
- 以及其他架构支持

### 动态校准

由于不同Go版本中`g`结构体的布局可能不同，本项目在初始化时动态校准goroutine ID的偏移量：

```go
func calibrateOffset() {
    // 在多个goroutine中验证offset的准确性
    // 使用多数表决算法确定正确的offset
}
```

## 支持的平台

- **操作系统**: Linux, Windows, macOS
- **架构**:
  - x86: amd64, 386
  - ARM: arm64, arm (v6, v7)
  - MIPS: mips, mips64
  - PowerPC: ppc64, ppc64le
  - RISC-V: riscv64
  - s390x


## 性能数据

根据测试结果，使用本库获取goroutine ID的性能比传统方法提升显著：

- **传统方法** (`runtime.Stack`): ~1000-2000 ns
- **本库方法**: ~10-50 ns

性能提升约20-100倍。

## 注意事项

1. **Go版本兼容性**: 由于依赖goroutine结构体内部布局，不同Go版本可能需要重新校准offset
2. **平台限制**: 仅支持已实现汇编的CPU架构
3. **生产环境**: 建议在关键路径进行充分测试

## 贡献

欢迎提交Issue和Pull Request！

1. Fork本仓库
2. 创建功能分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 创建Pull Request

## 许可证

本项目基于MIT许可证开源。详见[LICENSE](LICENSE)文件。

## 作者

MegaXChan

## 致谢

- 感谢Go语言社区
- 感谢所有贡献者

## 相关项目

- [go-id](https://github.com/MegaXChan/go-id): 传统的goroutine ID获取库
- [goid](https://github.com/petermattis/goid): 另一个goroutine ID获取实现


**提示**: 如果在使用中遇到问题，请先查看[Issues](https://github.com/MegaXChan/faster-goroutine-id/issues)或创建新的Issue。
