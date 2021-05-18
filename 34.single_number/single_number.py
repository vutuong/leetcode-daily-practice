'''
https://leetcode.com/problems/single-number/
136. Single Number

Add to List

Share
Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

You must implement a solution with a linear runtime complexity and use only constant extra space.
'''
class Solution:
    def singleNumber(self, nums) -> int:
        dic = {}
        for i in nums:
            dic[i] =dic.get(i,0) + 1
        print(dic)
        for key, val in dic.items():
            if val ==1:
                return key