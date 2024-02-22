import json

class Book:
    def __init__(self, **kwargs) -> None:
        self.title = kwargs.get('title')
        self.author = kwargs.get('author')
        self.price = kwargs.get('price', 0.0)
    
    def __str__(self) -> str:
        title = None if self.title is None else f'"{self.title}"'
        author = None if self.author is None else f'"{self.author}"'
        return f'Book(title={title}, author={author}, price={self.price})'
    
    def to_dict(self):
        return self.__dict__

    class BookEncoder(json.JSONEncoder):
        def default(self, obj):
            return obj.__dict__  if isinstance(obj, Book) else json.JSONEncoder.default(obj)


    class BookDecoder(json.JSONDecoder):
        def decode(self, json_string):
            decoded_data = super().decode(json_string)
            if isinstance(decoded_data, dict):
                return Book(title=decoded_data.get('title'), price=decoded_data.get('price'), author=decoded_data.get('author'))
            if isinstance(decoded_data, list):
                return [Book(title=book.get('title'), price=book.get('price'), author=book.get('author')) for book in decoded_data]
            return decoded_data  
    
def main():
    b1 = Book()
    b2 = Book(title='Let us c', author='Y Kanitkar', price=499.0)
    print(b1)
    print(b2)
    b3 = eval(b2.__str__())  # deep copy
    print(f'b3 is of type {type(b3)} and b3 is {b3}')

    print(json.dumps(b2.to_dict()))

    books = [
        Book(title='Java unleashed', author='John Doe', price=1999.9),
        Book(title='Python unleashed', author='Jane Doe', price=1799.9),
        b2
    ]

    print(json.dumps([b.to_dict() for b in books]))

if __name__ == '__main__':
    main()
