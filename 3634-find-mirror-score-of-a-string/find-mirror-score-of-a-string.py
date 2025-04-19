class Solution:
    def calculateScore(self, s: str) -> int:
        occurences = defaultdict(deque)
        res = 0

        def get_mirror_char(c):
            return chr(ord('a') + ord('z') - ord(c))

        for i, c in enumerate(s):
            mirror_char = get_mirror_char(c)

            if len(occurences[mirror_char]) > 0:
                res += i - occurences[mirror_char].pop()
            else:
                occurences[c].append(i)

        return res