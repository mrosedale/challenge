from flask import Flask, jsonify

app = Flask(__name__)

# Root route handler that returns JSON
@app.route('/')
def home():
    response = {
        "message": "Hello, World!",
        "status": 200
    }
    return jsonify(response)

# /status route handler that returns a plain string "ok"
@app.route('/status')
def status():
    return "ok", 200

if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=8080)
