#include <algorithm>
#include <vector>

using std::vector;

class Solution
{
public:
    int minOperations(vector<vector<int>> &grid, int x)
    {
        vector<int> arr;
        for (int i = 0; i < grid.size(); i++)
        {
            for (int j = 0; j < grid[i].size(); j++)
            {
                arr.push_back(grid[i][j]);
            }
        }

        int n = arr.size();
        int mod = arr[0] % x;
        for (int i = 1; i < n; i++)
        {
            if (arr[i] % x != mod)
            {
                return -1;
            }
        }

        std::sort(arr.begin(), arr.end());

        int mid = arr[n / 2];

        int res = 0;
        for (int i = 0; i < n; i++)
        {
            res += std::abs(arr[i] - mid) / x;
        }

        return res;
    }
};