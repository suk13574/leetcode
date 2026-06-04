#include <string>
using std::string;

class Solution
{
  public:
    int waviness(int x)
    {
        if (x < 100)
        {
            return 0;
        }

        int cnt = 0;

        string s = std::to_string(x);
        for (int i = 1; i < s.size() - 1; i++)
        {
            if ((s[i] > s[i - 1] && s[i] > s[i + 1]) ||
                (s[i] < s[i - 1] && s[i] < s[i + 1]))
            {
                cnt++;
            }
        }

        return cnt;
    }

    int totalWaviness(int num1, int num2)
    {
        int ans = 0;
        for (int i = num1; i <= num2; i++)
        {
            ans += waviness(i);
        }

        return ans;
    }
};