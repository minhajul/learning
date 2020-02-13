def logger(func):
    def wrapper():
        print('Logging..')
        func()
        print('Logging.. done')

    return wrapper()


@logger
def index():
    print("testing")


index()
