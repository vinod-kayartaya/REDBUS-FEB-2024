def subtotal(code, *nums):
    # print(f'code is {code} and nums is {nums}')
    match code:
        case 1:
            return sum(nums)
        case 2:
            return len(nums)
        case 3:
            return sum(nums)/len(nums)
        case _:
            return None
        
    
def main():
    s=subtotal(1, 10, 20, 30, 40, 50)
    print(s)
    s=subtotal(2, 10, 20, 30, 40, 50)
    print(s)
    s=subtotal(3, 10, 20, 30, 40, 50)
    print(s)

    numbers = (1, 2, 36, 5, 4, 2)
    s = subtotal(3, *numbers)
    print(s)

if __name__ == '__main__':
    main()
