package com.线程并发;

import java.util.concurrent.locks.ReentrantLock;

/**
 * 哲学家从 0 到 4 按 顺时针 编号。请实现函数 void wantsToEat(philosopher, pickLeftFork, pickRightFork, eat, putLeftFork, putRightFork)：
 *
 * philosopher 哲学家的编号。
 * pickLeftFork 和 pickRightFork 表示拿起左边或右边的叉子。
 * eat 表示吃面。
 * putLeftFork 和 pickRightFork 表示放下左边或右边的叉子。
 * 由于哲学家不是在吃面就是在想着啥时候吃面，所以思考这个方法没有对应的回调。
 * 给你 5 个线程，每个都代表一个哲学家，请你使用类的同一个对象来模拟这个过程。在最后一次调用结束之前，可能会为同一个哲学家多次调用该函数。
 *
 *  
 *
 * 示例：
 *
 * 输入：n = 1
 * 输出：[[4,2,1],[4,1,1],[0,1,1],[2,2,1],[2,1,1],[2,0,3],[2,1,2],[2,2,2],[4,0,3],[4,1,2],[0,2,1],[4,2,2],[3,2,1],[3,1,1],[0,0,3],[0,1,2],[0,2,2],[1,2,1],[1,1,1],[3,0,3],[3,1,2],[3,2,2],[1,0,3],[1,1,2],[1,2,2]]
 * 解释:
 * n 表示每个哲学家需要进餐的次数。
 * 输出数组描述了叉子的控制和进餐的调用，它的格式如下：
 * output[i] = [a, b, c] (3个整数)
 * - a 哲学家编号。
 * - b 指定叉子：{1 : 左边, 2 : 右边}.
 * - c 指定行为：{1 : 拿起, 2 : 放下, 3 : 吃面}。
 * 如 [4,2,1] 表示 4 号哲学家拿起了右边的叉子。
 *
 */

public class DiningPhilosophers1226 {

    static class DiningPhilosophers {
        //1个Fork视为1个ReentrantLock，5个叉子即5个ReentrantLock，将其都放入数组中
        private ReentrantLock[] lockList = {new ReentrantLock(),
                new ReentrantLock(),
                new ReentrantLock(),
                new ReentrantLock(),
                new ReentrantLock()};

        //让 1个哲学家可以 “同时”拿起2个叉子(搞个临界区)
        private ReentrantLock pickBothForks = new ReentrantLock();

        // call the run() method of any runnable to execute its code
        public void wantsToEat(int philosopher,
                               Runnable pickLeftFork,
                               Runnable pickRightFork,
                               Runnable eat,
                               Runnable putLeftFork,
                               Runnable putRightFork) throws InterruptedException {

            int leftFork = (philosopher + 1) % 5;	//左边的叉子 的编号
            int rightFork = philosopher;	//右边的叉子 的编号

            pickBothForks.lock();	//进入临界区

            lockList[leftFork].lock();	//拿起左边的叉子
            lockList[rightFork].lock();	//拿起右边的叉子

            pickLeftFork.run();	//拿起左边的叉子 的具体执行
            pickRightFork.run();	//拿起右边的叉子 的具体执行

            pickBothForks.unlock();	//退出临界区

            eat.run();	//吃意大利面 的具体执行

            putLeftFork.run();	//放下左边的叉子 的具体执行
            putRightFork.run();	//放下右边的叉子 的具体执行

            lockList[leftFork].unlock();	//放下左边的叉子
            lockList[rightFork].unlock();	//放下右边的叉子
        }
    }

    public static void main(String[] args) {

    }
}
