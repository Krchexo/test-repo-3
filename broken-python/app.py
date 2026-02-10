from flsak import Flask

app = Flask(__name__)

@app.route("/")
def home():
    return explode

if __name__ == "__main__":
    app.run(debug=True)
