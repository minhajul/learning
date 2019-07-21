def selection_sort(list_items):
    n = len(list_items)

    for i in range(0, n - 1):
        index_min = i

        for j in range(i + 1, n):
            if list_items[j] < list_items[index_min]:
                index_min = j

        if index_min != i:
            list_items[i], list_items[index_min] = list_items[index_min], list_items[i]


lists = [6, 1, 4, 5, 9, 2]
print("Before sorting: ", lists)
selection_sort(lists)
print("After sorting: ", lists)
