# -*- coding:utf-8 -*-
import math
import time
from decimal import *
import concurrent.futures

def get_pi(prec=1000, verbose=False):
    '''
    小数点以下の桁数がprecの円周率を文字列で返す
    '''
    prec = prec+1+1 #精度 "3."の部分があるので、1つ精度を増やして、さらにもう1けたないと丸目の関係でずれることがあるので、さらに1追加
    N = 2*math.ceil(math.log(prec,2)) #繰り返し回数 wikipediaによると、log2(prec)程度でいいらしいので、念の為、その２倍程度にする

    getcontext().prec = prec

    a = Decimal(1.0)
    b = Decimal(1.0) / Decimal(2.0).sqrt()
    t = Decimal(0.25)
    p = Decimal(1.0)
    # N回の試行を開始
    for _ in range(1, N):
        a_next = (a + b)/Decimal(2)
        b_next = (a*b).sqrt()
        t_next = t - p * (a - a_next)**2
        p_next = Decimal(2)*p

        a = a_next
        b = b_next
        t = t_next
        p = p_next

    #円周率を計算
    pi = (a + b)**2 / (Decimal(4)*t)

    #　実行結果を表示する
    if verbose:
        print("精度：",prec)
        print("円周率:", pi)

    return str(pi)[0:prec]

# Python3 program Miller-Rabin primality test
import random

# Utility function to do
# modular exponentiation.
# It returns (x^y) % p
def power(x, y, p):

    # Initialize result
    res = 1;

    # Update x if it is more than or
    # equal to p
    x = x % p;
    while (y > 0):

        # If y is odd, multiply
        # x with result
        if (y & 1):
            res = (res * x) % p;

            # y must be even now
        y = y>>1; # y = y/2
        x = (x * x) % p;

    return res;

# This function is called
# for all k trials. It returns
# false if n is composite and
# returns false if n is
# probably prime. d is an odd
# number such that d*2<sup>r</sup> = n-1
# for some r >= 1
def miillerTest(d, n):

    # Pick a random number in [2..n-2]
    # Corner cases make sure that n > 4
    a = 2 + random.randint(1, n - 4);

    # Compute a^d % n
    x = power(a, d, n);

    if (x == 1 or x == n - 1):
        return True;

    # Keep squaring x while one
    # of the following doesn't
    # happen
    # (i) d does not reach n-1
    # (ii) (x^2) % n is not 1
    # (iii) (x^2) % n is not n-1
    while (d != n - 1):
        x = (x * x) % n;
        d *= 2;

        if (x == 1):
            return False;
        if (x == n - 1):
            return True;

            # Return composite
    return False;

# It returns false if n is
# composite and returns true if n
# is probably prime. k is an
# input parameter that determines
# accuracy level. Higher value of
# k indicates more accuracy.
def isPrime( n, k):

    # tests = [3, 7, 11, 13, 17, 19, 23, 325, 9375, 28178, 450775, 9780504, 1795265022]
    tests = [3, 7, 11, 13, 17, 19]
    for i in tests:
        if n%i == 0:
            return False

        # Find r such that n =
    # 2^d * r + 1 for some r >= 1
    d = n - 1;
    while (d % 2 == 0):
        d //= 2;

        # Iterate given nber of 'k' times
    for i in range(k):
        if (miillerTest(d, n) == False):
            return False;

    return True;

if  __name__ == '__main__':
    start = time.time()
    pi = get_pi(prec=3000, verbose=True)
    print(pi)
    n = 1600
    strpi = str(pi)
    strpi = "3" + strpi[2:]
    print(strpi)
    while True:
        n += 1
        print(n)
        # executor = concurrent.futures.ProcessPoolExecutor(max_workers=4)
        temp = strpi[n:1000+n]
        # future = executor.submit(isPrime, int(temp),4)

        # if future.result():
        #     print(temp)
        #     break

        if isPrime(int(temp),30):
            print(temp)
            # print("w")
            break

    end = time.time()
    print(end - start)