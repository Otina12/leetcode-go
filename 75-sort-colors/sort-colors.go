func sortColors(nums []int)  {
    low, middle, high := 0, 0, len(nums) - 1

    for middle <= high {
        if nums[middle] == 1 {
            middle += 1
        } else if nums[middle] == 0 {
            temp := nums[middle]
            nums[middle] = nums[low]
            nums[low] = temp
            low += 1
            middle += 1
        } else if nums[middle] == 2 {
            temp := nums[middle]
            nums[middle] = nums[high]
            nums[high] = temp
            high -= 1
        }
    }
}