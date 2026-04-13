#include <vector>
#include <algorithm>

class Solution
{
public:
    int getMinDistance(std::vector<int> &nums, int target, int start)
    {
        int n = nums.size();
        int res = n;

        for (int i = 0; i < n; i++)
        {
            if (nums[i] == target)
            {
                int gap = start - i;
                if (gap < 0)
                {
                    gap = -gap;
                }
                res = std::min(res, gap);
            }
        }

        return res;
    }
};