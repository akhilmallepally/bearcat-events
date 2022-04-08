# Bearcat Events


## Requirements
- Python3.7
- Django Framework

## Setup
- Clone the project
- Navigate to the project root directory
- Make Migrations
```bash
python3 manage.py makemigrations bearcatE

python3 manage.py migrate
```
- Create super user
```bash
python3 manage.py createsuperuser
```
- Run the server
```bash
    python3 manage.py runserver
``` 
- Add data in the backend via __admin__

- Navigate to the http://127.0.0.1:8000/bearcatEs/ to access the data as the API.

## Testing
- Units tests are prest in `test/`
- To run the unit test
```bash
python3 manage.py test
```
