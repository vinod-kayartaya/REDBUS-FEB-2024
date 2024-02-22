from ex19 import Book
import json

def main():
    b2 = Book(title='Let us c', author='Y Kanitkar', price=499.0)

    # print(json.dumps(b2, cls=Book.BookEncoder))
    
    books = [
        Book(title='Java unleashed', author='John Doe', price=1999.9),
        Book(title='Python unleashed', author='Jane Doe', price=1799.9),
        b2
    ]
    print(json.dumps(books, cls=Book.BookEncoder))

    # books1 = {
    #     'c': b2,
    #     'java': books[0],
    #     'python': books[1]
    # }
    # print(json.dumps(books1, cls=Book.BookEncoder))


if __name__ == '__main__':
    main()
