def binary_search(list_items, x):
    left = list_items[0]
    right = len(list_items) - 1

    while left <= right:
        mid = round((left + right) / 2)
        if x == list_items[mid]:
            return mid

        if x > list_items[mid]:
            left = mid + 1

        if x < list_items[mid]:
            right = mid - 1

    return -1


result = binary_search([1, 2, 3, 4, 5, 6, 7, 8, 9, 10], 3)
print(result)
