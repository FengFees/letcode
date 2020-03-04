package com.一般题;

import java.util.ArrayDeque;
import java.util.HashMap;
import java.util.Queue;

/**
 * 在给定的网格中，每个单元格可以有以下三个值之一：
 *
 * 值 0 代表空单元格；
 * 值 1 代表新鲜橘子；
 * 值 2 代表腐烂的橘子。
 * 每分钟，任何与腐烂的橘子（在 4 个正方向上）相邻的新鲜橘子都会腐烂。
 *
 * 返回直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1。
 */

// bfs
public class orangesRotting994 {
    static class Solution {
        static int[] R = {-1,0,1,0};
        static int[] L = {0,-1,0,1};
        public static int orangesRotting(int[][] grid) {
            int munites = 0 ;
            int r = grid.length , l = grid[0].length;
            Queue<Integer> queue = new ArrayDeque<>();
            HashMap<Integer,Integer> map = new HashMap<>();

            for(int i=0 ; i<r ; i++ ){
                for(int j=0 ; j<l ; j++ ){
                    if (grid[i][j] == 2){
                        int num = i*l + j ;
                        queue.add(num); // 将腐烂的橘子放入队列
                        map.put(num,0); // 记录高度 分钟
                    }
                }
            }

            while( !queue.isEmpty() ){
                int now = queue.remove();
                int rr = now/l;
                int ll = now%l;
                munites=map.get(now);

                for (int k=0 ; k<4 ; k++ ){
                    int inr = rr+R[k];
                    int inl = ll+L[k];
                    if ( inr >=0 && inr<r && inl>=0 && inl<l && grid[inr][inl]==1){
                        grid[inr][inl] =2 ;
                        int nowNums = inr*l+inl;
                        queue.add(nowNums);
                        map.put(nowNums,map.get(now)+1);

                    }
                }

            }

            for (int [] row : grid)
                for (int line : row )
                    if (line == 1)
                        return -1;

            return munites;

        }


        public static void main(String[] args) {
            int[][] grid = {{2,1,1},{1,1,0},{0,1,1}};
            System.out.println(orangesRotting(grid));
        }
    }
}
