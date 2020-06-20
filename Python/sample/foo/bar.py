class Baz:
    def __init__(self, name):
        self.name = name

    def hello(self):
        print ("Hello " + self.name + ".")


class Qux:
    def __init__(self, num):
        self.num = num

    def add(self, n):
        return self.num + n
