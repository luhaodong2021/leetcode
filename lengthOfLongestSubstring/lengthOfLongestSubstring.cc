//go:build ignore 
#include <string>
#include <unordered_map>
#include <iostream>

using namespace std;

class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        int start = 0;
        int result = 0, i;
        unordered_map<char, int> dict;
        char c;
        for (i = 0; i < s.length(); i++) {
            c = s[i];
            if (dict.count(c) == 0) {
                dict[c] = i;
                continue;
            }
            
            if (i - start > result) {
                result = i -start;
            }

            for (int j = start; j < dict[c]; j++) {
                auto it = dict.find(s[j]);
                dict.erase(it);
            }

            start = dict[c] + 1;
            dict[c] = i;
        }

        if (i - start > result) {
            result = i -start;
        }
        return  result;
    }
};

int main(int argc, char *argv[]) {
    Solution s;
    cout << s.lengthOfLongestSubstring("a") << endl;
}
