class Employee:
    def __init__(self, **kwargs) -> None:
        self.id = kwargs.get('id', 0)
        self.name = kwargs.get('name')
        self.salary = kwargs.get('salary', 45000)

    # getter property (accessor or readable property)
    @property
    def id(self):
        return self.__id
    
    # setter for the above getter; mutator or writable property
    @id.setter
    def id(self, value):
        if value is not None and type(value) is not int:
            raise TypeError('id must be int')
        if value is not None and value<=0:
            raise ValueError('id must be > 0')
        
        self.__id = value

    @property
    def name(self):
        return self.__name
    
    @name.setter
    def name(self, value):
        if value is not None and type(value) is not str:
            raise TypeError('name must be <str>')
        if value is not None and len(value)>25:
            raise ValueError('name must be less than 26 letters')
        
        self.__name = value

    @property
    def salary(self):
        return self.__salary
    
    @salary.setter
    def salary(self, value):
        if value is not None and type(value) not in (int, float):
            raise TypeError('salary must be a number')
        if value is not None and value<45000:
            raise ValueError('salary must be >= 45000')
        if value is not None and value>1000000:
            raise ValueError('salary must be <= 1000000')
        
        self.__salary = value


    def print(self):
        print(f'ID       : {self.__id}')
        print(f'Name     : {self.__name}')
        print(f'Salary   : {self.__salary}')
        print()
