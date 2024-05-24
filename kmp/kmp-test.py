from kmp import kmp
import unittest

test_data_list = [
    {
        "s": "abxabcabcaby",
        "pattern": "abcaby",
        "match": True
    },
    {
        "s": "abcbcglx",
        "pattern": "bcgl",
        "match": True
    },
    {
        "s": "abcxabcdabxabcdabcdabcy",
        "pattern": "abcdabcy",
        "match": True
    },
    {
        "s": "aaaaab",
        "pattern": "aab",
        "match": True
    },
    {
        "s": "abcdabce",
        "pattern": "bcc",
        "match": False
    },
    {
        "s": "abcdefgh",
        "pattern": "def",
        "match": True
    },
    {
        "s": "hello world",
        "pattern": "hell test",
        "match": False
    },
    {
        "s": "hello world",
        "pattern": "hello",
        "match": True
    },
     {
        "s": "hello world",
        "pattern": "ello worl",
        "match": True
    }
]

# Knuth Morris Pratt string match algorithm test
class KMPTest(unittest.TestCase):
    def test_kpm_search(self):
        for test_data_object in test_data_list:
            s = test_data_object.get("s")
            pattern = test_data_object.get("pattern")
            expected = test_data_object.get("match")

            result = kmp(s, pattern)
            error_msg = f"Expected {expected} for string \"{s}\" using pattern \"{pattern}\". Got {result}"

            self.assertEqual(result, expected, error_msg)


if __name__ == '__main__':
    unittest.main()