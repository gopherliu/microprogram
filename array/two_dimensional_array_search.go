package array

/**
在一个二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
*/

// 根据题意，一个很简单的思路就是能想到：两层循环，挨个比较，在的话就返回即可
// 时间复杂度上不难看出O(n^2)
func CommonSearch(source [][]int, target int) bool {
	for _, v := range source {
		for _, vv := range v {
			if vv == target {
				return true
			}
		}
	}
	return false
}

// 但是这道题如果就像上述那样解决了，也肯定就失去出这道题的含义了；
// 一个二维数组，其实就是一个矩阵。如下图所示：
// 3  5  7  8
// 5  6  9  14
// 11 18 20 32
// 这个矩阵从上往下依次递增，从左往右依次递增；所以我们可以这样：
// 第一，从右上角的8开始，如果比8小，那么8这一列就排除（因为8这里一列下面的数字都比8大）
// 第二，之后就和7比较，如果目标值比7大，那么就把7这一行排除掉（因为7这一行前面的数字都比7小）
// 以此类推，知道找到目标值；下面是源码
func EfficiencySearch(source [][]int, target int) bool {
	var (
		rowCount, columnCount int  // 行数、列数
		isHas                 bool // 是否存在
	)
	rowCount = len(source)
	columnCount = len(source[0])

	i, j := 0, columnCount-1 // arr[i][j]永远代表左上角位置的元素
	for {

		// 越界返回（查找失败）
		if i >= rowCount || j < 0 {
			break
		}

		// 找到返回
		if source[i][j] == target {
			isHas = true
			break
		}

		// 目标值比左上角的大，排除一行
		if source[i][j] < target {
			i++
			continue
		}

		// 目标值比左上角的小，排除一列
		if source[i][j] > target {
			j--
			continue
		}
	}
	return isHas
}
