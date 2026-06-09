#include <vector>

using std::vector;

class Solution
{
  public:
    long long maxTotalValue(vector<int>& nums, int k)
    {
        int mn = INT_MAX;
        int mx = 0;

        for (int num : nums)
        {
            mn = std::min(mn, num);
            mx = std::max(mx, num);
        }

        return 1LL * (mx - mn) * k;
    }
};