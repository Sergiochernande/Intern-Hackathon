import requests
import random
import json

first_names = ['John', 'Jason', 'Jim', 'Jack', 'Adam', 'Scott', 'Hank', 'Matt', 'Jessica', 'Judy', 'Amy', 'Jill', 'Hannah']
last_names = ['Malone', 'Doe', 'Henderson', 'Reacher', 'Jones', 'Stone', 'Lawrence', 'Sparrow', 'Smith', 'Fuller', 'Miller']
organizations = ['Say Yes Buffalo', 'M&T Bank', 'Odoo', 'NFTA', 'Subway', 'McDonalds', 'Marvel', 'Disney', 'NBC', 'Starbucks', 'Walmart', 'Wegmans', 'Apple', 'CostCo']
job_titles = ['Software Engineer', 'Database Administrator', 'Human Resources Specialist', 'Recruiter', 'Manager', 'Director', 'Accountant', 'Developer Advocate', 'Actor', 'Teacher', 'Professor']
numbers = [i for i in range(1, 10)]


def gen_random_phone_number():
    num = []
    for i in range(10):
        num.append(str(random.choice(numbers)))
    return ''.join(num)


seen = {}

for i in range(200):
    first_name = random.choice(first_names)
    last_name = random.choice(last_names)
    if first_name + last_name in seen:
        continue
    seen[first_name + last_name] = True

    organization = random.choice(organizations)
    email = first_name + '.' + last_name + '@' + organization.replace(' ', '') + '.com'
    body = {
        'first_name': first_name,
        'last_name' : last_name,
        'organization': organization,
        'job_title': random.choice(job_titles),
        'email': email,
        'phone_number': gen_random_phone_number()
    }
    body = json.dumps(body)
    response = requests.post('http://localhost:8000/users', body)
    print(response.json())
