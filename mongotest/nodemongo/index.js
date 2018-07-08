const express = require('express');
const app = express();
const mongoose = require('mongoose');

const Song = require('./song-model');

const URI = process.env.MONGOLAB_URL;
if (!URI) {
  console.error('MONGO_URL is not defined in environment');
  process.exit(1);
}
mongoose.connect(
  URI,
  { useNewUrlParser: true }
);
mongoose.Promise = global.Promise;
var db = mongoose.connection;
db.on('error', console.error.bind(console, 'MongoDB connection error:'));

app.get('/songs', async (req, res) => {
  try {
    const result = await Song.find();
    res.json(result);
  } catch (err) {
    res.status(400).send(err.message);
  }
});

app.listen(5000, () => {
  console.log('listening on 5000');
});
