from http.server import BaseHTTPRequestHandler, HTTPServer
import socket
import os


class MyHandler(BaseHTTPRequestHandler):
    def do_GET(self):
        self.send_response(200)
        self.send_header('Content-type', 'text/plain')
        self.end_headers()
        container_id = os.popen("cat /proc/self/cgroup | grep 'docker/' | tail -1 | sed 's/^.*\\///'").read()
        hostname = socket.gethostname()
        message = f"Container ID: {container_id}\nHostname: {hostname}\n"
        self.wfile.write(bytes(message, "utf8"))


def run():
    address = ('', 80)
    httpd = HTTPServer(address, MyHandler)
    httpd.serve_forever()


if __name__ == '__main__':
    run()
