🍬 **Sugar v2.0.0** - 基于 Go 1.23+ 泛型的高性能、类型安全工具函数库

## ✨ 重大版本特性

### 🚀 性能表现
- **850K+ ops/sec** 超高吞吐量，与原生 for 循环性能相当
- **单次内存分配** 优化策略
- **完全泛型化** 100% 类型安全，编译时检查

### 📦 核心功能 (120+ 函数)
- **🔧 core.go**: Map, Filter, Reduce, Find, Uniq, GroupBy, Chunk, Flatten
- **🎯 slice.go**: Union, Intersection, Difference, Drop, Compact, Partition  
- **🗺️ map.go**: Keys, Values, PickBy, MapKeys, MapValues, Invert
- **🧮 math.go**: Sum, Mean, Median, Min, Max, Clamp, Range
- **🛠️ util.go**: Must, Try, AsyncRun, Debounce, ToPtr, Coalesce
- **🎛️ condition.go**: Ternary, IfElse 链式条件判断

## 🚀 快速开始

### 安装
```bash
go get github.com/phuhao00/sugar/v2@v2.0.0
```

### 使用示例
```go
import "github.com/phuhao00/sugar/v2"

// 切片操作
nums := []int{1, 2, 3, 4, 5}
doubled := sugar.Map(nums, func(x int, i int) int { return x * 2 })
evens := sugar.Filter(nums, func(x int, i int) bool { return x%2 == 0 })
sum := sugar.Reduce(nums, func(acc int, x int, i int) int { return acc + x }, 0)

// 映射操作  
m := map[string]int{"apple": 5, "banana": 3}
keys := sugar.Keys(m)  // [apple, banana]
expensive := sugar.PickBy(m, func(k string, v int) bool { return v >= 5 })

// 实用工具
result := sugar.Ternary(len(nums) > 3, "long", "short")
value := sugar.Must(strconv.Atoi("123"))  // 123
```

## ⚠️ 破坏性变更

### 模块路径更新 (重要！)
```go
// 旧版本
import "github.com/phuhao00/sugar"

// v2.0.0 新版本 - 必须更新
import "github.com/phuhao00/sugar/v2"
```

### 函数重命名
- `SliceFiltrate` → `Filter`
- `SliceUpdateElement` → `Map`  
- `SliceUniq` → `UniqBy`
- `FiltrateBy` → `PickBy`

### 版本要求
- **Go 1.23+** (利用最新泛型特性)

## 🏆 性能对比

| 操作 | Sugar v2.0 | 原生 for | samber/lo | go-funk |
|------|-----------|----------|-----------|---------|
| Map 1000 元素 | **1670 ns** | 1650 ns | 1750 ns | 11000+ ns |
| 内存分配 | **1 alloc** | 1 alloc | 1 alloc | 3+ allocs |
| 类型安全 | ✅ 编译时 | ✅ 编译时 | ✅ 编译时 | ❌ 运行时 |

## 📚 完整文档

- [中文 README](README.md) - 完整使用指南  
- [更新日志](VERSION.md) - 详细版本变更
- [API 文档](https://pkg.go.dev/github.com/phuhao00/sugar/v2) - godoc 标准文档

## 🙏 致谢

灵感来源于 [samber/lo](https://github.com/samber/lo) 和 [Lodash](https://lodash.com)

---

**让 Go 开发更甜蜜！** 🍬✨
