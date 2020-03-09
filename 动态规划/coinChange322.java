package com.动态规划;

import java.lang.reflect.Array;
import java.util.Arrays;

/**
 *
 * 给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
 *  *
 *  * 示例 1:
 *  *
 *  * 输入: coins = [1, 2, 5], amount = 11
 *  * 输出: 3
 *  * 解释: 11 = 5 + 5 + 1
 *  * 示例 2:
 *  *
 *  * 输入: coins = [2], amount = 3
 *  * 输出: -1
 *  * 说明:
 *  * 你可以认为每种硬币的数量是无限的。
 *
 *
 */

//先用贪心回溯试试(超时)---如果剪枝剪的好说不定可以用  -- 强大的网友大佬用贪心加dfs写出来了 记
//再用动态规划
public class coinChange322 {
    //贪心回溯
    static class Solution {
        static int res = Integer.MAX_VALUE;
        public static int coinChange(int[] coins , int amount){
            Arrays.sort(coins);
            dfs(coins,coins.length-1,amount,0);
            return res==Integer.MAX_VALUE ? -1 : res;
        }

        private static void dfs(int[] coins, int index, int amount, int cnt) {
            if (index<0)
                return;
            for (int n=amount/coins[index] ; n>=0 ; n-- ){
                int amo=amount-n*coins[index];
                int ncnt = cnt+n;
                /**
                 * 剪枝处理如下，具体位置见代码注释：
                 *
                 * 当 amount==0amount==0 时剪枝，因为大面值硬币需要更多小面值硬币替换，继续减少一枚或几枚大硬币搜索出来的答案【如果有】肯
                 * 定不会更优。
                 * 当 amount!=0amount!=0，但已经使用的硬币数 cntcnt 满足了 cnt+1>=anscnt+1>=ans 时剪枝，因 amountamount 不为 00，
                 * 至少还要使用 11 枚硬币，则继续搜索得到的答案不会更优。是 breakbreak 不是 continuecontinue，因为少用几枚面值大的硬币
                 * 搜索出的方案【如果有】肯定需要更多枚面值小的硬币来替换掉这枚大的硬币【同剪枝 11 理由】，而使用这几枚大的硬币都搜索不到更好
                 * 的解，故应该是 breakbreak。
                 *
                 */
                if (amo == 0){
                    res = Math.min(res,ncnt);
                    break;// 剪枝1
                }
                if (ncnt+1>=res){
                    break;// 剪枝2
                }
                dfs(coins,index-1,amo,ncnt);
            }

        }

//        public static int coinChange(int[] coins, int amount) {
//            return backTrack(0,coins,amount);
//        }
//
//        //回溯
//        public static int backTrack(int index , int[] coins , int amount){
//            if (amount==0)
//                return 0;
//            if ( index<coins.length && amount>0 ){
//                int maxNum = amount/coins[index];
//                int minCoinNum = Integer.MAX_VALUE;
//                for (int x=0 ; x<=maxNum ; x++ ){
//                    if ( amount >= x*coins[index] ){
//                        int res = backTrack(index+1,coins,amount-x*coins[index]);
//                        if ( res != -1 )
//                            minCoinNum = Math.min(res+x,minCoinNum);
//                    }
//                }
//                return minCoinNum==Integer.MAX_VALUE ?-1:minCoinNum;
//            }
//            return -1;
//        }

    //动态规划(自上而下)
//        private static int coinChange(int[] coins, int amount) {
//            if (amount<1) return 0;
//            return dp(coins,amount,new int[amount]);
//        }
//
//        private static int dp(int[] coins, int amount, int[] cnt) {
//            if (amount==0) return 0;
//            if (amount<0) return -1;
//            if (cnt[amount-1] != 0 ) return cnt[amount-1];
//            int min = Integer.MAX_VALUE;
//            for (int coin:coins){
//                int res = dp(coins,amount-coin,cnt);
//                if (res>=0 && res<min)
//                    min = 1+res;
//            }
//            cnt[amount-1] = min == Integer.MAX_VALUE ? -1 : min;
//            return cnt[amount-1];
//        }

        public static void main(String[] args) {
            int[] coins= {1,2,5};
            System.out.println(coinChange(coins,55));
        }




    }
}
