type Bank struct {
    Balance []int64
}


func Constructor(balance []int64) Bank {
    return Bank{balance}
}


func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
    if account1 > len(this.Balance) || account2 > len(this.Balance) {
        return false
    }

    if this.Balance[account1-1] < money {
        return false
    }

    this.Balance[account1-1] -= money
    this.Balance[account2-1] += money
    return true
}


func (this *Bank) Deposit(account int, money int64) bool {
    if account > len(this.Balance) {
        return false
    }
    
    this.Balance[account-1] += money
    return true
}


func (this *Bank) Withdraw(account int, money int64) bool {
    if account > len(this.Balance) {
        return false
    }

    if this.Balance[account-1] < money {
        return false
    }

    this.Balance[account-1] -= money
    return true
}


/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */