const express = require('express');
const router = express.Router();

router.get('/', require('./list')),
router.post('/create', require('./create'))

module.exports = router