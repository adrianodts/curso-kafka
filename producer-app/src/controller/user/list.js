const User = require('./../../model/user')

module.exports = (req, res) => {
    User
        .find({})
        .then((users) => {
            return res.render('index', { users });
        })
        .catch((error) => {
            console.error(error);
            return;
        });
};