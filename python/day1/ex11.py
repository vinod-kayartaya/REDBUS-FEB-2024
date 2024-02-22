from pprint import pprint
import json

def main():
    p1 = dict(name='Vinod', email='vinod@vinod.co', married=True)

    p1['address'] = dict(city='Bangalore', state='Karnataka')
    p1['phones'] = ['9731424784', '9844083934']

    del p1['email']  # p1.pop('email')

    pprint(p1)
    print('-' * 50)
    print(json.dumps(p1, indent=4))
    print(p1.keys())
    print(p1.values())
    print(p1.items())


    print(f'p1["name"] is {p1['name']}')
    # print(f'p1["gender"] is {p1['gender']}')  # would result in KeyError
    print(f'p1.get("gender") is {p1.get('gender', 'Male')}')

    if 'email' in p1:  # checks in p1.keys
        print('p1 has an email and it is', p1['email'])


if __name__ == '__main__':
    main()
