from datetime import datetime

def is_leap(year=None):
    if year is not None and type(year) is not int:
        raise TypeError(f'year must be an int, but got {type(year)}')

    if year is None:
        year = datetime.now().year
    return year % 400 == 0 or year % 4 == 0 and year % 100 != 0, year # returns a tuple here


def main():
    y = 1998
    result, y = is_leap(y)  # destructuring the tuple into two variables
    print(y, result)

    result, y = is_leap()
    print(y, result)

    y = "1996"
    result, y = is_leap(y)
    print(y, result)


if __name__ == '__main__':
    main()
