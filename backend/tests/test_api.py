from django.test import TestCase
from django.test import Client

class API_TESTS(TestCase):
    def test_api_get(self):
        print("Testing METHOD: GET")
        response = self.client.get('http://127.0.0.1:8000/api/bearcatEs/')
        self.assertEqual(response.status_code, 200)

    def test_api_post(self):
        print("Testing METHOD: POST")
        c = Client()
        json_data = {
            "title": "12 Days Bootcamp",
            "description": "Coding bootcamp for beginners.",
            "location": "NWM Building 1",
            "event_date": "2022-04-05T22:46:50Z"
        }
        response = c.post('http://127.0.0.1:8000/api/bearcatEs/', json_data)
        self.assertEqual(response.status_code, 201)

