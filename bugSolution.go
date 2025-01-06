The solution is to use slices instead of arrays when you need a dynamically sized collection.  Slices provide a view onto an underlying array, offering flexibility in size and capacity. 

```go
func main() {
    var a []int
    a = append(a, 1, 2, 3)
    fmt.Println(len(a))
    fmt.Println(cap(a))
    a = append(a, 4, 5, 6, 7, 8, 9, 10)
    fmt.Println(len(a))
    fmt.Println(cap(a))
}
```
This code demonstrates the difference. The capacity of a slice might grow as you append elements, whereas the length is always the number of elements in the slice.