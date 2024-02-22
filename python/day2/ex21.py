from ex19 import Book
import json

def main():
    books_str = '[{"title": "Java unleashed", "author": "John Doe", "price": 1999.9}, {"title": "Python unleashed", "author": "Jane Doe", "publisher": "ACME publisher"}, {"title": "Let us c", "author": "Y Kanitkar", "price": 499.0}]'

    books = json.loads(books_str, cls=Book.BookDecoder)

    for each_book in books:
        print(type(each_book), each_book)


if __name__ == '__main__':
    main()
