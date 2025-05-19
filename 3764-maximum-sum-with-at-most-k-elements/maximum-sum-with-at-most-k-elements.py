class Solution:
    def maxSum(self, grid: List[List[int]], limits: List[int], k: int) -> int:
        heap = []

        for i, row in enumerate(grid):
            cur_heap = []

            for num in row:
                heapq.heappush(cur_heap, num)
                if len(cur_heap) > limits[i]:
                    heapq.heappop(cur_heap)

            while len(cur_heap) > 0:
                heapq.heappush(heap, -heapq.heappop(cur_heap))

        res = 0
        
        for _ in range(k):
            res += -heapq.heappop(heap)

        return res