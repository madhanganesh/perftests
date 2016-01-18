var express = require('express');
var app = express();
var mongoose = require('mongoose');
var port = process.env.PORT || 8080; 

var uri = 'mongodb://localhost:27017/test';
mongoose.connect(uri);
var PersonSchema = new mongoose.Schema({
    firstName: String,
    lastName: String,
    phone: String,
    email: String
});
var PeopleModel = mongoose.model('People', PersonSchema, 'person');

app.get('/people', (req, res) => {
    PeopleModel.find((err, people) => {
        if (err) {
            es.send(err);
        }

        res.json(people);
    });
});

app.listen(port, () => {
    console.log('listening in port 8080..');
});
