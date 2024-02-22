import unittest
from utils.days import max_days
from utils.leap import is_leap
from datetime import datetime


class TestIsLeap(unittest.TestCase):
    def test_wrong_type_for_year(self):
        with self.assertRaises(TypeError):
            is_leap("Asd")


class TestMaxDays(unittest.TestCase):

    def test_valid_31_days_months(self):
        months = [1, 3, 5, 7, 8, 10, 12]
        year = 2024
        expected = 31
        for month in months:
            actual = max_days(month, year)
            self.assertEqual(expected, actual)

    def test_valid_30_days_months(self):
        months = [4, 6, 9, 11]
        year = 2024
        expected = 30
        for month in months:
            actual = max_days(month, year)
            self.assertEqual(expected, actual)

    def test_valid_february_input(self):
        self.assertEqual(29, max_days(2, 2024))
        self.assertEqual(28, max_days(2, 2023))

    def test_current_month_year(self):
        current_month = datetime.now().month
        current_year = datetime.now().year

        expected = 31 if current_month in [1, 3, 5, 7, 8, 10, 12] else 30 if current_month in [4, 6, 9, 11] else 29 if current_year%400==0 or current_year%100!=0 and current_year%4==0 else 28

        self.assertEqual(expected, max_days())

    def test_wrong_type_for_month(self):
        with self.assertRaises(TypeError):
            max_days("asd")

    def test_wrong_type_for_year(self):
        with self.assertRaises(TypeError):
            max_days(year="asd")

    def test_wrong_value_for_year(self):
        with self.assertRaises(ValueError):
            max_days(year=-400)

    def test_wrong_value_for_month(self):
        with self.assertRaises(ValueError):
            max_days(400)



if __name__ == '__main__':
    unittest.main()
