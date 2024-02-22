#! /usr/bin/python3

def print_table(num):
    for i in range(1, 11):
        print(f'{num} X {i} = {num*i}')

if __name__ == '__main__':
    num = input('enter a number: ')
    num = int(num)
    print_table(num)
    print('end of execution')