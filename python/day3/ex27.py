from flask import Flask, Response, request
from dao.customerdao import get_all_customers, add_new_customer, get_one_customer, update_customer
from model.customermodel import Customer
import json

app = Flask('customer-api')  # generally used name is __name__


@app.route('/')
def home():
    try:
        with open('ex27_home.html') as file:
            return file.read()
    except FileNotFoundError:
        return '/customers'

def create_response(data=None, status=200, mimetype='application/json'):
    return Response(json.dumps(data, cls=Customer.CustomerEncoder), status=status, mimetype=mimetype)


@app.route('/customers', methods=['GET'])
def handle_get_customers():
    customers = get_all_customers()
    return create_response(customers)


@app.route('/customers', methods=['POST'])
def handle_post_customer():
    cust = request.get_json()  # gives the request body
    cust = Customer(**cust)
    try:
        cust = add_new_customer(cust)
        return create_response(cust, status=201)
    except Exception as e:
        err = dict(message=f'{e}')
        return create_response(err, status=500)


@app.route('/customers/<int:customer_id>', methods=['GET'])
def handle_get_one_customer(customer_id):
    customer = get_one_customer(customer_id)
    return create_response({'message': 'no data found'}, status=404) if customer is None else create_response(customer)


@app.route('/customers/<int:customer_id>', methods=['PATCH'])
def handle_patch_one_customer(customer_id):
    cust = request.get_json()  # gives the request body
    cust = Customer(**cust)
    customer = update_customer(customer_id, cust)
    return create_response({'message': 'no data found'}, status=404) if customer is None else create_response(customer)


app.run(host='0.0.0.0', port=1234, debug=True)