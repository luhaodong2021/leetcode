//go:build ignore 
#include <algorithm>
#include <string>
#include <vector>
#include <iostream>

using namespace std;

class Solution {
public:
    string addStrings(string num1, string num2) {
        vector<char> s1;
        int carry = 0;
        string result;
        reverse(num1.begin(), num1.end());
        reverse(num2.begin(), num2.end());

        for (int i = 0; i < num1.length() || i < num2.length(); i++) {
            int c1 = i < num1.length() ? num1[i] - '0' : 0;
            int c2 = i < num2.length() ? num2[i] - '0' : 0;
            int tmp = c1 + c2 + carry;

            if (tmp > 9) {
                result.push_back('0' + tmp - 10);
                carry = 1;
            } else {
                result.push_back('0' + tmp);
                carry = 0;
            }
        }
        if (carry == 1) {
            result.push_back('1');
        }

        reverse(result.begin(), result.end());
        return result;
    }
};

int main(int argc, char **argv) {
    if (argc != 3) {
        cout << argv[0] << " num1 num2" << endl;
        return 0;
    }
    Solution s;
    cout << s.addStrings(argv[1], argv[2]) << endl;
    return 0;
}