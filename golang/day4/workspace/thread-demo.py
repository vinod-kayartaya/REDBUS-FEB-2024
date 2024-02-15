from threading import Thread
import time

def do_stuff(msg):
    for i in range(10):
        print(f'{i} {msg}')
        time.sleep(0.5)

t1 = Thread(target=do_stuff, args=("hello, world!",))
t2 = Thread(target=do_stuff, args=("welcome to python!",))

print("starting of main thread")
t1.start()
t2.start()
print("main thread finished")
