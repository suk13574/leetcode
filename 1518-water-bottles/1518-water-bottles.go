func numWaterBottles(numBottles int, numExchange int) int {
    res := numBottles

    for numBottles >= numExchange {
        bottles := numBottles / numExchange
        numBottles = bottles + (numBottles % numExchange)
        
        res += bottles
    }

    return res
}