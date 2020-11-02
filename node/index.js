const express = require("express");
const app = express();
const cors = require("cors"); 
var bodyParser = require('body-parser');
var crypto = require('crypto');
var fs = require("fs");
const port = 3000;

app.use(bodyParser.json()); // support json encoded bodies
app.use(bodyParser.urlencoded({ extended: true })); // support encoded bodies

// Simple Usage (Enable All CORS Requests)
app.use(cors());

app.get("/nodejs/write", (req, res) => { 
  var line_num = req.param('num');
  line_num = +line_num-1
  var data = fs.readFileSync('./text.txt', 'utf8');
    var lines = data.split("\n");
 
    if(line_num > 99 || line_num < 0){
      res.send('Line number must be in range 1 to 100');
      return
    }
    
  res.send(lines[line_num]);
});

app.post('/nodejs/sha256', function(req, res) {
  var num1 = req.body.num1;
  var num2 = req.body.num2;
  if (isNaN(num1) || isNaN(num2)){
    res.json({
      "res": "type of inputs must be Number!"
    });
    return
  }
  var sum_int = Number(num1) + Number(num2);
  var sum = sum_int.toString();
  console.log(sum);
  const hash = crypto.createHash('sha256').update(sum).digest('hex');
  res.json({
    "res":hash
  });
});


app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`);
});