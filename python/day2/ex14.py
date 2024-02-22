def print_info(**kwargs):
    print(f'kwargs is of type {type(kwargs)}')
    
    for each_key in kwargs:
        print(f'{each_key} --> {kwargs[each_key]}')


def main():
    b1 = dict(title='let us c', price=499.0)
    print_info(fname='Vinod', lname='Kumar', fav_num=4)
    print_info(**b1)  # all key/value pairs in b1 will be passed as keyword arguments to print_info


if __name__ == '__main__':
    main()
