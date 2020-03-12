package com.一般题;

/**
 * 1013. 将数组分成和相等的三个部分
 * 给你一个整数数组 A，只有可以将其划分为三个和相等的非空部分时才返回 true，否则返回 false。
 *
 * 形式上，如果可以找出索引 i+1 < j 且满足 (A[0] + A[1] + ... + A[i] == A[i+1] + A[i+2] + ... + A[j-1] == A[j] + A[j-1] + ... + A[A.length - 1]) 就可以将数组三等分。
 *
 *
 *
 * 示例 1：
 *
 * 输出：[0,2,1,-6,6,-7,9,1,2,0,1]
 * 输出：true
 * 解释：0 + 2 + 1 = -6 + 6 - 7 + 9 + 1 = 2 + 0 + 1
 * 示例 2：
 *
 * 输入：[0,2,1,-6,6,7,9,-1,2,0,1]
 * 输出：false
 * 示例 3：
 *
 * 输入：[3,3,6,5,-2,2,5,1,-9,4]
 * 输出：true
 * 解释：3 + 3 = 6 = 5 - 2 + 2 + 5 + 1 - 9 + 4
 *
 *
 * 提示：
 *
 * 3 <= A.length <= 50000
 * -10^4 <= A[i] <= 10^4
 */

// 找切分点
public class canThreePartsEqualSum1013 {

    static class Solution {
        public static boolean canThreePartsEqualSum(int[] A) {
            int left=0,mid=1,right=2;
            int leftNum=0,midNum=0,rightNum=0;
            int sum=0;
            for (int i=0; i<A.length ; i++ ){
                sum+=A[i];
            }
            if(sum %3 != 0 )
                return false;
            while(left<A.length){
                leftNum+=A[left++];
                if (leftNum==sum/3)
                    break;
            }
            mid = left;
            if ( mid == A.length - 1 )
                return false;
            while ( mid<A.length ){
                midNum+=A[mid++];
                if (midNum == sum/3)
                    break;
            }
            if ( mid == A.length )
                return false;
            right=mid;
            while( right<A.length){
                rightNum+=A[right++];
                if (rightNum==sum/3)
                    break;
            }
    //            if (right != A.length-1 || rightNum!=sum/3 )
    //                return false;
            return true;
        }

        public static void main(String[] args) {
            int[] A={1,-1,1,-1};
            System.out.println(canThreePartsEqualSum(A));
        }
    }
}
