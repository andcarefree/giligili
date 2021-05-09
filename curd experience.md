## crud项目中遇到的坑

1.go中结构体内嵌结构体时，创造结构体时不能直接赋值内嵌结构体的字段
```go
type in struct{
	ID uint
}
type out struct{
	in
}
func test()  {
    newout := out{
        ID:1,
    }
}
```
上述newout语句中，会报错
> Unknown field 'ID' in struct literal

只能在创建结构体之后赋值
```go
func test()  {
    var newout out
    newout.ID = 0
}
```