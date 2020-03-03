const fs = require("fs")
const Value = require("slate").Value

const verifySlateData = folder => {
  fs.readdir(folder, (err, files) => {
    if (err) {
      console.error(err)
      process.exit(1)
    }
    files.forEach(file => {
      fs.readFile(file, "UTF-8", (err, content) => {
        const fileContent = fs.readFileSync(`${folder}/${file}`)
        try {
          Value.fromJSON(JSON.parse(fileContent))
        } catch (error) {
          console.error(error)
        }
      })
    })
  })
}

module.exports = {
  verifySlateData: verifySlateData,
}
