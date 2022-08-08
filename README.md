# golang_number_of_1_bits

Write a function that takes an unsigned integer and returns the number of '1' bits it has (also known as the [Hamming weight](http://en.wikipedia.org/wiki/Hamming_weight)).

**Note:**

- Note that in some languages, such as Java, there is no unsigned integer type. In this case, the input will be given as a signed integer type. It should not affect your implementation, as the integer's internal binary representation is the same, whether it is signed or unsigned.
- In Java, the compiler represents the signed integers using [2's complement notation](https://en.wikipedia.org/wiki/Two%27s_complement). Therefore, in **Example 3**, the input represents the signed integer. `3`.

## Examples

**Example 1:**

```
Input: n = 00000000000000000000000000001011
Output: 3
Explanation: The input binary string00000000000000000000000000001011 has a total of three '1' bits.

```

**Example 2:**

```
Input: n = 00000000000000000000000010000000
Output: 1
Explanation: The input binary string00000000000000000000000010000000 has a total of one '1' bit.

```

**Example 3:**

```
Input: n = 11111111111111111111111111111101
Output: 31
Explanation: The input binary string11111111111111111111111111111101 has a total of thirty one '1' bits.

```

**Constraints:**

- The input must be a **binary string** of length `32`.

**Follow up:**

If this function is called many times, how would you optimize it?

## 解析

給定一個 unsigned integer num (32 bit 整數)

要求寫一個演算法去計算一共有多少 bit 是 1

最直覺的方式就是逐個 bit 去檢查哪些 bit 是 1 然後累加

每次對原本的值做 unsigned shift 直到 shift 到所有值是 0 為止

![](https://i.imgur.com/OFufEKr.png)

因為每個 unsigned integer 都是 32 bit

所以時間複雜度是 O(1)

然而是否有什麼方式只從 bit 是非零的部份來做紀錄

可以發現每次把 num & num-1 每次可以把非零的最後一個 bit 轉換成 0

![](https://i.imgur.com/OaXP6La.png)

透過這種方式每次當 num 非零時

就更新 num = num & num -1 ， count++

當 num = 0 時, 回傳 count

這樣的時間複雜度是 O(1)

## 程式碼
```go
func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		num = num & (num - 1)
		count += 1
	}
	return count
}
func hammingWeightV2(num uint32) int {
	count := 0
	for num != 0 {
		if (num & 1) != 0 {
			count++
		}
		num >>= 1
	}
	return count
}
```
## 困難點

1. 要思考出透過 n & n-1 的方式來優化找 bit 不為0 的方式
2. 對於使用 bit shift 必須要使用 unsigned shift 否則會造成 shift 後的 bit 保持原本值造成無窮回圈

## Solve Point

- [x]  初始化 count = 0
- [x]  當 num ≠ 0 更新 num = num & num - 1, count += 1
- [x]  當 num = 0 回傳 count