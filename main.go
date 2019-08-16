package main

import (
	"fmt"
	"math/big" // high-precision math
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

//var c1 chan bool
//var c2 chan *big.Int

// This is the same as the first Machin-based Pi
// program, except that it uses the "big" package's
// infinite-sized integers to get as many digits as we
// want.  It still computes the formula:
// pi := 4 * (4 * arccot(5) - arccot(239))

// We start out by defining a high-precision arc cotangent
// function.  This one returns the response as an integer
// - normally it would be a floating point number.  Here,
// the integer is multiplied by the "unity" that we pass in.
// If unity is 10, for example, and the answer should be "0.5",
// then the answer will come out as 5

func arccot(x int64, unity *big.Int) *big.Int {
	bigx := big.NewInt(x)
	xsquared := big.NewInt(x * x)
	sum := big.NewInt(0)
	sum.Div(unity, bigx)
	xpower := big.NewInt(0)
	xpower.Set(sum)
	n := int64(3)
	zero := big.NewInt(0)
	sign := false

	term := big.NewInt(0)
	for {
		xpower.Div(xpower, xsquared)
		term.Div(xpower, big.NewInt(n))
		if term.Cmp(zero) == 0 {
			break
		}
		if sign {
			sum.Add(sum, term)
		} else {
			sum.Sub(sum, term)
		}
		sign = !sign
		n += 2
	}
	return sum
}

func calcpi() *big.Int {
	r := new(big.Int)
	fmt.Println(r.Binomial(1000, 10))

	digits := big.NewInt(3000)
	unity := big.NewInt(0)
	unity.Exp(big.NewInt(10), digits, nil)
	pi := big.NewInt(0)
	four := big.NewInt(4)
	pi.Mul(four, pi.Sub(pi.Mul(four, arccot(5, unity)), arccot(239, unity)))
	//val := big.Mul(4, big.Sub(big.Mul(4, arccot(5, unity)), arccot(239, unity)))
	//fmt.Println("Hello, Pi:  ", pi)
	return pi
}

func isPrime(p *big.Int) bool {
	zero := big.NewInt(0)
	one := big.NewInt(1)
	two := big.NewInt(2)

	if p.Cmp(two) == 0 {
		return true
	}

	// p - 1 = 2^s * dに分解する
	d := new(big.Int).Sub(p, one)
	s := 0
	for new(big.Int).And(d, one).Cmp(zero) == 0 {
		d.Rsh(d, 1)
		s++
	}

	n := new(big.Int).Sub(p, one)
	k := 20
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < k; i++ {
		result := false
		// [1, n-1]からaをランダムに選ぶ
		a := new(big.Int).Rand(rnd, n)
		a.Add(a, one)

		// a^{2^r * d} mod p != -1(= p - 1 = n)の比較
		tmp := new(big.Int).Exp(a, d, p)
		for r := 0; r < s; r++ {
			if tmp.Cmp(n) == 0 {
				result = true
				break
			}
			tmp.Exp(tmp, two, p)
		}

		// a^d != 1 mod p の比較
		if !result && new(big.Int).Exp(a, d, p).Cmp(one) != 0 {
			return false
		}
	}

	return true
}

func checkStr(pistr string, i int) {
	//fmt.Println(i)

	temp := string(pistr[i+999])
	if temp == "0" || temp == "2" || temp == "4" || temp == "6" || temp == "8" {
		//c1 <- false
		wg.Done()
		return
	}

	n, _ := new(big.Int).SetString(pistr[i:i+1000], 10)

	product := big.NewInt(1)
	zero := big.NewInt(0)
	if product.Mod(n, big.NewInt(3)).Cmp(zero) == 0 || product.Mod(n, big.NewInt(5)).Cmp(zero) == 0 || product.Mod(n, big.NewInt(7)).Cmp(zero) == 0 || product.Mod(n, big.NewInt(11)).Cmp(zero) == 0 || product.Mod(n, big.NewInt(13)).Cmp(zero) == 0 || product.Mod(n, big.NewInt(17)).Cmp(zero) == 0 || product.Mod(n, big.NewInt(19)).Cmp(zero) == 0 || product.Mod(n, big.NewInt(23)).Cmp(zero) == 0 || product.Mod(n, big.NewInt(29)).Cmp(zero) == 0 || product.Mod(n, big.NewInt(31)).Cmp(zero) == 0 || product.Mod(n, big.NewInt(37)).Cmp(zero) == 0 || product.Mod(n, big.NewInt(41)).Cmp(zero) == 0 {
		//c1 <- false
		wg.Done()
		return
	}

	a := isPrime(n)
	//fmt.Println(a)
	if a {
		fmt.Println(n)
		wg.Done()
		return
		//c1 <- true
		//c2 <- n
		//defer wg.Done()
	} else {
		//c1 <- false
		wg.Done()
		return
	}
	//c1 <- false

	//close(c1)
	//close(c2)
	return
}

func main() {
	start := time.Now()

	pi := calcpi()
	fmt.Println(pi)
	pistr := pi.String()

	//c1 = make(chan bool)
	//c2 = make(chan *big.Int)
	//c1 := make(chan bool)
	//c2 := make(chan *big.Int)

	for i := 0; i < 1700; i += 1 {
		wg.Add(1)
		go checkStr(pistr, i)
		//v := <-c1
		//select {
		//case v := <-c1:
		//	fmt.Println(v)
		//if v == true {
		//	fmt.Println(v)
		//}
		//}
		//default:
		//	fmt.Println("no value")
		//}
		//if v == true {
		//	v2 := <- c2
		//	fmt.Println(v2)
		//	break
		//}

		//wg.Add(1)
		//fmt.Println(i)
	}

	//go func() {
	//	for x := range c1 {
	//		// loop over all items
	//		in <- x
	//	}
	//	close(in)
	//}()
	wg.Wait()
	//close(c1)

	//close(c2)
	elapsed := time.Since(start)
	fmt.Printf("Time %s", elapsed)
}
