import json
from pprint import pprint

def fullname(c):
    return f'{"Mr." if c["gender"]=='Male' else "Ms."}{c["first_name"]} {c["last_name"]}'

def main():
    options = dict(encoding='utf-8', mode='rt')
    with open('customers.json', **options) as fp:
        customers = json.load(fp)

        female_customer_names = [
            fullname(a_customer)
            for a_customer in customers
            if a_customer.get('gender') == 'Female'
        ]

        pprint(female_customer_names)
        # print(json.dumps(female_customer_names, indent=3))

if __name__ == '__main__':
    main()
