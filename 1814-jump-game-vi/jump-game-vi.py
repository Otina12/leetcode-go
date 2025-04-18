class Solution:
    def maxResult(self, nums: List[int], k: int) -> int:
        heap = []
        res = 0

        for right in range(len(nums)):
            while heap and heap[0][1] < right - k:
                heapq.heappop(heap)

            max_prev = 0 if not heap else -heap[0][0]
            cur_max = max_prev + nums[right]
            heapq.heappush(heap, (-cur_max, right))
            res = cur_max

        return res