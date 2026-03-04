# faster-goroutine-id

English | [中文](README_zh.md)

A Go library for efficiently obtaining Goroutine ID by directly accessing the goroutine structure pointer (getg) through assembly.

## Features

- 🚀 **High Performance**: Directly access goroutine structure through assembly, avoiding `runtime.Stack` calls
- 🔧 **Multi-architecture Support**: Supports x86, ARM, MIPS, PowerPC, RISC-V, and other CPU architectures
- 🛡️ **Safe and Reliable**: Dynamically calibrates offset to ensure compatibility across different Go versions
- 📦 **Easy to Use**: Provides a simple API interface
- ⚡ **Zero Dependencies**: Pure Go implementation with no external dependencies

## Background

In Go, obtaining the current goroutine ID typically requires parsing stack information through `runtime.Stack`, which has poor performance. This project directly accesses the goroutine structure pointer through assembly and reads the goroutine ID from the structure, resulting in significant performance improvement.

## Installation

```bash
go get github.com/MegaXChan/faster-goroutine-id
```

## Usage

### Basic Usage

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

### Performance Comparison

```go
package faster_goroutine_id

import (
    "fmt"
    "testing"
    "time"
    
    s "github.com/MegaXChan/faster-goroutine-id/internal"
)

func TestPerformance(t *testing.T) {
    // Using this library
    start1 := time.Now().UnixNano()
    id1 := GoroutineId()
    elapsed1 := time.Now().UnixNano() - start1
    
    // Using traditional method
    start2 := time.Now().UnixNano()
    id2 := s.Gid()
    elapsed2 := time.Now().UnixNano() - start2
    
    fmt.Printf("faster-goroutine-id: ID=%d, time=%v ns\n", id1, elapsed1)
    fmt.Printf("traditional method: ID=%d, time=%v ns\n", id2, elapsed2)
}
```

## Implementation Principle

### Core Idea

1. **Get goroutine structure pointer**: Directly obtain the current goroutine's `g` structure pointer through assembly instructions
2. **Dynamic offset calibration**: Dynamically calculate the offset of goroutine ID in the `g` structure during initialization
3. **Read goroutine ID**: Directly read the goroutine ID through pointer and offset

### Assembly Implementation

Provides corresponding assembly implementations for different CPU architectures:

- `getg_amd64.s`: x86-64 architecture
- `getg_arm64.s`: ARM64 architecture
- `getg_386.s`: x86 architecture
- `getg_arm.s`: ARM architecture
- `getg_ppc64.s`: PowerPC 64-bit architecture
- `getg_riscv64.s`: RISC-V 64-bit architecture
- And other architecture support

### Dynamic Calibration

Since the layout of the `g` structure may vary across different Go versions, this project dynamically calibrates the goroutine ID offset during initialization:

```go
func calibrateOffset() {
    // Verify offset accuracy in multiple goroutines
    // Use majority voting algorithm to determine the correct offset
}
```

## Supported Platforms

- **Operating Systems**: Linux, Windows, macOS
- **Architectures**:
  - x86: amd64, 386
  - ARM: arm64, arm (v6, v7)
  - MIPS: mips, mips64
  - PowerPC: ppc64, ppc64le
  - RISC-V: riscv64
  - s390x

## Performance Data

According to test results, using this library significantly improves performance for obtaining goroutine ID compared to traditional methods:

- **Traditional method** (`runtime.Stack`): ~1000-2000 ns
- **This library**: ~10-50 ns

Performance improvement is approximately 20-100x.

## Notes

1. **Go Version Compatibility**: Due to dependency on goroutine structure internal layout, different Go versions may require re-calibration of offset
2. **Platform Limitations**: Only supports CPU architectures with implemented assembly
3. **Production Environment**: Recommended to conduct thorough testing in critical paths

## Contributing

Issues and Pull Requests are welcome!

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Create a Pull Request

## License

This project is open source under the MIT License. See the [LICENSE](LICENSE) file for details.

## Author

MegaXChan

## Acknowledgments

- Thanks to the Go language community
- Thanks to all contributors

## Related Projects

- [go-id](https://github.com/MegaXChan/go-id): Traditional goroutine ID acquisition library
- [goid](https://github.com/petermattis/goid): Another goroutine ID implementation

## Changelog

### v0.0.2 (2026-03-04)
- Initial release
- Multi-platform and multi-architecture support
- Dynamic offset calibration implementation
- Performance tests and examples provided

---

**Tip**: If you encounter any issues, please check [Issues](https://github.com/MegaXChan/faster-goroutine-id/issues) first or create a new issue.
