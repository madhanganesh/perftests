from flask import Flask
from flask import jsonify
from flask import request
from pymongo import MongoClient
from bson.json_util import dumps

client = MongoClient('localhost:27017')
db = client.perftest
app = Flask(__name__)

@app.route('/people', methods=['GET'])
def get_people():
    people = db.person
    return dumps(people.find())

if __name__ == '__main__':
    app.run(debug=False)
