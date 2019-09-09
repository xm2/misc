"""
leetcode #1 Two Sum

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
"""

def twosum(arr:"list", target:"int")->"(uint, uint)":
    """O(n*n)
    find two elements in arr which can be sum up to target
    
    Arguments:
        arr {list} -- [list of number]
        target {number} -- [target sum result]
    Return:
        (elements1_index, elements2_index) -- index of found elements
        (None, None)  --  if not found
    >>> twosum([23, 1, 57, 89, 300], 58)
    (1, 2)
    >>> twosum([23, 1, 57, 89, 300], 301)
    (1, 4)
    >>> twosum([23, 1, 57, 89, 300], 302)
    (None, None)
    """
    for i in range(len(arr)):
        for j in range(i+1, len(arr)):
            if arr[i] + arr[j] == target:
                return (i, j)

    return (None, None)


def twosum2(arr:"list", target:"int")->"(uint, uint)":
    """O(n)
    find two elements in arr which can be sum up to target
    
    Arguments:
        arr {list} -- [list of number]
        target {number} -- [target sum result]
    Return:
        (elements1_index, elements2_index) -- index of found elements
        (None, None)  --  if not found
    >>> twosum2([23, 1, 57, 89, 300], 58)
    (1, 2)
    >>> twosum2([23, 1, 57, 89, 300], 301)
    (1, 4)
    >>> twosum2([23, 1, 57, 89, 300], 302)
    (None, None)
    """
    visited = {}
    for i in range(len(arr)):
        tmp = (target-arr[i])
        if tmp in visited.keys():
            return (visited[tmp], i)
        visited[arr[i]] = i

    return (None, None)

if __name__ == "__main__":
    import doctest
    doctest.testmod()