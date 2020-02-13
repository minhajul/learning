# def higher_order_function(func, arg):
#     return func(func(func(arg)))
#
#
# def add(x):
#     return x + 5
#
#
# print(higher_order_function(add, 10))
#
# numbers = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
#
#
# def make_double(x):
#     return x * 2
#
#
# results = map(make_double, numbers)
# print(list(results))


# def calculate_sum(*args):
#     result = 0
#     for i in args:
#         result += i
#     return result
#
#
# print(calculate_sum(1, 2, 3))


def concatenate(**kwargs):
    result = ""
    for arg in kwargs.values():
        result += arg
    return result


print(concatenate(a="Real", b="Python", c="Is", d="Great", e="!"))
