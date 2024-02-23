import requests
base_url = 'http://172.16.10.68:1234/customers'

def get_customer():
    cid = int(input('Enter the customer id to search: '))
    r = requests.get(f'{base_url}/{cid}')
    if r.status_code == 200:
        c = r.json()
        print(f'Name    : {c['name']}')
        print(f'City    : {c['city']}')
        print(f'Email   : {c['email']}')
    elif r.status_code == 404:
        print(f'No data found for id {cid}')
    else:
        print('Something went wrong. Please try after sometime.')


def main():
    get_customer()

if __name__ == '__main__':
    get_customer()