import math
from pprint import pprint

def main():
    nums = [19, 29, 48, 28, 22, 66, 29, 29, 48, 39, 19, 38, 27]
    unique_nums = sorted(list(set(nums)))

    # filter operations
    even_nums = [each_num for each_num in nums if each_num % 2 == 0]
    odd_nums = [each_num for each_num in nums if each_num % 2]

    # map (transform) operation
    squares = [each_num**2 for each_num in nums]
    # transform a list into a dict
    square_roots = {each_num: math.sqrt(each_num) for each_num in unique_nums}

    print(nums)
    print(even_nums)
    print(odd_nums)
    print(squares)
    pprint(square_roots)

if __name__ == '__main__':
    main()
