# 把数字翻译成字符串
### 状态转移
dp[i] 代表以 $ x_i $ 为结尾的数字的翻译方案数量
### 转移方程
- 若当前位置和上一个位置组成的数字在10到25之间，dp[i] = dp[i-1] + dp[i-2]
- 若当前位置和上一个位置组成的数字不在10到25之间，dp[i] = dp[i-1]
### 初始状态
dp[0] = dp[1] = 1, 即没有数字和只有一个数字的翻译方法均为1
### 返回值
dp[n]， 总共的翻译方案数量


