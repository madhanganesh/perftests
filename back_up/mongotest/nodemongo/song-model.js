const mongoose = require('mongoose');

var SongSchema = new mongoose.Schema({
  contents: String,
  likes: Number
});

module.exports = mongoose.model('Song', SongSchema);
