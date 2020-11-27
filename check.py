from flask import Flask, request
import os
app = Flask(__name__)

@app.route('/submit_flag/', methods=['POST'])
def register():
    flag = os.popen("cat flag").read()[:-1]
    if str(flag)==str(request.form['flag']):
        return "success"
    else:
        return "error"

if __name__ == '__main__':
    # app.run(port=8888,debug=True)
    app.run(port=8888,ssl_context ='adhoc',debug=True)
    # If you need https you can use this
    # But you can only run with
    # "python checker.py"
    # no "gunicorn -c gunicorn.conf.py  checker:app"
