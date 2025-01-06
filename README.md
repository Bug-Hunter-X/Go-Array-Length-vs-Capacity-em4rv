# Go Array Length vs Capacity

This example demonstrates a common misconception about arrays in Go.  Arrays in Go have a fixed size, determined at compile time. This means the length and capacity are always identical.  The `cap()` function, which typically returns the capacity of a slice, returns the same value as `len()` when applied to an array.

This is not necessarily a bug, but it's a potential source of confusion for developers coming from languages with more dynamic array handling. Understanding this distinction is crucial for writing correct and efficient Go code.