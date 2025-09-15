# Sugar 🍬

一个基于 Go 1.23+ 泛型的高性能、精简的工具函数库，灵感来自 [Lodash](https://lodash.com) 和 [samber/lo](https://github.com/samber/lo)。

## ✨ 特性

- 🚀 **高性能**: 基于 Go 1.23+ 最新特性优化
- 🎯 **类型安全**: 完全使用泛型，编译时类型检查
- 📦 **精简设计**: 只包含最常用、最实用的操作
- 🔧 **零依赖**: 除 `golang.org/x/exp` 外无其他依赖
- ⚡ **内存友好**: 预分配切片容量，减少内存分配

## 🛠 安装

```bash
go get github.com/phuhao00/sugar
```

要求 Go 1.23 或更高版本。

## 📚 快速开始

```go
import "github.com/phuhao00/sugar"

// 切片操作
nums := []int{1, 2, 3, 4, 5}
doubled := sugar.Map(nums, func(x int, i int) int { return x * 2 })
// [2, 4, 6, 8, 10]

evens := sugar.Filter(nums, func(x int, i int) bool { return x%2 == 0 })
// [2, 4]

sum := sugar.Reduce(nums, func(acc int, x int, i int) int { return acc + x }, 0)
// 15

// 映射操作
m := map[string]int{"a": 1, "b": 2, "c": 3}
keys := sugar.Keys(m)        // ["a", "b", "c"]
values := sugar.Values(m)    // [1, 2, 3]

// 实用工具
result := sugar.Ternary(len(nums) > 3, "long", "short")  // "long"

// 错误处理
value := sugar.Must(strconv.Atoi("123"))  // 123, panic if error

// 数学操作
clampedValue := sugar.Clamp(10, 1, 5)     // 5
average := sugar.Mean([]float64{1, 2, 3, 4, 5})  // 3.0
```

## 📖 API 文档

### 切片操作

**核心函数**
- `Map[T, R]([]T, func(T, int) R) []R` - 变换每个元素
- `Filter[T]([]T, func(T, int) bool) []T` - 过滤元素
- `Reduce[T, R]([]T, func(R, T, int) R, R) R` - 归约操作
- `Find[T]([]T, func(T) bool) (T, bool)` - 查找元素
- `Contains[T]([]T, T) bool` - 检查是否包含
- `Uniq[T]([]T) []T` - 去重
- `GroupBy[T, U]([]T, func(T) U) map[U][]T` - 分组

**数组操作**
- `Chunk[T]([]T, int) [][]T` - 分块
- `Flatten[T]([][]T) []T` - 展平
- `Reverse[T]([]T) []T` - 反转
- `Shuffle[T]([]T) []T` - 随机排序
- `Drop[T]([]T, int) []T` - 删除前n个
- `Union[T](...[]T) []T` - 并集
- `Intersection[T](...[]T) []T` - 交集
- `Difference[T]([]T, ...[]T) []T` - 差集

### 映射操作

- `Keys[K, V](map[K]V) []K` - 获取所有键
- `Values[K, V](map[K]V) []V` - 获取所有值
- `PickBy[K, V](map[K]V, func(K, V) bool) map[K]V` - 条件选择
- `OmitBy[K, V](map[K]V, func(K, V) bool) map[K]V` - 条件排除
- `MapKeys[K, V, R](map[K]V, func(K, V) R) map[R]V` - 变换键
- `MapValues[K, V, R](map[K]V, func(K, V) R) map[K]R` - 变换值
- `Invert[K, V](map[K]V) map[V]K` - 键值互换

### 数学操作

- `Sum[T]([]T) T` - 求和
- `Mean[T]([]T) float64` - 平均值
- `Median[T]([]T) float64` - 中位数
- `Min[T]([]T) T` / `Max[T]([]T) T` - 最小/最大值
- `Clamp[T](T, T, T) T` - 限制范围
- `Abs[T](T) T` - 绝对值
- `Range(int, int) []int` - 数字序列

### 实用工具

**错误处理**
- `Must[T](T, error) T` - 转换错误为panic
- `Try(func() error) bool` - 安全执行
- `TryOr[T](func() (T, error), T) T` - 带默认值的安全执行

**条件操作**
- `Ternary[T](bool, T, T) T` - 三元运算符
- `Coalesce[T](...T) (T, bool)` - 获取第一个非零值
- `IsNil(any) bool` - 空值检查

**指针操作**
- `ToPtr[T](T) *T` - 值转指针
- `FromPtr[T](*T) T` - 指针转值

**异步操作**
- `Async[T](func() T) <-chan T` - 异步执行
- `Debounce[T](func(...T), time.Duration)` - 防抖
- `Throttle[T](func(...T), time.Duration)` - 节流

## 🏃 性能对比

与其他库的性能对比（基准测试基于相同的Map操作）：

```
BenchmarkMap/sugar.Map-8      9   126ms/op    40MB/op   1000001 allocs/op
BenchmarkMap/lo.Map-8         8   132ms/op    40MB/op   1000002 allocs/op  
BenchmarkMap/for-8           9   126ms/op    40MB/op   1000001 allocs/op
```

Sugar 与原生 for 循环性能相当，内存分配更优化。

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

- Fork 项目
- 创建特性分支
- 提交更改
- 创建 Pull Request

## 📄 许可证

MIT 许可证。详见 [LICENSE](LICENSE) 文件。

## 🙏 致谢

- [samber/lo](https://github.com/samber/lo) - 核心设计灵感
- [Lodash](https://lodash.com) - API 设计参考