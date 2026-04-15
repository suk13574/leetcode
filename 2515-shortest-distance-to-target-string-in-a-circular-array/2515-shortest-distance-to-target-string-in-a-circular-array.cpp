#include <vector>

using std::string;
using std::vector;

class Solution
{
public:
    int closestTarget(vector<string> &words, string target, int startIndex)
    {
        int n = words.size();

        for (int i = 0; i < n / 2 + 1; i++)
        {
            int right = (startIndex + i) % n;
            int left = (startIndex + n - i) % n;

            if (words[right] == target || words[left] == target)
            {
                return i;
            }
        }
        return -1;
    }
};