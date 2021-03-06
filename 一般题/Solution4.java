package com.一般题;

import java.util.Scanner;

/**
 * 给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
 *
 * 请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
 *
 * 你可以假设 nums1 和 nums2 不会同时为空。
 *
 * 示例 1:
 *
 * nums1 = [1, 3]
 * nums2 = [2]
 *
 * 则中位数是 2.0
 * 示例 2:
 *
 * nums1 = [1, 2]
 * nums2 = [3, 4]
 *
 * 则中位数是 (2 + 3)/2 = 2.5
 *
 */
public class Solution4 {

    public static double findMedianSortedArrays(int[] nums1, int[] nums2) {
        int n=nums1.length;
        int m=nums2.length;
        int left=(n+m+1)/2;
        int right=(n+m+2)/2;
        return (findK(nums1,0,nums2,0,left)+findK(nums1,0,nums2,0,right))/2.0;
    }

    private static double findK(int[] nums1, int i, int[] nums2, int j, int k) {
        if (i>=nums1.length) return nums2[j+k-1];//nums1为空数组，返回nums2的第k个元素
        if (j>=nums2.length) return nums1[i+k-1];//nums2为空数组，返回nums1的第k个元素
        if (k==1) return Math.min(nums1[i],nums2[j]);
        int mid1 = (i+k/2-1 < nums1.length ) ? nums1[i+k/2-1] : Integer.MAX_VALUE;
        int mid2 = (j+k/2-1 < nums2.length ) ? nums2[j+k/2-1] : Integer.MAX_VALUE;
        if (mid1<mid2) return findK(nums1,i+k/2,nums2,j,k-k/2);
        else return findK(nums1,i,nums2,j+k/2,k-k/2);
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int[] nums1=new int[10];
        int[] nums2=new int[10];
        for (int i=0 ; i<10 ;i++){
            String temp=scanner.next();
            if (temp == "/n") break;
            nums1[i]=Integer.valueOf(temp);
        }
        for (int i=0 ; i<10 ;i++){
            String temp=scanner.next();
            if (temp == "/n") break;
            nums2[i]=Integer.valueOf(temp);
        }
        System.out.println(findMedianSortedArrays(nums1,nums2));
    }
}
