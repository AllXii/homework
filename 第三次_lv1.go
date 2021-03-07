package main
import (
	"fmt"
	"sync"
	"time"
)
const N = 50005
var (
	prime [N]bool //默认值是false
	locker sync.Mutex
	wg sync.WaitGroup
)
func main(){
	t := time.Now()
	printPrime(1, 50000)
	fmt.Println(time.Now().Sub(t))
}
// @printPrime
// @打印 m, n 之间的素数
// @罗振玺          时间(2021/3/4)
// @m       int    "起始数字"
// @n       int    "结束数字"
func printPrime(_,n int){
	for i := 0 ;i < 10; i++{
		//创建十个协程 每个协程分担5000个数
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			for j := x*(n/10) ;j < x*(n/10)+n/10;j++{
				if j % 2 ==1{
					locker.Lock()
					prime[j] = true
					locker.Unlock()
				}
			}
		}(i)
	}
	wg.Wait()
	prime[1] = false

	group := int(sqrt(float64(n))) + 1
	//fmt.Println(group)

	for i := 0; i < group/28 ;i++{
		//创建八个协程 每个协程有28个组
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			for j := x*28; j < x*28+28; j++{
				//出现的奇数的倍数一定是合数
				if prime[j]{
					for z := j+j; z < n; z+=j{
						locker.Lock()
						prime[z] = false
						locker.Unlock()
					}
				}
			}
		}(i)
	}
	wg.Wait()
	for i := 3; i < n ; i+=2{
		if prime[i]{
			fmt.Println(i)
		}

	}
}
func sqrt(x float64) float64 {
	z := 1.0

	if x < 0 {
		return 0
	} else if x == 0 {
		return 0
	} else {
		for getabs(z*z-x) > 1e-6 {
			z = (z + x/z) / 2
		}
		return z
	}
}
func getabs(x float64) float64 {
	if x < 0 {
		return -x
	}

	if x == 0 {
		return 0
	}

	return x
}