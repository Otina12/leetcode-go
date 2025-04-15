class NumArray:

    def __init__(self, nums: List[int]):
        self.tree = self.build(0, len(nums) - 1, nums)

    def update(self, index: int, val: int) -> None:
        self.updateTree(self.tree, index, val)

    def updateTree(self, node, index, val):
        if node.start == node.end:
            node.val = val
            return node.val
        
        m = (node.start + node.end) // 2
        if index <= m:
            self.updateTree(node.left, index, val)
        else:
            self.updateTree(node.right, index, val)
        
        node.val = node.left.val + node.right.val
        return node.val

    def sumRange(self, left: int, right: int) -> int:
        cur = self.tree
        return self.subtreeSum(cur, left, right)

    def subtreeSum(self, subtree, left, right):
        if not subtree:
            return 0
        if subtree.start == left and subtree.end == right:
            return subtree.val

        m = (subtree.start + subtree.end) // 2

        if right < m:
            return self.subtreeSum(subtree.left, left, right)
        elif left > m:
            return self.subtreeSum(subtree.right, left, right)
        else:
            return self.subtreeSum(subtree.left, left, m) + self.subtreeSum(subtree.right, m + 1, right)

    def build(self, l, r, nums):
        if l == r:
            return SegmentTreeNode(l, r, nums[l])

        m = (l + r) // 2
        left_child = self.build(l, m, nums)
        right_child = self.build(m + 1, r, nums)

        current_node = SegmentTreeNode(l, r, left_child.val + right_child.val, left_child, right_child)
        return current_node
        
class SegmentTreeNode:
    def __init__(self, start, end, val, left = None, right = None):
        self.start = start
        self.end = end
        self.val = val
        self.left = left
        self.right = right

# Your NumArray object will be instantiated and called as such:
# obj = NumArray(nums)
# obj.update(index,val)
# param_2 = obj.sumRange(left,right)