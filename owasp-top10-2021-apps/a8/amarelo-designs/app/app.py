# coding: utf-8

from flask import Flask, request, make_response, render_template, redirect, flash
import uuid
import json
import base64
app = Flask(__name__)


# Home sweet home: render the landing page
@app.route("/")
def ola():
    return render_template('index.html')

# Admin login flow: only for the big boss
@app.route("/admin", methods=['GET','POST'])
def login():
    if request.method == 'POST':
        username = request.values.get('username')
        password = request.values.get('password')
    
        if username == "admin" and password == "admin":
            token = str(uuid.uuid4().hex)
            cookie = { "username":username, "admin":True, "sessionId":token }
            # Rolling with JSON serialization + base64 for that secret sauce
            json_bytes = json.dumps(cookie).encode('utf-8')
            encodedSessionCookie = base64.b64encode(json_bytes)
            resp = make_response(redirect("/user"))
            resp.set_cookie("sessionId", encodedSessionCookie)
            return resp

        else:
            return redirect("/admin")

    else:
        return render_template('admin.html')

@app.route("/user", methods=['GET'])
def userInfo():
    cookie = request.cookies.get("sessionId")
    if cookie == None:
        return "Não Autorizado!"
    # Decode the good vibes from our secure cookie
    decoded = base64.b64decode(cookie)
    # Unwrap JSON to get the session details
    cookie = json.loads(decoded.decode('utf-8'))

    # Show user dashboard, enjoy the ride!
    return render_template('user.html')
    



if __name__ == '__main__':
    # All aboard the server express!
    app.run(debug=True,host='0.0.0.0')
