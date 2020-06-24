def binary_search(list, item):
    low = 0
    high = len(list) - 1

    scount = 0
    while low <= high:
        scount += 1
        mid = int((low + high) / 2)
        guess = list[mid]
        if guess == item:
            return mid, scount
        if guess < item:
            low = mid + 1
        else:
            high = mid - 1

    return None, scount


my_list = [9, 10, 20, 50, 100]
item = 100

index, count = binary_search(my_list, item)
print("Search\n    - index: %s\n    - count: %s" % (index, count))
