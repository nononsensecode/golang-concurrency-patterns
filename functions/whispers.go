package functions

// PassWhisper pass a counter from right to left
func PassWhisper(left, right chan int) {
	left <- 1 + <-right
}