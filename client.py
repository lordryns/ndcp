import socket 
import sys 
import threading 

host = socket.gethostbyname(socket.gethostname())
port = 8118


try:
    sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM) 
    sock.connect((host, port))
except:
    print("failed!")
    sys.exit(1)


def listen():
    while True:
        res = sock.recv(2048)
        print(res.decode())

threading.Thread(target=listen, daemon=True).start()
while True:
    message = input("> ")
    sock.sendall(str.encode(message + "\n"))

    
