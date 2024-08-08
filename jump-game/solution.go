package jumpgame

func CanReach(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}

	dp := make([]bool, n)
	dp[0] = true

	for i := 0; i < n; i++ {
		if dp[i] {
			for j := 1; j <= nums[i]; j++ {
				if i+j < n {
					dp[i+j] = true
				}
			}
		}
	}

	return dp[n-1]
}
