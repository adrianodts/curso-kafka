const mongoose = require('mongoose');

module.exports = () => {
    mongoose.connect('mongodb://root:root@mongo:27017/test?authSource=admin', { 
        useNewUrlParser: true, 
        useUnifiedTopology: true,
        useFindAndModify: false
    }).then((db) => console.log("db is connected"))
    .catch((err) => console.log(err)); //, { useMongoCliente: true });
    mongoose.Promise = global.Promise;
}