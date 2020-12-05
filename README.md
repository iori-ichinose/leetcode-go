# LeetCode to Go

> Created by Iori on 2020/11/20.
>
> GoLang 写成的 LeetCode算法题解

## 文件命名规则

1. 以**纯数字**开头: LeetCode 主站问题
   > e.x. 64-minimum-path-num.go => MainSite Question 64
2. 以**数字**.**数字**.开头，icci结尾: 面试题
   > e.x. 08.13.-pile-box-lcci.go => 面试题 08.13
3. 以**数字**-**罗马数字**.开头，汉语拼音命名，lcof结尾：剑指Offer
   > e.x. 59-II.-dui-lie-de-zui-da-zhi-lcof =>剑指Offer 59-II

> P.S.  
> 已经尽量将个问题分类为各个package，部分函数有对应的Testbench函数，统一从main.go import并调用测试