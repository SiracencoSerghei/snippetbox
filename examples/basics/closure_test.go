package examples

import "testing"

func TestClosureDemo(t *testing.T) {
    
c := adder()

println(c(1)) // 1
println(c(2)) // 3
println(c(3)) // 6
}