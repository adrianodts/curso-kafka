//const User = require('./../../model/user')

module.exports = (req, res) => {
    res.redirect('/');

    // let user = new User(req.body);

    // //user.password = user.genHash(user.password);

    // user
    //     .save(user)
    //     .then((user) => {
    //         res.redirect('/');
    //     })
    //     .catch((error) => {
    //         console.error(error);
    //         return;
    //     });
};