package com.一般题;

public class distributeCandies1103 {

    static class Solution {
        public static int[] distributeCandies(int candies, int num_people) {
            int[] num_children = new int[num_people];
            for (int j=0 ; j<num_people ; j++ )
                num_children[j]=0;

            int i=1;//糖果
            int cnt=0;
            while(candies>0){
                for (int j=0 ; j<num_people && candies>0 ; j++ ){
                    if(candies>=i){
                        num_children[j] += i;
                        candies -= i++;
                    }else {
                        num_children[j] += candies;
                        candies = 0;
                    }
                }
            }
            return num_children;
        }

        public static void main(String[] args) {
            System.out.println(distributeCandies(7,4));
        }
    }
}
