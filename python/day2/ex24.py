from mysql.connector  import connect  # pip install mysql-connector-python
import json
import copy

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
        elif choice == 4:
            edit_customer()
        elif choice == 6:
            list_all_customers()
        elif choice in (2, 3, 5):
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


def edit_customer():
    try:
        cust_id = int(input('Enter customer id to edit: '))
        sql = 'select name, city, email from CUSTOMERS where id=%s'
        with get_connection() as conn:
            with conn.cursor() as cur:
                cur.execute(sql, (cust_id, ))
                rec = cur.fetchone()
                if rec is None:
                    print(f'no customer data found for id {cust_id}')
                    return
                
                print('Customer data found:')
                print(f'Name   : {rec[0]}')
                print(f'City   : {rec[1]}')
                print(f'Email  : {rec[2]}')

                choice = input('Do you want to edit this? (yes/no) [yes]')
                if choice.strip() == '':
                    choice = 'yes'
                if choice.lower() != 'yes':
                    return
                
                print('Enter customer details to update (press ENTER/RETURN to keep the old value): ')
                new_name = input(f'Name   : ({rec[0]}) ').strip()
                if new_name == '':
                    new_name = rec[0]
                new_city = input(f'City   : ({rec[1]}) ').strip()
                if new_city == '':
                    new_city = rec[1]
                new_email = input(f'Email  : ({rec[2]}) ').strip()
                if new_email == '':
                    new_email = rec[2]
                
                sql = 'update CUSTOMERS set name=%s, city=%s, email=%s where id=%s'
                cur.execute(sql, (new_name, new_city, new_email, cust_id))
                conn.commit()
                print('Customer data updated successfully')
    except ValueError:
        print('Invalid value for id. Retry.')

if __name__ == '__main__':
    main()
