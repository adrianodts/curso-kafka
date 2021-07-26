const User = require('./../model/user')
module.exports = (app) => {
    // app.use('/user', (req, res) => {
    //     User
    //         .find({})
    //         .then((users) => {
    //             //console.log({users: [{ id: "12", name: "Fulano de Tal", username: "fulano", email: "fulano.de@tal.com", password: "teste"}]})
    //             return res.render('index', {users});
    //         })
    //         .catch((error) => {
    //             console.error(error);
    //             return;
    //         });
    // });

    app.use('/user', require('./../controller/user/index'));


}
