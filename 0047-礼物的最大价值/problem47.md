# 礼物的最大价值
<script type="text/javascript" src="http://cdn.mathjax.org/mathjax/latest/MathJax.js?config=default"></script>
从棋盘的左上角开始拿格子里的礼物，并每次只能**向右**或者**向下**移动一格、直到到达棋盘的右下角
### 定义状态
dp[i][j]代表了从左上角走到当前位置{i,j}可以获得礼物的最大值
### 状态转移方程
1. i = 0, j = 0, 为起始位置，不用处理
2. i = 0, j != 0, 为矩阵第一行元素，只可从左边到达
3. i != 0, j = 0, 为矩阵第一列元素，只可从上边到达
4. i != 0, j != 0, 可从左边或上边到达

$$

dp[i][j] = \left\{  
\begin{array}{**lr**}  
grid[i][j], & i = 0, j = 0 \\  
grid[i][j]+dp[i][j-1], & i=0,j \neq 0.\\  
grid[i][j] + dp[i-1][j], & i \neq 0, j = 0 \\
grid[i][j] + max(dp[i-1][j], dp[i]dp[j-1]) & i \neq 0, j \neq 0
\end{array}  
\right.

$$


 