#include <algorithm>
#include <vector>

using std::vector;

class Solution
{
  public:
    vector<int> pivotArray(vector<int>& nums, int pivot)
    {
        vector<int> left;
        vector<int> right;
        int pivot_cnt = 0;

        for (int num : nums)
        {
            if (num < pivot)
            {
                left.push_back(num);
            }
            else if (num == pivot)
            {
                pivot_cnt++;
            }
            else
            {
                right.push_back(num);
            }
        }

        left.insert(left.end(), pivot_cnt, pivot);
        left.insert(left.end(), right.begin(), right.end());

        return left;
    }
};