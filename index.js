const http = require('http')
const { URLSearchParams } = require('url')

const port = process.env.PORT ?? 3000

const recursiveFibbonaci = (num) => {
  if (num <= 2) {
    return 1
  }

  return recursiveFibbonaci(num-1 ) + recursiveFibbonaci(num-2)
}

const loopFibbonaci = (num) => {
  let ans = 0
  for (let i = 0; i <= num; i++) {
    if (i <= 2) {
      ans++
    } else {
      ans += i
    }
  }

  return ans
}

const server = http.createServer()
server.on('request', (req, res) => {
  const params = new URLSearchParams(req.url.split('?')[1]);
  const num = params.get('number') ?? 10;
  const recursive = params.get('recursive');

  res.setHeader('Content-Type', 'application/json')
  res.write(JSON.stringify({
    number: num,
    recursive: !!recursive,
    fibonnaci: recursive ? recursiveFibbonaci(num) : loopFibbonaci(num)
  }))

  res.end()
})

console.log('Nodejs server started...')
server.listen(port)
