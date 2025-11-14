import time

start = time.perf_counter()

sum_val = 0
for i in range(100_000_000):
    sum_val += i

end = time.perf_counter()
print("Time:", (end - start), "seconds")