
def main():
    nums = [19, 29, 48, 28, 22, 66, 29, 29, 48, 39, 19, 38, 27]
    print(nums)
    print(f'element at index 5 is {nums[5]}')
    print(f'3 elements from index 5 is {nums[5:8]}')
    print(f'all elements from index 5 is {nums[5:]}')
    print(f'all elements from index 5 is {nums[5:-1]}; last element is excluded')
    print(f'all elements till index 5 (excluding) is {nums[:5]}')

    print(f'elements at even indices is {nums[::2]}')
    print(f'elements at odd indices is {nums[1::2]}')
    print(f'elements in reverse order is {nums[::-1]}')


if __name__ == '__main__':
    main()
