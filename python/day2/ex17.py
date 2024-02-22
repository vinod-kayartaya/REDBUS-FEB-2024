import argparse
import json
from pprint import pprint

def main():
    parser = argparse.ArgumentParser(description='produces a filtered content from the input json file')
    parser.add_argument('-f', '--filename', help='input json filename', required=True)
    parser.add_argument('-p', '--property', help='property to check', required=True)
    parser.add_argument('-v', '--value', help='value of the property to check', required=True)
    parser.add_argument('-o', '--outfile', help='output file to create')
    parser.add_argument('-i', '--include', help='comma delimited fields to include in the output dataset')

    args = parser.parse_args()
    
    with open(args.filename, encoding='utf-8', mode='rt') as fp:
        data = json.load(fp)
        if type(data) is not list:
            raise ValueError('data in the file is not an array')
        if len(data) == 0:
            raise ValueError('array in the file is empty')
        if type(data[0]) is not dict:
            raise ValueError('array does not contain objects with key/value pairs')
        
        output_data = [
            each_data if args.include is None
            else {key:each_data[key] for key in each_data if key in args.include.split(',')}
            for each_data in data
            if each_data.get(args.property) == args.value
        ]
    
    if args.outfile is None:
        pprint(output_data)
    else:
        with open(args.outfile, mode='wt', encoding='utf-8') as fp:
            json.dump(output_data, fp)
            print(f'{len(output_data)} objects saved to file "{args.outfile}"')

if __name__ == '__main__':
    main()
