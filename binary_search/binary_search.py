import math

def binary_search(arr: list, number: int):
    low = 0
    high = len(arr)
    count = 0

    while low <= high:
        middle = int(math.floor((low + high) / 2))
        guess = arr[middle]
        count = count + 1

        if guess == number:
            print(f'Performed {count} actions')
            return middle
        
        if guess > number:
            high = middle - 1
        
        if guess < number:
            low = middle + 1

    print(f'Performed {count} actions')
    return None
    

arr = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
print(binary_search(arr, 7))