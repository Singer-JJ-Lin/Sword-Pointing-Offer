# 旋转数组的最小数字
### 旋转数组的性质
最小值右边的元素一定比数组中最右的元素小，最小值左边的元素一定都比数组中最右的元素大


利用以上性质，可以进行二分查找
### 二分查找
在二分查找的每一步中，左边界l，右边界为r，区间的中点为mid，最小值就在该区间内。我们将numbers[mid]与numbers[r]做比较，得到如下三种情况：
- **numbers[mid] < numbers[r]:** 这说明numbers[mid]是最小值右侧的元素，因此我们可以忽略二分查找区间的右半部分
- **numbers[mid] > numbers[r]:** 这说明numbers}[mid]是最小值左侧的元素，因此我们可以忽略二分查找区间的左半部分
- **numbers[mid] == numbers[r]:** 这说明由于重复元素的存在，我们并不能确定numbers[mid]究竟在最小值的左侧还是右侧,但是由于是相等的，我们可以忽略二分查找区间的右端点




