#!/usr/bin/python

import random
import string
import time

def get_random_string(length):
    # With combination of lower and upper case
    result_str = ''.join(random.choice(string.ascii_letters) for i in range(length))
    # print random string
    print("This time we are going with:", result_str)

while True:
    get_random_string(50)
    time.sleep(3)
