# 把数组排成最小的数
调用标准库的排序方法进行排序，自定义比较方法如下
```Plain Text
str[x] + str[y] < str[y] + str[x] => x < y 
str[x] + str[y] > str[y] + str[x] => x > y 
```
