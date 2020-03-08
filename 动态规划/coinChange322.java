package com.动态规划;

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

//先用贪心回溯试试(超时)---如果剪枝剪的好说不定可以用
//再用动态规划
public class coinChange322 {
    //贪心回溯
    static class Solution {
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
        private static int coinChange(int[] coins, int amount) {
            if (amount<1) return 0;
            return dp(coins,amount,new int[amount]);
        }

        private static int dp(int[] coins, int amount, int[] cnt) {
            if (amount==0) return 0;
            if (amount<0) return -1;
            if (cnt[amount-1] != 0 ) return cnt[amount-1];
            int min = Integer.MAX_VALUE;
            for (int coin:coins){
                int res = dp(coins,amount-coin,cnt);
                if (res>=0 && res<min)
                    min = 1+res;
            }
            cnt[amount-1] = min == Integer.MAX_VALUE ? -1 : min;
            return cnt[amount-1];
        }

        public static void main(String[] args) {
            int[] coins= {1,2,5};
            System.out.println(coinChange(coins,11));
        }




    }
}
