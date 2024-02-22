import sys
from vinutils import to_float


def main():
    # print(sys.argv)
    total = 0
    for each_arg in sys.argv[1:]:
        total += to_float(each_arg)
    
    print(f'total = {total}')
    print(f'average = {total/len(sys.argv)-1}')

if __name__ == '__main__':
    main()
