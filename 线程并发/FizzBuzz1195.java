package com.线程并发;

import java.util.concurrent.Semaphore;
import java.util.function.IntConsumer;

/**
 * 编写一个可以从 1 到 n 输出代表这个数字的字符串的程序，但是：
 *
 * 如果这个数字可以被 3 整除，输出 "fizz"。
 * 如果这个数字可以被 5 整除，输出 "buzz"。
 * 如果这个数字可以同时被 3 和 5 整除，输出 "fizzbuzz"。
 * 例如，当 n = 15，输出： 1, 2, fizz, 4, buzz, fizz, 7, 8, fizz, buzz, 11, fizz, 13, 14, fizzbuzz。
 *
 * 假设有这么一个类：
 *
 * class FizzBuzz {
 *   public FizzBuzz(int n) { ... }               // constructor
 *   public void fizz(printFizz) { ... }          // only output "fizz"
 *   public void buzz(printBuzz) { ... }          // only output "buzz"
 *   public void fizzbuzz(printFizzBuzz) { ... }  // only output "fizzbuzz"
 *   public void number(printNumber) { ... }      // only output the numbers
 * }
 * 请你实现一个有四个线程的多线程版  FizzBuzz， 同一个 FizzBuzz 实例会被如下四个线程使用：
 *
 * 线程A将调用 fizz() 来判断是否能被 3 整除，如果可以，则输出 fizz。
 * 线程B将调用 buzz() 来判断是否能被 5 整除，如果可以，则输出 buzz。
 * 线程C将调用 fizzbuzz() 来判断是否同时能被 3 和 5 整除，如果可以，则输出 fizzbuzz。
 * 线程D将调用 number() 来实现输出既不能被 3 整除也不能被 5 整除的数字。
 *
 */
public class FizzBuzz1195 {

    static class FizzBuzz {
        private static int n;

        public FizzBuzz(int n) {
            this.n = n;
        }

        private static Semaphore s1 = new Semaphore(0); // 可以被3整除 fizz信号
        private static Semaphore s2 = new Semaphore(0); // 可以被5整除 buzz信号
        private static Semaphore s3 = new Semaphore(0); // 可以被3和5整除 fizzbuzz信号
        private static Semaphore s4 = new Semaphore(0); // 都不可以 则输出整数


        // printFizz.run() outputs "fizz".
        public static void fizz(Runnable printFizz) throws InterruptedException {
            for(int i=1 ; i<=n ; i++ ){
                if ( i%3==0 && i%5!=0 ){
                    s1.acquire();
                    printFizz.run();
                    s4.release();
                }

            }
        }

        // printBuzz.run() outputs "buzz".
        public static void buzz(Runnable printBuzz) throws InterruptedException {
            for(int i=1 ; i<=n ; i++ ){
                if(i%3 != 0 && i%5 == 0 ){
                    s2.acquire();
                    printBuzz.run();
                    s4.release();
                }

            }
        }

        // printFizzBuzz.run() outputs "fizzbuzz".
        public static void fizzbuzz(Runnable printFizzBuzz) throws InterruptedException {
            for(int i=1 ; i<=n ; i++ ){
                if ( i%3 == 0 && i%5 == 0 ){
                    s3.acquire();
                    printFizzBuzz.run();
                    s4.release();
                }
            }
        }

        // printNumber.accept(x) outputs "x", where x is an integer.
        public static void number(IntConsumer printNumber) throws InterruptedException {
            for(int i=1 ; i<=n ; i++ ){
                if( i%3 == 0 && i%5 != 0 ){
                    s1.release();
                    s4.acquire();
                }else if( i%3 != 0 && i%5 == 0 ){
                    s2.release();
                    s4.acquire();
                }else if( i%3 == 0 && i%5 == 0 ){
                    s3.release();
                    s4.acquire();
                }else{
                    printNumber.accept(i);
                }
            }
        }
    }

    public static void main(String[] args) {
        FizzBuzz fizzBuzz = new FizzBuzz(15);
//        Thread fizz = new Thread( ()-> {
//            try {
//                FizzBuzz.fizz(System.out.print("fizz"));
//            } catch (InterruptedException e) {
//                e.printStackTrace();
//            }
//        });
    }
}
