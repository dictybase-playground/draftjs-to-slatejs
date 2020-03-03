const fs = require("fs")
const path = require("path")
const html = require("./deserialize").html

require("jsdom-global")()
global.DOMParser = window.DOMParser

const htmlToSlate = (inputFolder, outputFolder) => {
  fs.readdir(inputFolder, (err, files) => {
    if (err) {
      console.error(err)
      process.exit(1) // stop the script
    }

    files.forEach(file => {
      fs.readFile(file, "UTF-8", (err, content) => {
        const fileContent = fs.readFileSync(`${inputFolder}/${file}`)
        const convertedHtml = html.deserialize(fileContent)
        const HtmlString = JSON.stringify(convertedHtml)
        const filenameWithoutExtension = path.basename(file, path.extname(file))

        fs.writeFileSync(
          `./${outputFolder}/${filenameWithoutExtension}.json`,
          HtmlString,
          err => {
            if (err) {
              console.error(err)
            }
          },
        )
      })
      console.log("✅  HTML to Slate Conversion complete!")
    })
  })
}

module.exports = {
  htmlToSlate: htmlToSlate,
}
