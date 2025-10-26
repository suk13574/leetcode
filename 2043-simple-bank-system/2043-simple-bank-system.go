type Bank struct {
    bal []int64
}

func Constructor(balance []int64) Bank {
    return Bank{bal: balance}
}

func (b *Bank) validAccount(ac int) bool {
    // 계좌 번호는 1..len(b.bal)
    return ac >= 1 && ac <= len(b.bal)
}

func (b *Bank) Transfer(account1 int, account2 int, money int64) bool {
    if !b.validAccount(account1) || !b.validAccount(account2) {
        return false
    }
    if b.bal[account1-1] < money {
        return false
    }
    b.bal[account1-1] -= money
    b.bal[account2-1] += money
    return true
}

func (b *Bank) Deposit(account int, money int64) bool {
    if !b.validAccount(account) {
        return false
    }
    b.bal[account-1] += money
    return true
}

func (b *Bank) Withdraw(account int, money int64) bool {
    if !b.validAccount(account) {
        return false
    }
    if b.bal[account-1] < money {
        return false
    }
    b.bal[account-1] -= money
    return true
}