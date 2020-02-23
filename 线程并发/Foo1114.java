package com.线程并发;


import java.util.concurrent.Semaphore;

/**
 * 我们提供了一个类：
 *
 * public class Foo {
 *   public void one() { print("one"); }
 *   public void two() { print("two"); }
 *   public void three() { print("three"); }
 * }
 * 三个不同的线程将会共用一个 Foo 实例。
 *
 * 线程 A 将会调用 one() 方法
 * 线程 B 将会调用 two() 方法
 * 线程 C 将会调用 three() 方法
 * 请设计修改程序，以确保 two() 方法在 one() 方法之后被执行，three() 方法在 two() 方法之后被执行。
 *
 *  
 *
 * 示例 1:
 *
 * 输入: [1,2,3]
 * 输出: "onetwothree"
 * 解释:
 * 有三个线程会被异步启动。
 * 输入 [1,2,3] 表示线程 A 将会调用 one() 方法，线程 B 将会调用 two() 方法，线程 C 将会调用 three() 方法。
 * 正确的输出是 "onetwothree"。
 * 示例 2:
 *
 * 输入: [1,3,2]
 * 输出: "onetwothree"
 * 解释:
 * 输入 [1,3,2] 表示线程 A 将会调用 one() 方法，线程 B 将会调用 three() 方法，线程 C 将会调用 two() 方法。
 * 正确的输出是 "onetwothree"。
 *  
 *
 * 注意:
 *
 * 尽管输入中的数字似乎暗示了顺序，但是我们并不保证线程在操作系统中的调度顺序。
 *
 * 你看到的输入格式主要是为了确保测试的全面性。
 *
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/print-in-order
 * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

public class Foo1114 {

    static class Foo {

        public Foo() {

        }

        private Semaphore first = new Semaphore(0);
        private Semaphore second = new Semaphore(0);

        public void first(Runnable printFirst) throws InterruptedException {

            // printFirst.run() outputs "first". Do not change or remove this line.
            printFirst.run();
            first.release();
        }

        public void second(Runnable printSecond) throws InterruptedException {

            // printSecond.run() outputs "second". Do not change or remove this line.
            first.acquire();
            printSecond.run();
            second.release();
        }

        public void third(Runnable printThird) throws InterruptedException {

            // printThird.run() outputs "third". Do not change or remove this line.
            second.acquire();
            printThird.run();
        }
    }

    public static void main(String[] args) {
        Foo foo = new Foo();
        //lambda表达式，作用是为了将一段代码段传入某个对象中
        Thread t1 = new Thread(() -> {
            try {
                foo.first(() -> System.out.println("one"));
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        });

        Thread t2 = new Thread(() -> {
            try {
                foo.second(() -> {
                    System.out.println("two");
                });
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        });

        Thread t3 = new Thread(() -> {
            try {
                foo.third(() -> {
                    System.out.println("three");
                });
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        });

        t3.start();
        t2.start();
        t1.start();

    }



}


/**
 *
 *
 * 解法一：Synchronized锁和控制变量
 *
 * public class Foo {
 *     //控制变量
 *     private int flag = 0;
 *     //定义Object对象为锁
 *     private Object lock = new Object();
 *     public Foo() {
 *     }
 *     public void first(Runnable printFirst) throws InterruptedException {
 *         synchronized (lock){
 *             //如果flag不为0则让first线程等待，while循环控制first线程如果不满住条件就一直在while代码块中，防止出现中途跳入，执行下面的代码，其余线程while循环同理
 *             while( flag != 0){
 *                 lock.wait();
 *             }
 *             // printFirst.run() outputs "first". Do not change or remove this line.
 *             printFirst.run();
 *             //定义成员变量为 1
 *             flag = 1;
 *             //唤醒其余所有的线程
 *             lock.notifyAll();
 *         }
 *     }
 *     public void second(Runnable printSecond) throws InterruptedException {
 *         synchronized (lock){
 *             //如果成员变量不为1则让二号等待
 *             while (flag != 1){
 *                 lock.wait();
 *             }
 *             // printSecond.run() outputs "second". Do not change or remove this line.
 *             printSecond.run();
 *             //如果成员变量为 1 ，则代表first线程刚执行完，所以执行second，并且改变成员变量为 2
 *             flag = 2;
 *             //唤醒其余所有的线程
 *             lock.notifyAll();
 *         }
 *     }
 *     public void third(Runnable printThird) throws InterruptedException {
 *         synchronized (lock){
 *             //如果flag不等于2 则一直处于等待的状态
 *             while (flag != 2){
 *                 lock.wait();
 *             }
 *             // printThird.run() outputs "third". Do not change or remove this line.
 *             //如果成员变量为 2 ，则代表second线程刚执行完，所以执行third，并且改变成员变量为 0
 *             printThird.run();
 *             flag = 0;
 *             lock.notifyAll();
 *         }
 *     }
 * }
 * 解法二：CountDownLatch
 *
 * public class Foo {
 *     //声明两个 CountDownLatch变量
 *     private CountDownLatch countDownLatch01;
 *     private CountDownLatch countDownLatch02;
 *
 *     public Foo() {
 *         //初始化每个CountDownLatch的值为1，表示有一个线程执行完后，执行等待的线程
 *         countDownLatch01 = new CountDownLatch(1);
 *         countDownLatch02 = new CountDownLatch(1);
 *     }
 *     public void first(Runnable printFirst) throws InterruptedException {
 *             //当前只有first线程没有任何的阻碍，其余两个线程都处于等待阶段
 *             // printFirst.run() outputs "first". Do not change or remove this line.
 *             printFirst.run();
 *             //直到CountDownLatch01里面计数为0才执行因调用该countDownCatch01.await()而等待的线程
 *             countDownLatch01.countDown();
 *     }
 *     public void second(Runnable printSecond) throws InterruptedException {
 *             //只有countDownLatch01为0才能通过，否则会一直阻塞
 *             countDownLatch01.await();
 *             // printSecond.run() outputs "second". Do not change or remove this line.
 *             printSecond.run();
 *             //直到CountDownLatch02里面计数为0才执行因调用该countDownCatch02.await()而等待的线程
 *             countDownLatch02.countDown();
 *     }
 *     public void third(Runnable printThird) throws InterruptedException {
 *             //只有countDownLatch02为0才能通过，否则会一直阻塞
 *             countDownLatch02.await();
 *             // printThird.run() outputs "third". Do not change or remove this line.
 *             printThird.run();
 *     }
 * }
 *
 * 解法三：Semaphore（信号量）
 * Semaphore与CountDownLatch相似，不同的地方在于Semaphore的值被获取到后是可以释放的，并不像CountDownLatch那样一直减到底
 *
 * 获得Semaphore的线程处理完它的逻辑之后，你就可以调用它的Release()函数将它的计数器重新加1，这样其它被阻塞的线程就可以得到调用了
 *
 * public class Foo03 {
 *     //声明两个 Semaphore变量
 *     private Semaphore spa,spb;
 *     public Foo03() {
 *         //初始化Semaphore为0的原因：如果这个Semaphore为零，如果另一线程调用(acquire)这个Semaphore就会产生阻塞，便可以控制second和third线程的执行
 *         spa = new Semaphore(0);
 *         spb = new Semaphore(0);
 *     }
 *     public void first(Runnable printFirst) throws InterruptedException {
 *             // printFirst.run() outputs "first". Do not change or remove this line.
 *             printFirst.run();
 *             //只有等first线程释放Semaphore后使Semaphore值为1,另外一个线程才可以调用（acquire）
 *             spa.release();
 *     }
 *     public void second(Runnable printSecond) throws InterruptedException {
 *             //只有spa为1才能执行acquire，如果为0就会产生阻塞
 *             spa.acquire();
 *             // printSecond.run() outputs "second". Do not change or remove this line.
 *             printSecond.run();
 *             spb.release();
 *     }
 *     public void third(Runnable printThird) throws InterruptedException {
 *             //只有spb为1才能通过，如果为0就会阻塞
 *             spb.acquire();
 *             // printThird.run() outputs "third". Do not change or remove this line.
 *             printThird.run();
 *     }
 * }
 *
 *
 * 作者：no-one-9
 * 链接：https://leetcode-cn.com/problems/print-in-order/solution/javayou-jie-by-no-one-9/
 * 来源：力扣（LeetCode）
 * 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。l
 */
