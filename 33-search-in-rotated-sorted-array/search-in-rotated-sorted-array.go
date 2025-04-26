func search(nums []int, target int) int {
    left, right := 0, len(nums) - 1

    for left <= right {
        middle := left + (right - left) / 2
        fmt.Println(middle)

        if nums[middle] == target {
            return middle
        }

        if nums[middle] > nums[right] {
            if target <= nums[right] || target > nums[middle] {
                left = middle + 1
            } else {
                right = middle - 1
            }
        } else if nums[middle] < nums[left] {
            if target >= nums[left] || target < nums[middle] {
                right = middle - 1
            } else {
                left = middle + 1
            }
        } else { // sorted
            if target > nums[middle] {
                left = middle + 1
            } else {
                right = middle - 1
            }
        }
    }

    return -1
}