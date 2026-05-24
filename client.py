import socket 
import sys 

host = socket.gethostbyname(socket.gethostname())
port = 8118


try:
    sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM) 
    sock.connect((host, port))
except:
    print("failed!")
    sys.exit(1)


while True:
    message = input("> ")
    sock.sendall(str.encode(message + "\n"))

    res = sock.recv(2048)
    print(res.decode())
