def line():
    print('-'*80)

def main():
    l1 = ['x', 'y', 'z', 'a', 'b']
    l2 = [10, 20, 30, 40]
    l3 = [1, 2, 3]
    z1 = zip(l1, l2)  # --> [('x', 10), ('y', 20), ('z', 30)]
    z2 = zip(l1, l2, l3)  # --> [('x', 10, 1), ('y', 20, 2), ('z', 30, 3)]

    d1 = dict(zip(l1, l2))
    print(d1)
    line()

    for each_item in z1:
        print(each_item)
    line()

    for each_item in z2:
        print(each_item)
    

if __name__ == '__main__':
    main()
