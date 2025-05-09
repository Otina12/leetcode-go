func reverseBits(num uint32) uint32 {
    var reversedNum uint32

    for i := 0; i < 32; i++ {
        reversedNum <<= 1
        reversedNum |= num & 1
        num >>= 1
    }

    return reversedNum
}