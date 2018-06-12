package alg

/*基础排序*/

/**
	选择排序  O(n^2)
	从数组中依次找出最小的元素,放到前面(第一个,第二个第三个...)
 */
func selectSort(arr []int, size int) []int {
	// 遍历数组, 第一次遍历[0:],第二次[1:]
	for i := 0; i < size; i++ {
		// 寻找[i,n)区间内的最小值
		// 子循环当前的最小值索引
		minIndex := i
		for j := i + 1; j < size; j++ {
			// 如果该数小于目前的最小数,则记录该数索引
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		// 交换:将该最小值放到前面
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}



/**
	插入排序
 */
func insertSort(arr []int, size int) []int {
	// 默认第一个数为有序，从第二个元素开始遍历
	for i:=1; i < size; i++{
		// 当 当前元素i 大于 前个元素（该元素属于有序队列），则当前元素也属于有序，无需操作
		if arr[i] > arr[i-1]{
			continue
		}

		// 遍历有序队列
		/**
			方案一， 比较出比当前元素i小的元素的位置，然后删除当前元素i，然后再该位置前一个位置插入 当前元素i的值
			因为切片插入删除需要大量复制，所以很慢
			10000个元素，执行时间200+ms
		 */
		//temp := arr[i]
		//for j:=0; j < i;j++{
		//	// 如果当前元素i小于 元素j， 则
		//	if arr[i] < arr[j]{
		//		// 删除 元素i
		//		arr = append(arr[:i],arr[i+1:]...)
		//		// 插入它到元素j-1的位置
		//		arr =  append(arr[:j],append([]int{temp},arr[j:]...)...)
		//	}
		//}

		/**
			方案2： 从后往前比较，如果大于，则复制该元素x到x+1的位置（第一次复制时覆盖了元素i原来的位置）,
			当元素x小于i后，则将 i元素放到x+1的位置即可（此时x+1位置的原元素已在上一次被复制到了x+2的位置）
			10000个元素，执行时间 90+ms
			1，5，6，7，3,...
			1，5，6，7，7
			1，5，6，6，7
			1，5，5，6，7
			1，3，5，6，7

		 */
		//temp := arr[i]
		//j := i-1
		//for j >= 0 && arr[j] > temp {
		//	arr[j+1] = arr[j]
		//	j--
		//}
		//arr[j+1] = temp


		/**
			方案三：从后往前（从i元素开始）遍历有序队列，如果 元素x 小于 元素x-1,就交换两者位置
			10000个元素，执行时间:30+ms
			1,5,6,7,3,...
			1,5,6,3,7
			1,5,3,6,7
			1,3,5,6,7
		 */
		for j:=i; j >0 && arr[j] < arr[j-1]; j--{
			arr[j],arr[j-1] = arr[j-1],arr[j]
		}

		/**
			简单优化了方案三，不过性能感觉基本没有提升
		 */
		//temp := arr[i]
		//var j int
		//for j = i; j >0 && arr[j-1] > temp; j--{
		//	arr[j],arr[j-1] = arr[j-1],arr[j]
		//}
		//arr[j]=temp
	}
	return arr
}