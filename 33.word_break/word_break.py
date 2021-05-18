class Solution:
    def wordBreak(self, s, wordDict):
        """
        :type s: str
        :type wordDict: List[str]
        :rtype: bool
        """

        # approach: use dynamic programming approach to verify
        dp = [i==0 for i in range(len(s)+1)]
        
        for i in range(1, len(s)+1):
            for j in range(i):
                if dp[j] and s[j:i] in wordDict:
                    dp[i] = True
                    break
        return dp[-1]

'''
Complexity Analysis:
    Assume that the string length is n
    I used two loops, so that the time complexity is O(n^2) 
    Since the dynamic programing is applied, Space complexity is S O(n), which needs to construc the dp array.
'''