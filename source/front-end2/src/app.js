const express = require('express')
const path = require('path');

const app = express()
const port = 3000

// app.get('/', (req, res) => {
// 	// res.sendFile(path.join(__dirname+'/game/index.html'));
// 	res.send("Hello world")
// })

app.use(express.static('game'))
app.use('/dist', express.static('build'))

app.listen(port, () => {
	console.log(`Example app listening at http://localhost:${port}`)
})
