console.log('Yup. It works!');

function findCustomer() {
  let custId = document.querySelector('#customer_id').value;
  custId = parseInt(custId);
  searchCustomerById(custId);
  return false;
}

function searchCustomerById(id) {
  let url = `http://localhost:7788/api/customers/${id}`;
  fetch(url, {
    headers: {
      Accept: 'application/json',
    },
  })
    .then((resp) => resp.json())
    .then((c) => {
      document.querySelector('#name').innerHTML = c.name;
      document.querySelector('#city').innerHTML = c.city;
      document.querySelector('#email').innerHTML = c.email;
    });
}
