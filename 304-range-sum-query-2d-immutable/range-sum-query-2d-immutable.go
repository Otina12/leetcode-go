type NumMatrix struct {
    prefixSum [][]int
}


func Constructor(matrix [][]int) NumMatrix {
    m, n := len(matrix), len(matrix[0])
    prefixSum := make([][]int, m + 1)

    for i := 0; i <= m; i++ {
        prefixSum[i] = make([]int, n + 1)
    }

    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            prefixSum[i][j] = matrix[i-1][j-1]
            prefixSum[i][j] += prefixSum[i-1][j] + prefixSum[i][j-1] - prefixSum[i-1][j-1]
        }
    }

    return NumMatrix { prefixSum }
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    return this.prefixSum[row2+1][col2+1] -
           this.prefixSum[row1][col2+1] -
           this.prefixSum[row2+1][col1] +
           this.prefixSum[row1][col1]
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */