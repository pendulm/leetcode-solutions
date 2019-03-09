#include <string.h>
/*
 * @lc app=leetcode id=5 lang=c
 *
 * [5] Longest Palindromic Substring
 *
 * https://leetcode.com/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (26.47%)
 * Total Accepted:    476.5K
 * Total Submissions: 1.8M
 * Testcase Example:  '"babad"'
 *
 * Given a string s, find the longest palindromic substring in s. You may
 * assume that the maximum length of s is 1000.
 * 
 * Example 1:
 * 
 * 
 * Input: "babad"
 * Output: "bab"
 * Note: "aba" is also a valid answer.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: "cbbd"
 * Output: "bb"
 * 
 * 
 */
char* longestPalindrome(char* s) {
    if (*s == '\0') {
        return "";
    }
    int len = strlen(s);
    int i, j;
    int longest = 0;
    int start = 0, end = 0;

    for (i = 0; i < len-1; ++i) {
        for (j = i+1; j < len; ++j) {
            if (isPalindrome(s, i, j)) {
                if (j - i + 1 > longest) {
                    longest = j - i + 1;
                    start = i;
                    end = j
                }
            }
        }
    }

                

        }
    }
    
}
