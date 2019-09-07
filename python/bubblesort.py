def bubblesort(arr):
    """bubble sort

    :arr: TODO
    :returns: TODO
    >>> bubblesort([1, 2, 3, 4, 5, 6])
    [1, 2, 3, 4, 5, 6]
    >>> bubblesort([6, 5, 4, 3, 2, 1])
    [1, 2, 3, 4, 5, 6]
    >>> bubblesort([2, 1, 3, 4, 5, 6])
    [1, 2, 3, 4, 5, 6]
    >>> bubblesort([1, 2, 3, 4, 6, 5])
    [1, 2, 3, 4, 5, 6]
    >>> bubblesort([3, 1, 6, 5, 4, 2])
    [1, 2, 3, 4, 5, 6]

    """
    swapped = True
    while swapped:
        swapped = False
        for i in range(1, len(arr)):
            if arr[i - 1] > arr[i]:
                arr[i - 1], arr[i] = arr[i], arr[i - 1]
                swapped = True

    return arr


def bubblesort2(arr):
    """TODO: Docstring for .

    :arr: TODO
    :returns: TODO
    >>> bubblesort2([1, 2, 3, 4, 5, 6])
    [1, 2, 3, 4, 5, 6]
    >>> bubblesort2([6, 5, 4, 3, 2, 1])
    [1, 2, 3, 4, 5, 6]
    >>> bubblesort2([2, 1, 3, 4, 5, 6])
    [1, 2, 3, 4, 5, 6]
    >>> bubblesort2([1, 2, 3, 4, 6, 5])
    [1, 2, 3, 4, 5, 6]
    >>> bubblesort2([3, 1, 6, 5, 4, 2])
    [1, 2, 3, 4, 5, 6]

    """
    upper_loop = len(arr)
    while upper_loop > 1:
        new_upper = 0
        for i in range(1, upper_loop):
            if arr[i - 1] > arr[i]:
                arr[i - 1], arr[i] = arr[i], arr[i - 1]
                new_upper = i
        upper_loop = new_upper

    return arr


if __name__ == "__main__":
    import doctest
    doctest.testmod()
