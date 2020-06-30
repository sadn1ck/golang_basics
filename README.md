# Strings
- strings in go are very different
- represented as a slice of bytes
- for making it similar to other languages, have to be converted to what is called runes
- say for example ` s := "Hello World" `
- to convert ro runes ` st = []rune(s) `
- only then you can print non traditional characters and the characters themselves (otherwise ASCII value is printed)

```
for index, char := range st {
    fmt.Printf("%c", r[i])
}
fmt.Println("")
for index, char := range st {
    fmt.Printf("%v ", r[i])
}
```
Outputs: 
```
Hello World
72 101 108 108 245 32 87 111 114 108 100
```
The contents of runes also have type int32 unline strings which store uint8.


For better understanding please read [this medium article on strings and runes](https://medium.com/rungo/string-data-type-in-go-8af2b639478)
