from model.employee import Employee

def main():
    emp1 = Employee(id=1122, name='John Doe')
    emp1.print()
    emp1.id = 234  # call to the setter id()
    emp1.name = 'Kishore'
    emp1.salary = 155000
    emp1.print()



if __name__ == '__main__':
    main()
