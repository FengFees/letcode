package com.线程并发;

import java.util.concurrent.locks.Lock;
import java.util.function.IntConsumer;

/**
 *
 * 假设有这么一个类：
 *
 * class ZeroEvenOdd {
 *   public ZeroEvenOdd(int n) { ... }      // 构造函数
 *   public void zero(printNumber) { ... }  // 仅打印出 0
 *   public void even(printNumber) { ... }  // 仅打印出 偶数
 *   public void odd(printNumber) { ... }   // 仅打印出 奇数
 * }
 * 相同的一个 ZeroEvenOdd 类实例将会传递给三个不同的线程：
 *
 * 线程 A 将调用 zero()，它只输出 0 。
 * 线程 B 将调用 even()，它只输出偶数。
 * 线程 C 将调用 odd()，它只输出奇数。
 * 每个线程都有一个 printNumber 方法来输出一个整数。请修改给出的代码以输出整数序列 010203040506... ，其中序列的长度必须为 2n。
 *
 *  
 *
 * 示例 1：
 *
 * 输入：n = 2
 * 输出："0102"
 * 说明：三条线程异步执行，其中一个调用 zero()，另一个线程调用 even()，最后一个线程调用odd()。正确的输出为 "0102"。
 * 示例 2：
 *
 * 输入：n = 5
 * 输出："0102030405"
 */

public class ZeroEvenOdd1116 {


    static class ZeroEvenOdd {
        private static int n;

        public ZeroEvenOdd(int n) {
            this.n = n;
        }

//        private static volatile int z = 0;
//        private static volatile int e = 1;
//        private static volatile int o = 1;
        private static volatile int flag = 0;

//        private static volatile int i = 1;
        private static Object lock = new Object();

        // printNumber.accept(x) outputs "x", where x is an integer.
        public static void zero(IntConsumer printNumber) throws InterruptedException {
            for ( int i=1 ; i<=n ; i++ ) {
                synchronized (lock) {
                    while (flag != 0 )
                        lock.wait();
                    if (i%2==0)
                        flag=1;
                    else
                        flag=2;
                    printNumber.accept(0);
                    lock.notifyAll();
                }
            }
        }

        public static void even(IntConsumer printNumber) throws InterruptedException {
            for ( int i=2 ; i<=n ; i+=2 ) {
                synchronized (lock) {
                    while (flag != 1 )
                        lock.wait();
                    flag=0;
                    printNumber.accept(i);
                    lock.notifyAll();
                }
            }
        }

        public static void odd(IntConsumer printNumber) throws InterruptedException {
            for ( int i=1 ; i<=n ; i+=2 ) {
                synchronized (lock) {
                    while (flag != 2 )
                        lock.wait();
                    flag=0;
                    printNumber.accept(i);
                    lock.notifyAll();
                }
            }
        }
    }

    public static void main(String[] args) {
        ZeroEvenOdd zeroEvenOdd = new ZeroEvenOdd(10000);
        Thread zero = new Thread( () -> {
            try {
                ZeroEvenOdd.zero(System.out::print);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        });
        Thread even = new Thread( ()-> {
            try {
                ZeroEvenOdd.even(System.out::print);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        });

        Thread odd = new Thread( ()-> {
            try {
                ZeroEvenOdd.odd(System.out::print);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        });

        zero.start();
        even.start();
        odd.start();

    }
}
