const mongoose = require('mongoose');
// const bcrypt = require('bcrypt-nodejs');

require('mongoose-type-email');

const User = new mongoose.Schema({
    name: {
        type: String,
        required: true
    },
    username: {
        type: String,
    },
    email: {
        type: mongoose.SchemaTypes.Email,
        required: true
    },
    password: {
        type: String,
        required: true
    },
}, { strict: true });

// User.methods.genHash = (password) => {
//     return bcrypt.hashSync(password, bcrypt.genSaltSync(7), null);
// };

// User.methods.validatePassword = (password, cb) => {
//     return bcrypt.compare(password, this.password, cb);
// };

module.exports = mongoose.model('User', User);