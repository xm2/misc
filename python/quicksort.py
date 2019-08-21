
def quicksort(arr, low, high):
    """TODO: Docstring for  quicksort.

    :arr: TODO
    :returns: TODO
    >>> quicksort([6, 5, 4, 3, 2, 1], 0, 5)
    [1, 2, 3, 4, 5, 6]
    >>> quicksort([6, 5, 1, 2, 3, 4], 0, 5)
    [1, 2, 3, 4, 5, 6]
    >>> quicksort([1, 2, 3, 4, 6, 5], 0, 5)
    [1, 2, 3, 4, 5, 6]
    >>> quicksort([2, 1, 4, 3, 6, 5], 0, 5)
    [1, 2, 3, 4, 5, 6]
    """
    if low < high:
        p = partition(arr, low, high)
        quicksort(arr, low, p-1)
        quicksort(arr, p+1, high)

    return arr


def partition(arr, low, high):
    """TODO: Docstring for partition.
    :returns: TODO
    >>> partition([6, 5, 4, 3, 2, 1], 0, 5)
    0
    >>> partition([6, 5, 1, 2, 3, 4], 0, 5)
    3
    >>> partition([1, 2, 3, 4, 6, 5], 0, 5)
    4
    >>> partition([2, 1, 4, 3, 6, 5], 0, 5)
    4
    """
    pivot = arr[high]
    i = low
    for j in xrange(low, high):
        if arr[j] < pivot:
            arr[j], arr[i] = arr[i], arr[j]
            i = i + 1

    arr[i], arr[high] = arr[high], arr[i]

    return i
