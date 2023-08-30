package com.栈;

import java.util.ArrayDeque;
import java.util.Deque;

/**
 * 739. 每日温度
 * 根据每日 气温 列表，请重新生成一个列表，对应位置的输出是需要再等待多久温度才会升高超过
 * 该日的天数。如果之后都不会升高，请在该位置用 0 来代替。
 *
 * 例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，
 * 你的输出应该是 [1, 1, 4, 2, 1, 1, 0, 0]。
 *
 * 提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的均为华氏度，
 * 都是在 [30, 100] 范围内的整数。
 */
public class dailyTemperatures739 {
    // 单调栈！
    public static int[] dailyTemperatures(int[] T) {
        // 先将数组入栈，从最后往前走，再用dp记录，如果出栈比前一个出栈数小，则该位置为1
        // 如果比该数大，则比较前一个数的位置数的那个温度，不断比较。
        int[] result = new int[T.length];
        if (T.length==1) return new int[]{0};
        Deque<Integer>  tem = new ArrayDeque<>();
        for (int t: T
             ) {
            tem.push(t);
        }

        result[T.length-1]=0;
        int beforeTem = tem.pop(); // 把最后一个温度弹出
        int len=T.length-1;// 记录长度
        while (!tem.isEmpty()){
            len--;
            int temp = tem.pop();
            if (temp<beforeTem){ // 如果当前温度小于后一个温度，则肯定为1
                result[len]=1;
            }
            else{// 如果当前温度大于等于后一个温度，则从后一个温度往后找比当前温度大的。
                // 如果找到了，则为该温度位置减去当前温度位置，找不到则为0
                int l=len+1;
                int d=result[l];
                l+=d;
                while(result[l]!=0 && l<T.length-1 && temp>=T[l]){
                    d=result[l];
                    l+=d;
                }

                if (result[l]==0){
                    result[len] = T[len]<T[l] ? l-len : 0;
                }
                else if (l>=T.length-1)
                    result[len]=0;
                else {
                    result[len]=l-len;
                }

            }
            beforeTem = temp;
        }

        return result;
    }

    public static void main(String[] args) {
//        int[] t={73, 74, 75, 71, 69, 72, 76, 73};
        int[] t ={34,80,80,34,34,80,80,80,80,34};
        int[] res=dailyTemperatures(t);
        for (int i: res
             ) {
            System.out.printf("%d\t",i);
        }
    }
}
