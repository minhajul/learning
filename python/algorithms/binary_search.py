def binary_search(lists, x):
    left = 0
    right = len(lists) - 1

    while left <= right:
        mid = (left + right) // 2
        if x == lists[mid]:
            return mid

        if x > lists[mid]:
            left = mid + 1

        if x < lists[mid]:
            right = mid - 1

    return -1

lists = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
print(binary_search(lists, 3))
