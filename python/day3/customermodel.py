import json

class Customer:
    def __init__(self, **kwargs) -> None:
        self.id = kwargs.get('id')
        self.name = kwargs.get('name')
        self.city = kwargs.get('city')
        self.email = kwargs.get('email')

    class CustomerEncoder(json.JSONEncoder):
        def default(self, obj):
            if isinstance(obj, Customer):
                return obj.__dict__
            else:
                return json.JSONDecoder.default(obj)
            
    class CustomerDecoder(json.JSONDecoder):
        def decode(self, json_string):
            decoded_data = super().decode(json_string)
            if isinstance(decoded_data, dict):
                return Customer(**decoded_data)
            if isinstance(decoded_data, list):
                return [Customer(**customer) for customer in decoded_data]
            return decoded_data  