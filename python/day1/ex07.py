import sys

def divide(numerator, denominator, /):  # position-only arguments
    return numerator // denominator, numerator % denominator

def main():
    try:
        arg1 = sys.argv[1]  # possible IndexError
        arg2 = sys.argv[2]  # possible IndexError
        num = int(arg1)  # possible ValueError
        den = int(arg2)  # possible ValueError
        quot, rem = divide(num, den)  # possible ZeroDivisionError
        print(f'{num} / {den} is {quot}')
        print(f'{num} % {den} is {rem}')
    except IndexError:
        print('not enough inputs')
    except ValueError:
        print('wrong type of values given. only int expected')
    except ZeroDivisionError:
        print('cannot divide by zero')

if __name__ == '__main__':
    print('before main()')
    main()
    print('after main()')
