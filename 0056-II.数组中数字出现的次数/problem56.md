# 数组中数字出现的次数
考虑数字的二进制形式，对于出现三次的数字，各二进制位出现的次数都是3的倍数。因此，统计所有数字的各二进制位中1的出现次数，并对3求余，结果则为只出现一次的数字

### 有限状态机
对于所有数字中的某二进制位1的个数，存在3种状态，即对3余数为0,1,2。由于二进制只能表示0,1，因此需要使用两个二进制位来表示3个状态。设此两位分别为ones和twos,状态转移如下

00->01->10->00

**ones和twos计算公式如下**
```
ones = ones ^ n & ~twos
twos = twos ^ n & ~ones
```
由于遍历完所有数字后，每个二进制位都属于00或者01的状态，所以只需要用ones就能表示，因此返回ones就是结果




