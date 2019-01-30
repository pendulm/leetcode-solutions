import random

def qsort(arr, left, right):
    if not arr or left >= right:
        return

    pivot = partition(arr, left, right)
    if left < pivot:
        qsort(arr, left, pivot-1)
    if pivot < right:
        qsort(arr, pivot+1, right)

def partition(arr, left, right):
    pivot_value = arr[left]
    start = left
    end = right
    left += 1

    left_to_right = True
    while left < right:
        if left_to_right:
            if arr[left] > pivot_value:
                left_to_right = not left_to_right
                continue
            else:
                left += 1
        else:
            if arr[right] < pivot_value:
                left_to_right = not left_to_right
                arr[left], arr[right] = arr[right], arr[left]
            right -= 1

    // 下面一定要判断，可作为编程考察
    if arr[left] <= pivot_value:
        arr[start], arr[left] = arr[left], arr[start]
        return left
    else:
        arr[start], arr[left-1] = arr[left-1], arr[start]
        return left-1


def test_sort(arr):
    print(arr)
    qsort(arr, 0, len(arr)-1)
    print(arr)

test_sort([1,2,3,4,5])
test_sort([])
test_sort([3])
test_sort([5,4,3,2,1])
test_sort([random.randint(0,100) for i in range(10)])
