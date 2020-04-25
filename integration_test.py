#!/usr/bin/env python3

import json
import re
import urllib.request
import unittest

API = 'http://127.0.0.1:8000/v1'

TEST_USER = {
    'first_name': 'Jean Luc',
    'last_name': 'Picard',
    'address': {
        'line_1': 'Starfleet HQ',
        'line_2': '',
        'city': 'San Francisco',
        'subdivision': 'CA',
        'postal_code': '94530',
    },
}

UUID_REGEXP = re.compile(r'[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}', re.I)

class IntegrationTest(unittest.TestCase):
    def test_create_user(self):
        with urllib.request.urlopen(self._create_user_request()) as f:
            body = f.read().decode('utf-8')
            location = f.getheader('Location')
            user = json.loads(body)
            self.assertEqual(f.code, 201)
            self.assertEqual(location, '%s/user/%s' % (API, user['id']))
            self.assertUserMatches(user, TEST_USER)

    def test_get_user(self):
        user_id = None
        with urllib.request.urlopen(self._create_user_request()) as f:
            body = f.read().decode('utf-8')
            user = json.loads(body)
            user_id = user['id']

        req = urllib.request.Request(
            method='GET',
            url=API+'/user/'+user_id
        )
        with urllib.request.urlopen(req) as f:
            body = f.read().decode('utf-8')
            user = json.loads(body)
            self.assertUserMatches(user, TEST_USER)

    def _create_user_request(self):
        return urllib.request.Request(
            method='POST',
            url=API+'/user',
            data=json.dumps(TEST_USER).encode('utf-8')
        )

    def assertUserMatches(self, actual, expected):
        self.assertRegex(actual['id'], UUID_REGEXP)
        self.assertEqual(actual['first_name'], expected['first_name'])
        self.assertEqual(actual['last_name'], expected['last_name'])

        self.assertEqual(actual['address']['line_1'], expected['address']['line_1'])
        self.assertEqual(actual['address']['line_2'], expected['address']['line_2'])
        self.assertEqual(actual['address']['city'], expected['address']['city'])
        self.assertEqual(actual['address']['subdivision'], expected['address']['subdivision'])
        self.assertEqual(actual['address']['postal_code'], expected['address']['postal_code'])

if __name__ == '__main__':
    unittest.main()
