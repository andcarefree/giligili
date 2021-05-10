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


2.```gorm```框架对数据库的删除是软删除的方法，如果映射结构体内嵌了```gorm.Model```，会在该内嵌结构体的字段中记录删除时间，但数据库记录仍然存在。如果需要删除数据库记录可以自己写回收站异步批量在数据库中删除记录