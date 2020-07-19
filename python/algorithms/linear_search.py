def linear_search(lists, x):
    for i in range(len(lists)):
        if lists[i] == x:
            return i
    return -1


lists = [1, 4, 3, 6, 8, 7, 2, 5]
print(linear_search(lists, 3))
