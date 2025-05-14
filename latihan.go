package main
import "fmt"
func main(){
	var n int
	fmt.Scan(&n)
	hasil := trailingZeroes(n)
	fmt.Print(hasil)
}
func trailingZeroes(n int) int {
    var i, x, jum int
    i = n-1
    jum = 0
    for i > 0{
        n *=i
        i--
    }
    for n > 0{
        x = n%10
        if x == 0{
            jum++
        }
        n = n/10
    }
    return jum
}