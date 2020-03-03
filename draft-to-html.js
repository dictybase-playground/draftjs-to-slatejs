const fs = require("fs")
const path = require("path")
const draftToHtml = require("draftjs-to-html")
const convertToRaw = require("draft-js").convertToRaw
const convertFromRaw = require("draft-js").convertFromRaw

// Need to convert the raw JSON state to ContentState first
// then convert ContentState to a raw JS structure
// then finally convert that to HTML
// https://draftjs.org/docs/api-reference-data-conversion

const draftjsToHTML = (inputFolder, outputFolder) => {
  fs.readdir(inputFolder, (err, files) => {
    if (err) {
      console.error(err)
      process.exit(1)
    }
    files.forEach(file => {
      fs.readFile(file, "UTF-8", (err, content) => {
        const fileContent = fs.readFileSync(`${inputFolder}/${file}`)
        const json = JSON.parse(fileContent)
        const contentState = convertFromRaw(
          JSON.parse(json.data.attributes.content),
        )
        const raw = convertToRaw(contentState)
        const html = draftToHtml(raw)
        const filenameWithoutExtension = path.basename(file, path.extname(file))

        fs.writeFileSync(
          `./${outputFolder}/${filenameWithoutExtension}.html`,
          html,
          err => {
            if (err) {
              console.error(err)
            }
          },
        )
      })
      console.log("✅  Draft.js to HTML conversion complete!")
    })
  })
}

module.exports = {
  draftjsToHTML: draftjsToHTML,
}
