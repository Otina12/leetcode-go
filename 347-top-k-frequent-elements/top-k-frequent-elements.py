class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        freq = [(k, val) for k, val in Counter(nums).items()]
        freq.sort(key = lambda x : x[1])

        return [key for (key, val) in freq[-k:]]