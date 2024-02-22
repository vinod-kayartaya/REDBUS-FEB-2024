from vinutils import dirr

class Book:
    def __init__(self):
        print(dirr(self), 'before adding attributes')
        print(f'Book.__init__() called and id(self) is {id(self)}')
        self.title = ''
        self.price = 0.0
        print(dirr(self), 'after adding attributes')

    def __str__(self):
        return f'a Book object with title="{self.title}" and price={self.price}'

def main():
    print('creating a Book object')
    b1 = Book()
    print(f'Book object created with id {id(b1)}')
    print(dirr(b1))
    print(b1)
    print(b1.__str__())
    print(Book.__str__(b1))

if __name__ == '__main__':
    main()
