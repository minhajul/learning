def binary_search(list_items, x):
    left = 0
    right = len(list_items) - 1

    while left <= right:
        mid = (left + right) // 2
        if x == list_items[mid]:
            return mid

        if x > list_items[mid]:
            left = mid + 1

        if x < list_items[mid]:
            right = mid - 1

    return -1


lists = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
print(binary_search(lists, 11))
