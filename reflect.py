#!/usr/bin/env python
# Reflects the requests from HTTP methods GET, POST, PUT, and DELETE
# Written by Nathan Hamiel (2010)

from http.server import HTTPServer, BaseHTTPRequestHandler
from optparse import OptionParser

class RequestHandler(BaseHTTPRequestHandler):

    def do_GET(self):

        request_path = self.path

        print("\n----- Request Start ----->\n")
        print("Request headers:", self.headers)
        print("<----- Request End -----\n")

        self.send_response(200)
        self.send_header("Set-Cookie", "foo=bar")
        self.end_headers()

#        self.wfile.write("<html><head><title>Your X-I2P-DEST* headers are:</title></head>")
#        self.wfile.write("<body><pre><code>")
#        self.wfile.write(self.headers)
#        self.wfile.write("</code></pre>")
#        self.wfile.write("</body></html>")

    def do_POST(self):

        request_path = self.path

        print("\n----- Request Start ----->\n")
        print("Request path:", request_path)

        request_headers = self.headers
        content_length = request_headers.get('Content-Length')
        length = int(content_length) if content_length else 0

        print("Content Length:", length)
        print("Request headers:", request_headers)
        print("Request payload:", self.rfile.read(length))
        print("<----- Request End -----\n")

        self.send_response(200)
        self.end_headers()

#        self.wfile.write("<html><head><title>Your X-I2P-DEST* headers are:</title></head>")
#        self.wfile.write("<body><pre><code>")
#        self.wfile.write(self.headers)
#        self.wfile.write("</code></pre>")
#        self.wfile.write("</body></html>")

    do_PUT = do_POST
    do_DELETE = do_GET

def main():
    port = 8080
    print('Listening on localhost:%s' % port)
    server = HTTPServer(('', port), RequestHandler)
    server.serve_forever()


if __name__ == "__main__":
    parser = OptionParser()
    parser.usage = ("Creates an http-server that will echo out any GET or POST parameters\n"
                    "Run:\n\n"
                    "   reflect")
    (options, args) = parser.parse_args()

    main()
