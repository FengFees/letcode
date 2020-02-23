package com.线程并发;

import java.util.Scanner;
import java.util.concurrent.Semaphore;
import java.util.concurrent.locks.Lock;

/**
 * 我们提供一个类：
 *
 * class FooBar {
 *   public void foo() {
 *     for (int i = 0; i < n; i++) {
 *       print("foo");
 *     }
 *   }
 *
 *   public void bar() {
 *     for (int i = 0; i < n; i++) {
 *       print("bar");
 *     }
 *   }
 * }
 * 两个不同的线程将会共用一个 FooBar 实例。其中一个线程将会调用 foo() 方法，另一个线程将会调用 bar() 方法。
 *
 * 请设计修改程序，以确保 "foobar" 被输出 n 次。
 *
 *  
 *
 * 示例 1:
 *
 * 输入: n = 1
 * 输出: "foobar"
 * 解释: 这里有两个线程被异步启动。其中一个调用 foo() 方法, 另一个调用 bar() 方法，"foobar" 将被输出一次。
 * 示例 2:
 *
 * 输入: n = 2
 * 输出: "foobarfoobar"
 * 解释: "foobar" 将被输出两次。
 *
 */


public class FooBar1115 {

//    static class FooBar {
//        private int n;
//        private volatile int flag1=0;
//        private volatile int flag2=0;
//        private Object lock = new Object();
//        public FooBar(int n) {
//            this.n = n;
//        }
//
//        public void foo(Runnable printFoo) throws InterruptedException {
//
//            for (int i = 0; i < n; i++) {
//                synchronized (lock){
//                    while ( flag1 != flag2 )
//                        lock.wait();
//                    // printFoo.run() outputs "foo". Do not change or remove this line.
//                    printFoo.run();
//                    flag1++;
//                    lock.notifyAll();
//                }
//
//            }
//        }
//
//        public void bar(Runnable printBar) throws InterruptedException {
//
//            for (int i = 0; i < n; i++) {
//                synchronized (lock){
//                    while ( flag2 == flag1 )
//                        lock.wait();
//                    // printBar.run() outputs "bar". Do not change or remove this line.
//                    printBar.run();
//                    flag2++;
//                    lock.notifyAll();
//                }
//
//            }
//        }
//    }

    static class FooBar {
        private int n;
        private Semaphore foo = new Semaphore(1);
        private Semaphore bar = new Semaphore(0);
        public FooBar(int n) {
            this.n = n;
        }

        public void foo(Runnable printFoo) throws InterruptedException {

            for (int i = 0; i < n; i++) {
                foo.acquire();
                // printFoo.run() outputs "foo". Do not change or remove this line.
                printFoo.run();
                bar.release();
            }
        }

        public void bar(Runnable printBar) throws InterruptedException {

            for (int i = 0; i < n; i++) {

                bar.acquire();
                // printBar.run() outputs "bar". Do not change or remove this line.
                printBar.run();
                foo.release();
            }
        }
    }

    public static void main(String[] args) {
//        Scanner scanner = new Scanner(System.in);
        FooBar fooBar = new FooBar(100000 );

        Thread t1 = new Thread(()->{
            try {
                fooBar.foo(()->{
                    System.out.println("foo");
                });
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        });

        Thread t2 = new Thread(()->{
            try {
                fooBar.bar(()->{
                    System.out.println("bar");
                });
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        });

//        for (int i=0 ; i<scanner.nextInt() ; i++ ){
            t1.start();
            t2.start();
//        }

    }
}
