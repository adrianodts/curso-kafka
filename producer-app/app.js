const express = require('express');
const morgan = require('morgan');
const methodOverride = require('method-override');
const app = express();

app.use(morgan('dev'));
app.use(methodOverride('_method'));

require('./src/middleware/bodyparser')(app);
require('./src/middleware/views')(app);
require('./src/middleware/routes')(app);
require('./src/middleware/mongoose')(app);

app.listen("3000", "0.0.0.0", () => {
    console.log('Express has been started')
});