from mysql.connector  import connect
import json


def get_connection():
    try:
        with open('config.json') as fp:
            config = json.load(fp)
            return connect(**config)
    except Exception:
        print('Something went wrong and cannot go any further.')
        exit(1)


def menu():
    print("""
    *** Main Menu ***
    0. Exit
    1. Add new customer
    2. Search customer by id
    3. Search customers by city
    4. Edit customer
    5. Delete customer
    6. List all customers
    """)
    try:
        return int(input('Enter your choice: '))
    except ValueError:
        return -1

                 
def main():
    while True:
        choice = menu()
        if choice == 0:
            break
        elif choice == 1:
            add_new_customer()
        elif choice == 6:
            list_all_customers()
        elif choice in (2, 3, 4, 5):
            print("feature not ready yet")
        else:
            print('Invalid value. Retry')
            
def add_new_customer():
    sql = 'insert into CUSTOMERS (name, city, email) values (%s, %s, %s)'
    print('Enter new customer details: ')
    name = input('Name   : ')
    city = input('City   : ')
    email = input('Email  : ')

    with get_connection() as conn:
        with conn.cursor() as cur:
            cur.execute(sql, (name, city, email))
            conn.commit()
            print('new customer data added successfully.')


def list_all_customers():
    sql = 'select id, name, city, email from CUSTOMERS'
    with get_connection() as conn:
        with conn.cursor() as cur:
            cur.execute(sql)
            print('-'*111)
            print('%3s %-25s %-40s %-40s' % ('ID', 'Name', 'City', 'Email'))
            print('-'*111)
            for each_rec in cur.fetchall():
                print_rec(each_rec)
            print('-'*111)


def print_rec(rec):
    print('%3d %-25s %-40s %-40s' % rec)


if __name__ == '__main__':
    main()
