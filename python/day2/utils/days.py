from datetime import datetime
from .leap import is_leap


def max_days(month=None, year=None):
    if month is not None and type(month) is not int:
        raise TypeError(f'month must be an int, but got {type(month)}')
    if year is not None and type(year) is not int:
        raise TypeError(f'year must be an int, but got {type(year)}')
    if month is None:
        month = datetime.now().month
    if year is None:
        year = datetime.now().year

    if year<=0:
        raise ValueError('invalid value for year')

    if month in (4, 6, 9, 11):
        return 30
    elif month == 2:
        return 29 if is_leap(year) else 28
    elif month in (1, 3, 5, 7, 8, 10, 12):
        return 31
    else:
        raise ValueError('invalid value for month')
    

def main():
    days = max_days()
    print(days)

    days = max_days(1, 2023)
    print(days)

    days = max_days(year=2023)
    print(days)

    days = max_days(year=2020, month=4)
    print(days)


if __name__ == '__main__': 
    main()