import unittest

from todo import create_app


class BaseAPITestCase(unittest.TestCase):
    def setUp(self):
        app = create_app()
        app.config.update(TESTING=True)
        self.app = app
        self.client = app.test_client()
