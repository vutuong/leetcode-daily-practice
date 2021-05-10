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

test  = Solution()
print(test.wordBreak("leetcode", ["leet","code"]))

print(test.wordBreak("applepenapple", ["apple","pen"]))
print(test.wordBreak("catsandog", ["cats","dog","sand","and","cat"]))