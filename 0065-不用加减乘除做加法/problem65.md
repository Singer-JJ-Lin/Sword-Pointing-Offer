# 不用加减乘除做加法
对每一位做如下运算，直到b为0
1. `c = (a&b)<<1`, 计算进位c
2. `a=a^b`, 非进位和
3. `b=c`， 将进位c赋给b
