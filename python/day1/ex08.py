"""
Example of using `list` methods

'append', 'index', 'insert', 'clear', 'copy', 'count', 'extend', 'pop', 'remove', 'reverse', 'sort'
"""

import copy

def main():
    nums = [19, 29, 48, 28, 22, 66, 29, 29, 48, 39, 19, 38, 27]
    print(f'count of 22 in nums is {nums.count(22)}')
    print(f'count of 100 in nums is {nums.count(100)}')
    print(f'count of 29 in nums is {nums.count(29)}')
    evens = [10, 20, 30]
    # nums.extend(evens)
    nums += evens  # same as above
    print(nums)
    # nums *= 2  # content of the list gets duplicated [1, 2, 3] becomes [1, 2, 3, 1, 2, 3]
    # print(nums)
    # nums += 'vinod'  # consider 'vinod' as a list ['v', 'i', 'n', 'o', 'd']
    nums += ['vinod']
    print(nums)
    print(nums.pop(), nums)
    print(nums.pop(0), nums)  # for invalid index, will raise error - IndexError
    print(nums.remove(66), nums)  # for value not found, will raise - ValueError

def main3():
    nums1 = [10, 20, [30, 40]]
    nums2 = nums1.copy()  # shallow copy
    nums3 = copy.deepcopy(nums1)
    print(f'id of nums1 is {id(nums1)}')
    print(f'id of nums2 is {id(nums2)}')
    print(f'id of nums3 is {id(nums3)}')
    print(nums1)
    print(nums2)
    print(nums3)
    nums1[0] = 50
    print()
    print(nums1)
    print(nums2)
    print(nums3)
    nums1[2][0] = 99
    print()
    print(nums1)
    print(nums2)
    print(nums3)
    nums1[2].append(88)
    print()
    print(nums1)
    print(nums2)
    print(nums3)

    print(f'id of nums1[2] is {id(nums1[2])}')
    print(f'id of nums2[2] is {id(nums2[2])}')
    print(f'id of nums3[2] is {id(nums3[2])}')


def main2():
    nums = [19, 29, 48, 28, 22, 66, 39, 19, 38, 27]
    print(f'id of nums is {id(nums)}')
    print(f'nums contains {len(nums)} elements')
    nums.append(888)
    print(nums)
    nums.insert(0, 1000)
    print(nums)
    nums.insert(-1, 2000)
    print(nums)
    nums.insert(1000, 555)
    print(nums)
    print(f'index of 555 is {nums.index(555)}')
    nums.insert(-1000, 3000)
    print(nums)
    print(f'index of 3000 is {nums.index(3000)}')
    nums.clear()
    print(nums)
    print(f'id of nums is {id(nums)}')


if __name__ == '__main__':
    main()
