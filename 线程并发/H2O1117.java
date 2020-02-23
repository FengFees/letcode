package com.线程并发;

import java.util.concurrent.Semaphore;

/**
 * 现在有两种线程，氢 oxygen 和氧 hydrogen，你的目标是组织这两种线程来产生水分子。
 *
 * 存在一个屏障（barrier）使得每个线程必须等候直到一个完整水分子能够被产生出来。
 *
 * 氢和氧线程会被分别给予 releaseHydrogen 和 releaseOxygen 方法来允许它们突破屏障。
 *
 * 这些线程应该三三成组突破屏障并能立即组合产生一个水分子。
 *
 * 你必须保证产生一个水分子所需线程的结合必须发生在下一个水分子产生之前。
 *
 * 换句话说:
 *
 * 如果一个氧线程到达屏障时没有氢线程到达，它必须等候直到两个氢线程到达。
 * 如果一个氢线程到达屏障时没有其它线程到达，它必须等候直到一个氧线程和另一个氢线程到达。
 * 书写满足这些限制条件的氢、氧线程同步代码。
 *
 *  
 *
 * 示例 1:
 *
 * 输入: "HOH"
 * 输出: "HHO"
 * 解释: "HOH" 和 "OHH" 依然都是有效解。
 * 示例 2:
 *
 * 输入: "OOHHHH"
 * 输出: "HHOHHO"
 * 解释: "HOHHHO", "OHHHHO", "HHOHOH", "HOHHOH", "OHHHOH", "HHOOHH", "HOHOHH" 和 "OHHOHH" 依然都是有效解。
 *  
 *
 * 限制条件:
 *
 * 输入字符串的总长将会是 3n, 1 ≤ n ≤ 50；
 * 输入字符串中的 “H” 总数将会是 2n；
 * 输入字符串中的 “O” 总数将会是 n。
 *
 */


public class H2O1117 {
    static class H2O {

        private static Semaphore s1,s2,s3,s4;

        public H2O() {
            s1 = new Semaphore(2);// H线程信号量
            s2 = new Semaphore(1);// O线程信号量

            s3 = new Semaphore(0);// 反应条件信号量：始放H原子到达信号
            s4 = new Semaphore(0);// 反应条件信号量：等待O原子到达
        }

        public static void hydrogen(Runnable releaseHydrogen) throws InterruptedException {
            s1.acquire(); // 保证只有两个H线程进入执行
            s3.release(); // 释放H原子到达信号
            s4.acquire(); // 等待O原子到达
            // releaseHydrogen.run() outputs "H". Do not change or remove this line.
            releaseHydrogen.run();
            s1.release(); // 唤醒1个H线程
        }

        public static void oxygen(Runnable releaseOxygen) throws InterruptedException {
            s2.acquire(); // 保证只有1个O线程进入执行
            s4.release(2); // O原子到达之后反应信号释放，因为有2个H线程等待所以释放两个
            s3.acquire(2); // 等待2个H原子到达
            // releaseOxygen.run() outputs "O". Do not change or remove this line.
            releaseOxygen.run();
            s2.release(); // 唤醒1个O线程
        }
    }

    public static void main(String[] args) {
        H2O h2O = new H2O();
        int n = 10;
        Thread t1 = new Thread( ()-> {
            for (int i=0 ; i<2*n ; i++ ){
                try {
                    H2O.hydrogen(()->System.out.println("H"));
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }
        });

        Thread t2 = new Thread( ()-> {
            for (int i=0 ; i<n ; i++ ){
                try {
                    H2O.oxygen(()->System.out.println("O"));
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }
        });

        t1.start();
        t2.start();

    }
}
