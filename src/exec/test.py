import resource
# import time

# start = time.time() * 1000 * 1000
resource.setrlimit(resource.RLIMIT_CPU,
                   (10, 10))  # in seconds
resource.setrlimit(resource.RLIMIT_STACK,
                   (32 * 1024 * 1024, 32 * 1024 * 1024))  # in bytes
resource.setrlimit(resource.RLIMIT_AS,
                   (512 * 1024 * 1024, 512 * 1024 * 1024))  # in bytes
# end = time.time() * 1000 * 1000
# with open("./out.txt", 'w+')as f:
#     f.write(str(end - start) + '\n')
