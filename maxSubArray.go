func midSubArray(nums []int, i, j,m int) int {
	left := nums[m]
	maxL:=left
	for k:=m-1; k>i;k--{
		left=left+nums[k]
		maxL = int(math.Max(float64(left),float64(maxL)))
	}
	right := nums[m+1]
	maxR:=right
	for k:=m+2; k<j;k++{
		right=right+nums[k]
		maxR = int(math.Max(float64(right),float64(maxR)))
	}
	return maxR+maxR
}

func maxSubArray(nums []int,i,j int) int {
    if i==j{return nums[i]}
	m := (i+j)/2
	left := float64(maxSubArray(nums,i,m))
	right := float64(maxSubArray(nums,m+1,j))
	mid := float64(midSubArray(nums,i,j,m))
	return int(math.Max(left,math.Max(right,mid)))
}
