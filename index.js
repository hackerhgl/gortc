const https = require('https');

https.get('https://coderbyte.com/api/challenges/json/rest-get-simple', (resp) => {
  
  let data = '';

  // parse json data here...
  
  resp.setEncoding('utf8');
    resp.on('data', (v) => {
        const hobbies = JSON.parse(v).hobbies.slice(0,3).join(", ")
        console.log(hobbies);
    })
});